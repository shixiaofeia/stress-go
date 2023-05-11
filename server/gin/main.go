package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	go func() {
		var m runtime.MemStats
		for {
			runtime.ReadMemStats(&m)
			println(m.Sys / 1024 / 1024)
			time.Sleep(1 * time.Second)
		}
	}()

	r.GET("/download", func(c *gin.Context) {

		filePath := "/Users/xiaofei/Downloads/Docker.dmg"
		fileName := "Docker.dmg"

		file, err := os.Open(filePath)
		if err != nil {
			c.String(404, "文件不存在")
			return
		}
		defer file.Close()

		// 设置响应头
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/octet-stream")

		// 将文件内容复制到响应体中
		_, err = io.Copy(c.Writer, file)
		if err != nil {
			c.String(500, "文件下载失败")
			return
		}
	})

	_ = r.Run("0.0.0.0:8080")
}
