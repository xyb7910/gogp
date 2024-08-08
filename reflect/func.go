package reflect

import (
	"errors"
	"reflect"
)

type FuncInfo struct {
	Name   string
	In     []reflect.Type
	Out    []reflect.Type
	Result []any
}

// CallMethod 输出方法的信息并执行调用, 输出：方法名，方法参数，返回值
func CallMethod(val any) (map[string]*FuncInfo, error) {
	typ := reflect.TypeOf(val)
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Ptr {
		return nil, errors.New("非法类型")
	}
	// 构建结果集
	num := typ.NumMethod()
	result := make(map[string]*FuncInfo, num)
	for i := 0; i < num; i++ {
		fc := typ.Method(i)
		numIn := fc.Type.NumIn()
		ps := make([]reflect.Value, 0, fc.Type.NumIn())
		ps = append(ps, reflect.ValueOf(val))
		in := make([]reflect.Type, 0, fc.Type.NumIn())
		for j := 0; i < numIn; j++ {
			p := fc.Type.In(j)
			in = append(in, p)
			if j > 0 {
				ps = append(ps, reflect.Zero(p))
			}
		}
		ret := fc.Func.Call(ps)
		outNum := fc.Type.NumOut()
		out := make([]reflect.Type, 0, outNum)
		res := make([]any, 0, outNum)
		for k := 0; k < outNum; k++ {
			out = append(out, fc.Type.Out(k))
			res = append(res, ret[k].Interface())
		}
		result[fc.Name] = &FuncInfo{
			Name:   fc.Name,
			In:     in,
			Out:    out,
			Result: res,
		}
	}
	return result, nil
}
