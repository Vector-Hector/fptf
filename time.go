package fptf

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"strconv"
	"time"
)

// TimeUnix a time.Time object, that is (un)marshalled as unix integers
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

func (t *TimeUnix) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	var i int64
	err := bson.UnmarshalValue(typ, data, &i)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(i, 0)
	return nil
}

func (t TimeUnix) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(t).Unix())
}

// Is essentially just a time.Time object, but zero values are marshalled as null
type TimeNullable struct {
	time.Time
}

// MarshalJSON is used to convert the timestamp to JSON
func (t TimeNullable) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

func (t *TimeNullable) UnmarshalBSONValue(typ bsontype.Type, data []byte) error {
	if typ == bson.TypeNull {
		return nil
	}
	return bson.UnmarshalValue(typ, data, &t.Time)
}

func (t TimeNullable) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if t.IsZero() {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(t.Time)
}
