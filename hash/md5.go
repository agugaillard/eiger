package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
