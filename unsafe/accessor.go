package unsafe

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type FieldMeta struct {
	offset uintptr
}
type UnsafeAccessor struct {
	fields     map[string]FieldMeta // 结构体的字段信息
	entityAddr unsafe.Pointer       // 结构体的起始地址
}

func NewUnsafeAccessor(entity any) (*UnsafeAccessor, error) {
	if entity == nil {
		return nil, errors.New("entity is nil")
	}
	val := reflect.ValueOf(entity)
	typ := reflect.TypeOf(entity)

	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("entity must be a pointer to a struct")
	}
	fields := make(map[string]FieldMeta, typ.Elem().NumField())
	elemType := typ.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		fd := elemType.Field(i)
		fields[fd.Name] = FieldMeta{offset: fd.Offset}
	}
	return &UnsafeAccessor{
		entityAddr: val.UnsafePointer(),
		fields:     fields,
	}, nil
}

func (u *UnsafeAccessor) Field(field string) (int, error) {
	fdMeta, ok := u.fields[field]
	if !ok {
		return 0, errors.New("field not found")
	}
	ptr := unsafe.Pointer(uintptr(u.entityAddr) + fdMeta.offset)
	if ptr == nil {
		return 0, fmt.Errorf("invalid address of field %s", field)
	}
	res := *(*int)(ptr) // 读取： *(*T)(ptr)
	return res, nil
}

func (u *UnsafeAccessor) SetField(field string, value int) error {
	fdMeta, ok := u.fields[field]
	if !ok {
		return fmt.Errorf("invalid field %s", field)
	}
	ptr := unsafe.Pointer(uintptr(u.entityAddr) + fdMeta.offset)
	if ptr == nil {
		return fmt.Errorf("invalid address of field %s", field)
	}
	*(*int)(ptr) = value // 写入： *(*T)(ptr) = value
	return nil
}
