package common

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type DateTime time.Time

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt == nil {
		return []byte("null"), nil
	}
	return []byte(`"` + time.Time(*dt).Format("2006-01-02 15:04:05") + `"`), nil
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*dt = DateTime(t)
	return nil
}

func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid type for DateTime")
	}
	*dt = DateTime(t)
	return nil
}

func (dt DateTime) Value() (driver.Value, error) {
	return time.Time(dt), nil
}
