package shorturl

import (
	"crypto/md5"
	"fmt"
)

func Shorten(in string) string {
	data := []byte(in)
	sum := md5.Sum(data)
	md5str := fmt.Sprintf("%x", sum)
	return md5str[0:6]
}
