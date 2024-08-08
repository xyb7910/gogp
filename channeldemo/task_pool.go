package channeldemo

// TaskPool 固定大小的任务池
type TaskPool struct {
	ch chan struct{}
}

func NewTaskPool(size int) *TaskPool {
	t := &TaskPool{
		ch: make(chan struct{}, size),
	}
	// 提前准备好令牌
	for i := 0; i < size; i++ {
		t.ch <- struct{}{}
	}
	return t
}

func (tp *TaskPool) Do(f func()) {
	// 从令牌池中获取一个令牌
	token := <-tp.ch
	go func() {
		f()
		// 任务执行完毕，将令牌放回令牌池
		tp.ch <- token
	}()
}

type TaskPoolWithCache struct {
	cache chan func()
}

func NewTaskPoolWithCache(limit, cacheSize int) *TaskPoolWithCache {
	t := &TaskPoolWithCache{
		cache: make(chan func(), cacheSize),
	}

	for i := 0; i < limit; i++ {
		go func() {
			for {
				select {
				case task, ok := <-t.cache:
					if !ok {
						return
					}
					task()
				}
			}
		}()
	}
	return t
}

func (t *TaskPoolWithCache) Do(f func()) {
	t.cache <- f
}
