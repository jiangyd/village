package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
)

func Getuuid() string {
	uuid := fmt.Sprintf("%s", uuid.NewV4())
	return uuid
}

func Encrypt(data string) string {
	key := []byte("$#@%$^GFDER00DC#24$")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))

	return fmt.Sprintf("%x", mac.Sum(nil))
}
