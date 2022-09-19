package sign

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func Sign(key, data string) string {
	hmac := hmac.New(sha512.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}
