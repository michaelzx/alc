package alc_types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Time time.Time

const DefaultTimeFormat = "2006-01-02 15:04:05"

var TimeZone = time.FixedZone("CST", 8*3600)

func NewTime(t time.Time) Time {
	return Time(t)
}

func TimeNow() Time {
	return Time(time.Now())
}

// 转换成系统标准库time.Time类型
func (t Time) StdTime() time.Time {
	return time.Time(t)
}

func (t *Time) Scan(value interface{}) error {
	if stdTime, ok := value.(time.Time); ok {
		*t = NewTime(stdTime)
	}
	return nil
}

func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t Time) Format(timeFormat string) string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(DefaultTimeFormat))
}

func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return errors.New("null is not allow in alc_types.Time")
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	stdTime, err := time.ParseInLocation(`"`+DefaultTimeFormat+`"`, string(data), TimeZone)
	*t = NewTime(stdTime)
	return err
}
