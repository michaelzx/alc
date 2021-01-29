package alc_types

import (
	"bytes"
	"database/sql/driver"
	"strings"
)

var (
	bitTrue  = []byte{0x1}
	bitFalse = []byte{0x0}
)

type Bool bool

func (b *Bool) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		if bytes.Compare(src.([]byte), bitTrue) == 0 {
			*b = true
		} else {
			*b = false
		}
	}
	return nil
}
func (b Bool) Value() (driver.Value, error) {
	if !b {
		return bitFalse, nil
	}
	return bitTrue, nil
}
func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte(`true`), nil
	} else {
		return []byte(`false`), nil
	}
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if strings.ToLower(dataStr) == "true" {
		*b = true
	} else {
		*b = false
	}
	return nil
}
