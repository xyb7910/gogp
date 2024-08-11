package sql

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonColumn[T any] struct {
	Val   T
	Valid bool
}

func (jc JsonColumn[T]) Value() (driver.Value, error) {
	if !jc.Valid {
		return nil, nil
	}
	return json.Marshal(jc.Val)
}

func (jc *JsonColumn[T]) Scan(src any) error {
	var bs []byte
	switch val := src.(type) {
	case []byte:
		bs = val
	case *[]byte:
		bs = *val
	case string:
		bs = []byte(val)
	case sql.RawBytes:
		bs = val
	case *sql.RawBytes:
		bs = *val
	default:
		return fmt.Errorf("unsupported type %T", src)
	}
	if err := json.Unmarshal(bs, &jc.Val); err != nil {
		return err
	}
	jc.Valid = true
	return nil
}
