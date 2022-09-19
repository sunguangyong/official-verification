package req

import (
	"errors"
	"fmt"
	"strings"
)

type Float64 struct {
	field string
	value float64
	Err   error
}

func NewFloat64(field string, value float64) *Float64 {
	return &Float64{
		field: field,
		value: value,
		Err:   nil,
	}
}
func (r *Float64) GetField() string {
	return r.field
}
func (r *Float64) GetValue() float64 {
	return r.value
}

// value < min
func (r *Float64) Min(min float64) *Float64 {
	if r != nil && r.Err != nil {
		return r
	}
	if r.value < min {
		r.Err = errors.New(fmt.Sprintf("field:%s value:%f is less than %f", r.field, r.value, min))
	}
	return r
}

//value > max
func (r *Float64) Max(max float64) *Float64 {
	if r != nil && r.Err != nil {
		return r
	}
	if r.value > max {
		r.Err = errors.New(fmt.Sprintf("field:%s value:%f is more than %f", r.field, r.value, max))
	}
	return r
}

//value 小数位长度 >j的小数位长度
func (r *Float64) DPGreaterThan(j float64) *Float64 {
	if r != nil && r.Err != nil {
		return r
	}
	iStr := fmt.Sprint(r.value)
	jStr := fmt.Sprint(j)
	iStrArr := strings.Split(iStr, ".")
	jStrArr := strings.Split(jStr, ".")
	if len(iStrArr) > 1 && len(jStrArr) > 1 {
		if len(iStrArr[1]) > len(jStrArr[1]) {
			r.Err = errors.New(fmt.Sprintf("field:%s,value:%v The length of the decimal places is greater than:%v", r.field, r.value, j))
		}
	}
	return r
}
