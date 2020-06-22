package middleware

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		// TODO 实际上线请修改这里
		if setting.ApplicationConf.Env == "release" {
			if c.Request.Method != "GET" && strings.Index(c.FullPath(), "login") == -1 {
				c.AbortWithStatusJSON(200, gin.H{
					"code": 0,
					"data": "",
					"msg":  "仅供演示,请勿操作",
				})
				return
			}
		}

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Logger.Infof("[GIN]| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
