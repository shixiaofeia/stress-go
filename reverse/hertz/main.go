package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/hertz-contrib/reverseproxy"
)

const addr = "http://192.168.3.30:8080"

func main() {
	h := server.Default(
		server.WithHostPorts("0.0.0.0:8081"),
		server.WithTransport(standard.NewTransporter),
	)
	proxy, err := reverseproxy.NewSingleHostReverseProxy(addr)
	if err != nil {
		panic(err)
	}
	h.NoRoute(proxy.ServeHTTP)
	h.Spin()
}
