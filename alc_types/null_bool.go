package alc_types

import (
	"database/sql"
	"encoding/json"
)

// 弥补sql.NullBool无法解析json null的缺陷
type NullBool sql.NullBool

func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
	var x *bool
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		b.Valid = true
		b.Bool = *x
	} else {
		b.Valid = false
	}
	return nil
}
