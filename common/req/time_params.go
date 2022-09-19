package req

import "time"

type Time struct {
	field string
	value *time.Time
	Err   error
}

func NewTime(field string, value *time.Time) *Time {
	return &Time{
		field: field,
		value: value,
		Err:   nil,
	}
}

func (i *Time) GetField() string {
	return i.field
}

func (i *Time) GetValue() *time.Time {
	return i.value
}
