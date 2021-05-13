package middleware

import (
	"github.com/gin-gonic/gin"
)

// 处理 404 -- 如果没有命中路由直接返回 404 状态码
// gin 官方确实有 NoFound 方法，但是还是会走很多的中间件
func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Status() == 404 {
			c.AbortWithStatus(404)
			return
		}
	}
}
