package fptf

import (
	"strconv"
	"time"
)

// Is essentially just a time.Time object, but it is (un)marshalled as
// unix integers
type TimeUnix time.Time

// MarshalJSON is used to convert the timestamp to JSON
func (t TimeUnix) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *TimeUnix) UnmarshalJSON(s []byte) error {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// Is essentially just a time.Time object, but zero values are marshalled as null
type TimeNullable struct {
	time.Time
}

// MarshalJSON is used to convert the timestamp to JSON
func (t TimeNullable) MarshalJSON() ([]byte, error) {
	if t == (TimeNullable{}) {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}
