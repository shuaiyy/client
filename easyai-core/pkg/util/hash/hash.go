package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/big"
)

// MD5 MD5哈希值
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5哈希值
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 SHA1哈希值
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1哈希值
func SHA1String(s string) string {
	return SHA1([]byte(s))
}

// Md5ToBigInt string hash to big int
func Md5ToBigInt(s string) big.Int {
	s = MD5String(s)
	bi := big.Int{}
	bi.SetString(s, 16)
	return bi
}
