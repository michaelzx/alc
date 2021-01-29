package alc_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type aesBase64Kit struct {
}

func NewAesKit() *aesBase64Kit {
	return &aesBase64Kit{}
}

var AesBase64Kit = NewAesKit()

func (a *aesBase64Kit) EncryptBase64(src []byte, key []byte) string {
	eb := a.encryptAES(src, key)
	return a.encodeBase64(eb)
}
func (a *aesBase64Kit) DecryptBase64(src string, key []byte) ([]byte, error) {
	eb, err := a.decodeBase64(src)
	if err != nil {
		return nil, err
	}
	return a.decryptAES(eb, key)
}
func (a *aesBase64Kit) encodeBase64(b []byte) string {
	return base64.URLEncoding.EncodeToString(b)
}

func (a *aesBase64Kit) decodeBase64(s string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(s)
}

func (a *aesBase64Kit) padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func (a *aesBase64Kit) unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func (a *aesBase64Kit) encryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = a.padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func (a *aesBase64Kit) decryptAES(src []byte, key []byte) (res []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockmode := cipher.NewCBCDecrypter(block, key)
	defer func() {
		if info := recover(); info != nil {
			if v, ok := info.(error); ok {
				res = nil
				err = v
			}
		}
	}()
	blockmode.CryptBlocks(src, src)
	src = a.unpadding(src)
	return src, nil
}
