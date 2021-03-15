package alc_types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type NilErr struct {
}

func NewNilErr() *NilErr {
	return &NilErr{}
}

func (n NilErr) Error() string {
	return "对象为空"
}

type TypeErr struct {
	msg string
}

func NewTypeErr(msg string) *TypeErr {
	return &TypeErr{msg: msg}
}

func (t TypeErr) Error() string {
	return t.msg
}

type JsonArray []interface{}

func (t *JsonArray) Scan(src interface{}) error {
	// 从数据库中拿出来
	b, ok := src.([]byte)
	if !ok {
		return errors.New("JsonString must be []byte")
	}
	var _t []interface{}
	err := json.Unmarshal(b, &_t)
	if err != nil {
		return err
	}
	*t = _t
	return nil
}
func (t JsonArray) Value() (driver.Value, error) {
	// 放到数据库里面去
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (t JsonArray) ToStringSlice() ([]string, error) {
	list := make([]string, 0, 0)
	for _, si := range t {
		if s, ok := si.(string); ok {
			list = append(list, s)
		} else {
			return nil, NewTypeErr("不是string类型")
		}
	}
	return list, nil
}
