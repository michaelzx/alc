package alc_types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date time.Time

const DefaultDateFormat = "2006-01-02"

func (j *Date) Scan(value interface{}) error {
	if stdTime, ok := value.(time.Time); ok {
		*j = Date(stdTime)
	}
	return nil
}
func (j Date) Value() (driver.Value, error) {
	return time.Time(j), nil
}
func (j *Date) UnmarshalJSON(bytes []byte) error {
	t, err := time.Parse(`"`+DefaultDateFormat+`"`, string(bytes))
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

func (j Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(j).Format(DefaultDateFormat))
	return []byte(stamp), nil
}
func (j Date) StdTime() time.Time {
	return time.Time(j)
}
