package controllers

import (
	"fmt"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

var (
	// 设置上传到的空间
	bucket = "testwd"
	// 设置上传文件的key
	// key = "yourdefinekey"
	// 设置CallbackUrl字段
	callbackurl = "http://your.domain.com/qiniucallback"
	// 设置CallbackBody字段
	callbackbody = `{"key":"$(key)", "hash":"$(etag)","filesize":$(fsize)}`
)

// 构造返回值字段
type PutRet struct {
	Hash     string `json:"hash"`
	Key      string `json:"key"`
	FileSize int    `json:"filesize"`
}

func UpQiNiu(filepath, key string) PutRet {
	// 初始化AK，SK
	conf.ACCESS_KEY = "Wbm4-DTJs5TF3bRrfyOoOqNz7-RlpzWkXqxwxKil"
	conf.SECRET_KEY = "JRJGbaTmpZGTXXM-dt05MgzxYHVoT6xL_poQKmqe"
	// 创建一个Client
	c := kodo.New(0, nil)
	// 设置上传的策略
	policy := &kodo.PutPolicy{
		Scope: bucket + ":" + key,
		//设置Token过期时间
		Expires:      3600,
		CallbackUrl:  callbackurl,
		CallbackBody: callbackbody,
	}
	// 生成一个上传token
	token := c.MakeUptoken(policy)
	// 构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	// 设置上传文件的路径
	// filepath := "/xxx/xxx/sample.flv"
	// 调用PutFile方式上传，这里的key需要和上传指定的key一致
	res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
	// 打印返回的信息
	// 打印出错信息
	if res != nil {
		fmt.Println("io.Put failed:", res)
		return ret
	}
	return ret
}
