package controllers

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	// "strconv"
	"strings"
	"time"
)

const (
	secret      = "i9IkzpKWCuSwK808iNR6awWREjsvU4"
	AccessKeyId = "LTAIVF29fBCXjmYC"
)

// type Aliyun struct {
// 	RequestId string
// 	HostId    string
// 	Code      string
// 	Message   string
// }

var Aliyun map[string]interface{}

func PercentEncode(str string) string {
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)
	return str
}

func CreateSignature(secret, StringToSign string) string {
	//创建签名
	key := []byte(secret + "&")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(StringToSign))
	s := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return s
}

func RandSeq() int {

	return int(time.Now().UnixNano() / 1000000)
}

func GetUtcTime() string {
	now := time.Now()
	year, mon, day := now.UTC().Date()
	hour, min, sec := now.UTC().Clock()
	s := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02dZ", year, mon, day, hour, min, sec)
	return s
}

func SendMail(toaddress, htmlbody, subject string) *map[string]interface{} {
	//组织请求数据，方法get
	urldata := url.Values{}
	//获取utc时间
	utc := GetUtcTime()
	// randstr := strconv.Itoa(RandSeq())

	urldata.Add("Action", "SingleSendMail")
	urldata.Add("AccountName", "admin@testwd.cn")
	urldata.Add("ReplyToAddress", "true")
	urldata.Add("AddressType", "1")
	urldata.Add("ToAddress", toaddress)
	urldata.Add("FromAlias", "测试村")
	urldata.Add("Subject", subject)
	urldata.Add("HtmlBody", htmlbody)
	// urldata.Add("TextBody", "")
	urldata.Add("Format", "JSON")
	urldata.Add("Version", "2015-11-23")
	urldata.Add("SignatureMethod", "HMAC-SHA1")
	urldata.Add("SignatureNonce", utc)
	urldata.Add("SignatureVersion", "1.0")
	urldata.Add("AccessKeyId", AccessKeyId)
	urldata.Add("Timestamp", utc)
	fmt.Println(urldata.Encode())
	//字符替换
	percent := PercentEncode(urldata.Encode())
	//获取StringToSign
	StringToSign := "GET" + "&" + url.QueryEscape("/") + "&" + url.QueryEscape(percent)

	Signature := CreateSignature(secret, StringToSign)
	urldata.Add("Signature", Signature)

	res, err := http.Get("https://dm.aliyuncs.com/?" + PercentEncode(urldata.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &Aliyun)
	return &Aliyun
}
