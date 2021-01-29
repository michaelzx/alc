package alc_types

import (
	"encoding/json"
	"errors"
)

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
