package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
)

func PercentEncode(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
}

func CreateSignature(secret, StringToSign string) string {
	key := []byte(secret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(data))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s
}

func GetUtcTime() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)
	return s
}

func Body(toaddress, htmlbody, subject string) {
	urldata := url.Values{}
	utc := GetUtcTime()
	urldata.Add(key, value)
	urldata.Add("Action", "SingleSendMail")
	urldata.Add("AccountName", "admin@testwd.cn")
	urldata.Add("ReplyToAddress", "true")
	urldata.Add("AddressType", "1")
	urldata.Add("ToAddress", toaddress)
	urldata.Add("FromAlias", "测试村<admin@testwd.cn>")
	urldata.Add("Subject", subject)
	urldata.Add("HtmlBody", htmlbody)
	// urldata.Add("TextBody", "")
	urldata.Add("Format", "JSON")
	urldata.Add("Version", "2015-11-23")
	urldata.Add("SignatureMethod", "HMAC-SHA1")
	urldata.Add("SignatureNonce", randstr)
	urldata.Add("SignatureVersion", "1.0")
	urldata.Add("AccessKeyId", "LTAIjsEviSFAO1wp")
	urldata.Add("Timestamp", utc)
}
