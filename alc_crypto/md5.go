package alc_crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(psw string) string {
	newMd5 := md5.New()
	newMd5.Write([]byte(psw))
	return hex.EncodeToString(newMd5.Sum(nil))
}
