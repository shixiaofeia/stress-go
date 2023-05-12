package main

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

const addr = "192.168.3.30:8080"

type BytesBufferPool struct {
	defaultSize int
	pool        sync.Pool
}

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
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConnsPerHost:   1000, // 最大空闲连接, 最好为最大活跃连接的一半或偏小
		MaxConnsPerHost:       2000, // 最大活跃连接
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: time.Second,
		ForceAttemptHTTP2:     true,
	}

	// 使用缓冲池可以减少内存分配, 有效提升性能
	proxy.BufferPool = NewBytesBufferPool(64 * 1024)

	// 创建 HTTP 服务器
	server := &http.Server{
		Addr:    ":8081",
		Handler: proxy,
	}

	// 启动服务器
	log.Printf("Starting server on :%d\n", 8000)
	log.Fatal(server.ListenAndServe())
}

// NewBytesBufferPool 字节缓冲池.
func NewBytesBufferPool(defaultSize int) *BytesBufferPool {
	p := &BytesBufferPool{
		defaultSize: defaultSize,
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, defaultSize)
			},
		},
	}
	return p
}

func (p *BytesBufferPool) Get() []byte {
	return p.pool.Get().([]byte)
}

func (p *BytesBufferPool) Put(bs []byte) {
	p.pool.Put(bs)
}
