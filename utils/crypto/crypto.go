package crypto

import (
	"crypto/md5"
	"encoding/hex"
)
func Passwordhash(pass string)string{
	hasher := md5.New()
	defer hasher.Reset()
	hasher.Write([]byte(pass))
	return hex.EncodeToString(hasher.Sum(nil))
}