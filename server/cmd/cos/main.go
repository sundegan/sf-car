package main

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
)

func main() {
	u, _ := url.Parse("https://sfcar-1304689777.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{

			SecretID: "AKIDoZG0c3MYeYFj2A97zhGfy98fAXqzveEc",

			SecretKey: "8B5oJar22pbGptlT0NaaVK2kKo9KLGz6",
		},
	})

	ak := "AKIDoZG0c3MYeYFj2A97zhGfy98fAXqzveEc"
	sk := "8B5oJar22pbGptlT0NaaVK2kKo9KLGz6"

	// 获取上传预签名 URL
	name1 := "test.txt"
	ctx1 := context.Background()
	presignedURL_PUT, err := client.Object.GetPresignedURL(ctx1, http.MethodPut, name1, ak, sk, time.Hour, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL_PUT)

	name2 := "test.txt"
	ctx2 := context.Background()
	// 获取下载预签名 URL
	presignedURL_GET, err := client.Object.GetPresignedURL(ctx2, http.MethodGet, name2, ak, sk, time.Hour, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(presignedURL_GET)
}
