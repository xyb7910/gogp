package channeldemo

type Consumer struct {
	ch chan string
}

// Broker 实现方式一： 每一个消费者都有一个独立的channel
type Broker struct {
	consumer []*Consumer
}

func (b *Broker) Producer(msg string) {
	for _, v := range b.consumer {
		v.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumer = append(b.consumer, c)
}

// Broker1 实现方式二： 每一个消费者都共享一个channel, 轮询所有消费者
type Broker1 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *Broker1) Producer(msg string) {
	b.ch <- msg
}

func (b *Broker1) Subscribe(c func(s string)) {
	b.consumers = append(b.consumers, c)
}

func (b *Broker1) Start() {
	go func() {
		for {
			msg, ok := <-b.ch
			if !ok {
				return
			}
			for _, c := range b.consumers {
				c(msg)
			}
		}
	}()
}

func NewBroker1() *Broker1 {
	broker := &Broker1{
		ch:        make(chan string, 10),
		consumers: make([]func(s string), 0, 10),
	}
	go func() {
		for {
			msg, ok := <-broker.ch
			if !ok {
				return
			}
			for _, c := range broker.consumers {
				c(msg)
			}
		}
	}()
	return broker
}
