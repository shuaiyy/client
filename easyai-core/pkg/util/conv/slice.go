package conv

import (
	"strconv"
	"strings"
)

// UniqueStrSlice unique
func UniqueStrSlice(a, b []string, c ...string) []string {
	res := make([]string, 0, len(a)+len(b)+len(c))
	seen := make(map[string]bool)
	for _, k := range a {
		if seen[k] {
			continue
		}
		res = append(res, k)
		seen[k] = true
	}

	for _, k := range b {
		if seen[k] {
			continue
		}
		res = append(res, k)
		seen[k] = true
	}
	for _, k := range c {
		if seen[k] {
			continue
		}
		res = append(res, k)
		seen[k] = true
	}
	return res
}

// ParseStringSliceToUint64 ...
func ParseStringSliceToUint64(s []string) []uint64 {
	iv := make([]uint64, len(s))
	for i, v := range s {
		iv[i], _ = strconv.ParseUint(v, 10, 64)
	}
	return iv
}

// ParseStringSlice ...
func ParseStringSlice(s string, seps ...string) []string {
	sep := ","
	if len(seps) > 0 {
		sep = seps[0]
	}
	var res []string
	for _, v := range strings.Split(s, sep) {
		if v = strings.TrimSpace(v); v != "" {
			res = append(res, v)
		}
	}
	return res
}

// NewStrSlice got a new slice with given fn mapped
func NewStrSlice(ss []string, fn func(string) string) []string {
	res := make([]string, len(ss))
	for i, s := range ss {
		res[i] = fn(s)
	}
	return res
}

// StrSliceContains true if contains
func StrSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// StrSliceDelete delete element
func StrSliceDelete(old []string, ss ...string) []string {
	if len(ss) == 0 {
		return old
	}
	res := make([]string, 0, len(old))
	for _, item := range old {
		ignore := false
		for _, s := range ss {
			if item == s {
				ignore = true
				break
			}
		}
		if !ignore {
			res = append(res, item)
		}
	}
	return res
}
