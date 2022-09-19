package req

import (
	"errors"
	"fmt"
)

type Int64 struct {
	field string
	value int64
	Err   error
}

func NewInt64(field string, value int64) *Int64 {
	return &Int64{
		field: field,
		value: value,
		Err:   nil,
	}
}

func (i *Int64) GetField() string {
	return i.field
}

func (i *Int64) GetValue() int64 {
	return i.value
}

// value < min
func (r *Int64) Min(min int64) *Int64 {
	if r != nil && r.Err != nil {
		return r
	}
	if r.value < min {
		r.Err = errors.New(fmt.Sprintf("field:%s value:%f is less than %f", r.field, r.value, min))
	}
	return r
}

//value > max
func (r *Int64) Max(max int64) *Int64 {
	if r != nil && r.Err != nil {
		return r
	}
	if r.value > max {
		r.Err = errors.New(fmt.Sprintf("field:%s value:%f is more than %f", r.field, r.value, max))
	}
	return r
}
