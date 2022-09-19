package req

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Str struct {
	field string
	value string
	Err   error
}

func NewStr(field, value string) *Str {
	return &Str{
		field: field,
		value: value,
	}
}

func (r *Str) GetField() string {
	return r.field
}

func (r *Str) GetValue() string {
	return r.value
}

func (r *Str) Empty() *Str {
	if r != nil && r.Err != nil {
		return r
	}
	if strings.Trim(r.value, "") == "" {
		r.Err = errors.New(fmt.Sprintf("field:%s is empty", r.field))
	}
	return r
}

func (r *Str) In(m map[string]struct{}) *Str {
	if r != nil && r.Err != nil {
		return r
	}
	if _, isOk := m[r.value]; !isOk {
		strs := make([]string, 0, len(m))
		for k, _ := range m {
			strs = append(strs, k)
		}
		r.Err = errors.New(fmt.Sprintf("field:%s value not in %v", r.field, strs))
	}
	return r
}

func (r *Str) Len(min, max int) *Str {
	if r != nil && r.Err != nil {
		return r
	}
	if len(r.value) > max {
		r.Err = errors.New(fmt.Sprintf("field:%s len is more than %d", r.field, max))
	}
	if len(r.value) < min {
		r.Err = errors.New(fmt.Sprintf("field:%s len is less than %d", r.field, min))
	}
	return r
}

func (r *Str) ConvertToFloat64() *Float64 {
	rf := NewFloat64(r.field, 0)
	if r != nil && r.Err != nil {
		rf.Err = r.Err
		return rf
	}
	f, err := strconv.ParseFloat(r.value, 64)
	if err != nil {
		rf.Err = errors.New(fmt.Sprintf("field:%s,value:%v can't convert to float64 ", rf.field, r.value))
	}
	rf.value = f
	return rf
}

func (r *Str) ConvertToTime() *Time {
	rf := NewTime(r.field, nil)
	if r != nil && r.Err != nil {
		rf.Err = r.Err
		return rf
	}
	t, err := time.Parse("2006-01-02 15:04:05", r.GetValue())
	if err != nil {
		rf.Err = errors.New(fmt.Sprintf(`field:%s,value:%v can't convert to Datetime,please use format 2006-01-02 15:04:05`, rf.field, r.value))
	}
	rf.value = &t
	return rf
}

func (r *Str) ConvertToInt64() *Int64 {
	rf := NewInt64(r.field, 0)
	if r != nil && r.Err != nil {
		rf.Err = r.Err
		return rf
	}
	f, err := strconv.ParseInt(r.value, 10, 64)
	if err != nil {
		rf.Err = errors.New(fmt.Sprintf("field:%s,value:%v can't convert to int64 ", rf.field, r.value))
	}
	rf.value = f
	return rf
}
