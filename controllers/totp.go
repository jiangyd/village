package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	length = 16
)

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}

//密钥生成
func GetSecret() string {
	str := "234567abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

//动态6位密码
func Totp(secret string, ago int64) string {
	keynospaces := strings.Replace(secret, " ", "", -1)
	keynospacesupper := strings.ToUpper(keynospaces)
	key, err := base32.StdEncoding.DecodeString(keynospacesupper)
	if err != nil {
		fmt.Println(err)
	}
	epochsecond := time.Now().Unix()
	epochsecond -= ago //ago可以为0，也可以为30，这样可以应付2组密码的情况
	pwd := oneTimePassword(key, toBytes(epochsecond/30))

	secondsRemaining := 30 - (epochsecond % 30)
	//fmt.Sprintf("%06d (%d second(s) remaining)\n", pwd, secondsRemaining)
	fmt.Println(secondsRemaining) //这个secondsRemaining没有用到,只是打印下
	return fmt.Sprintf("%06d", pwd)

}

//二维码包含内容
func Getotpauth(name, secret, issuer string) string {
	otpauth := "otpauth://totp/" + "testwd" + ":" + name + "?secret=" + secret + "&issuer=" + issuer
	return otpauth
}
