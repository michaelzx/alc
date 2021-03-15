package alc_types

import (
	"database/sql/driver"
	"encoding/json"
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
	return []byte(t), nil
}

func (t JsonString) ToJsonObj() JsonObj {
	var v JsonObj
	_ = json.Unmarshal([]byte(t), &v)
	return v
}
func (t JsonString) ToJsonArray() JsonArray {
	var v JsonArray
	_ = json.Unmarshal([]byte(t), &v)
	return v
}
