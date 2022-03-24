package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func main() {
	u, err := url.Parse("https://coolcar-1309863650.cos.ap-nanjing.myqcloud.com")
	if err != nil {
		panic(err)
	}
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	secretID := "AKID2tUpthFWWLuCCxRjgPr2wZv0oLbbKnzu"
	secretKey := "g6F9m8D2arqY1C7mAI9k47sNbaIjtoSp"
	// 1.永久密钥
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: secretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	// 获取预签名URL
	presignedURL, err := client.Object.GetPresignedURL(
		context.Background(), http.MethodGet, "test123.png", secretID, secretKey, 20*time.Second, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL)

}
