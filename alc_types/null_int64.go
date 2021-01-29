package alc_types

import (
	"database/sql"
	"database/sql/driver"
	"strconv"
)

type NullInt64 sql.NullInt64

func (n *NullInt64) Scan(value interface{}) error {
	nv := (*sql.NullInt64)(n)
	return nv.Scan(value)
}
func (n NullInt64) Value() (driver.Value, error) {
	nv := sql.NullInt64(n)
	return nv.Value()
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(n.Int64, 10)), nil
}

func (n *NullInt64) UnmarshalJSON(data []byte) error {
	dataStr := string(data)
	if dataStr == "null" {
		return nil
	}
	if i64, err := strconv.ParseInt(dataStr, 10, 64); err != nil {
		return err
	} else {
		n.Int64 = i64
		n.Valid = true
		return nil
	}
}
