package base32

import (
	"encoding/base32"
)

func Base32encoder(src []byte) []byte {
	return []byte(base32.StdEncoding.EncodeToString(src))
}
func Base32decoder(src []byte) []byte {
	dst, _ := base32.StdEncoding.DecodeString(string(Base32encoder(src)))
	return dst
}
