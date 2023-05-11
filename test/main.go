package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 创建一个 TLS 配置对象，用于配置 TLS 连接的参数
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证，仅用于演示，实际使用中应该关闭
	}

	// 创建一个自定义的 HTTP 客户端，使用自定义的 TLS 配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// 发送一个 HTTPS GET 请求
	resp, err := client.Get("https://127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(body))
}
