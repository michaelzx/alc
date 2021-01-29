package alc_types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JsonObj map[string]interface{}

func (t *JsonObj) Scan(src interface{}) error {
	// 从数据库中拿出来
	b, ok := src.([]byte)
	if !ok {
		return errors.New("JsonString must be []byte")
	}
	var _t map[string]interface{}
	err := json.Unmarshal(b, &_t)
	if err != nil {
		return err
	}
	*t = _t
	return nil
}

func (t JsonObj) Value() (driver.Value, error) {
	// 放到数据库里面去
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}