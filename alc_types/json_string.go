package alc_types

import (
	"database/sql/driver"
	"errors"
)

type JsonString string

func (t *JsonString) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("JsonString must be []byte")
	}
	str := string(b)
	result := JsonString(str)
	*t = result
	return nil
}

func (t JsonString) Value() (driver.Value, error) {
	return []byte(string(t)), nil
}
