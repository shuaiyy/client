package conv

import (
	"fmt"
	"strconv"
)

// S string transfer
type S string

func (s S) String() string {
	return string(s)
}

// Int s to int
func (s S) Int() int {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(v)
}

// Int8 s to int8
func (s S) Int8() int8 {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	return int8(v)
}

// Int64  s to int
func (s S) Int64() int64 {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

// SS ss
type SS []string

// StrSlice ...
func (ss SS) StrSlice() []string {
	return ss
}

// IntSlice ...
func (ss SS) IntSlice() []int {
	var res []int
	for _, s := range ss {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, int(v))
	}
	return res
}

// Int64Slice ...
func (ss SS) Int64Slice() []int64 {
	var res []int64
	for _, s := range ss {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		res = append(res, v)
	}
	return res
}

// I ...
func I(v interface{}) S {
	if v, ok := v.(string); ok {
		return S(v)
	}
	return S(fmt.Sprintf("%v", v))
}
