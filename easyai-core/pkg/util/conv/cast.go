package conv

import "github.com/spf13/cast"

// ToInt8  i to int8
func ToInt8(v interface{}) int8 {
	return cast.ToInt8(v)
}
