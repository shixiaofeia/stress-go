package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const addr = "192.168.3.30:8080"

func main() {
	// 创建反向代理对象
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   addr,
	})

	proxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,

		Dial:                   nil,
		DialTLSContext:         nil,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    10 * time.Second,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           0,
		MaxIdleConnsPerHost:    1000, // 最大空闲连接, 最好为最大活跃连接的一半或偏小
		MaxConnsPerHost:        2000, // 最大活跃连接
		IdleConnTimeout:        90 * time.Second,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  time.Second,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      true,
	}

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    ":8081",
		Handler: proxy,
	}

	// 启动服务器
	log.Printf("Starting server on :%d\n", 8000)
	log.Fatal(server.ListenAndServe())
}
