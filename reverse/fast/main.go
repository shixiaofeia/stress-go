package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

const addr = "192.168.3.30:8080"

func main() {
	// 创建一个FastHTTP客户端
	client := &fasthttp.HostClient{
		Addr:                addr,
		MaxConns:            1024,
		MaxIdleConnDuration: 5 * time.Minute,
		ReadTimeout:         10 * time.Second, // 作为客户端如果超时的话连接不会被释放, 会继续等待只不过返回值会被丢弃, 所以需要设置readTimeout来释放连接, 参考 https://www.cnblogs.com/xavier-yang/p/11788500.html
	}

	// 创建一个FastHTTP请求处理函数
	handler := func(ctx *fasthttp.RequestCtx) {

		// 复制客户端请求
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)
		ctx.Request.CopyTo(req)

		// 设置请求的URL
		req.SetRequestURI("http://" + addr + string(ctx.Path()))

		// 发送请求到目标服务器，并获取响应
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)
		if err := client.Do(req, resp); err != nil {
			//err = fmt.Errorf("error when proxying the request: %s", err.Error())
			//log.Println(err.Error())
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		// 将目标服务器的响应写回到客户端
		_, _ = resp.WriteTo(ctx)
	}

	// 启动FastHTTP服务器并监听端口
	if err := fasthttp.ListenAndServe(":8081", handler); err != nil {
		fmt.Printf("Error when starting the server: %s", err.Error())
	}
}
