package fptf

import (
	"errors"
	"time"
)

func MustTime(t *time.Time) time.Time {
	if t == nil {
		panic(errors.New("time is nil"))
	}
	return *t
}

func MustInt(i *int) int {
	if i == nil {
		panic(errors.New("int is nil"))
	}
	return *i
}

func MustBool(b *bool) bool {
	if b == nil {
		panic(errors.New("bool is nil"))
	}
	return *b
}

func Time(t time.Time) *time.Time {
	return &t
}

func Int(i int) *int {
	return &i
}

func Delay(delay int) *int {
	return Int(delay)
}

func Bool(b bool) *bool {
	return &b
}
