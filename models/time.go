package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (t *DateTime) UnmarshalJSON(data []byte) error {
	var err error

	t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}

	return nil
}

func (t *DateTime) Scan(src interface{}) error {
	if src == nil {
		t.Time = time.Time{}
		return nil
	}
	switch v := src.(type) {
	case time.Time:
		t.Time = v
	case string:
		t.Time, _ = time.Parse("2006-01-02 15:04:05", v)
	case []byte:
		t.Time, _ = time.Parse("2006-01-02 15:04:05", string(v))
	}
	return nil
}

func (t *DateTime) Value() (driver.Value, error) {
	return t.Time.Format("2006-01-02 15:04:05"), nil
}
