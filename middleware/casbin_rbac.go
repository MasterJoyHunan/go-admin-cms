package middleware

import (
	"blog/pkg/casbin"
	"blog/pkg/logger"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// 鉴权
func Authentication() gin.HandlerFunc {

	return func(c *gin.Context) {
		user, hasUser := c.Get("user")
		if !hasUser {
			return
		}
		userInfo := user.(*util.Claims)
		userIdStr := "user:" + strconv.Itoa(userInfo.Id)
		e, err := casbin.InitCasbin()
		if err != nil {
			logger.Logger.Panic("初始化 Casbin 出现错误：", err)
		}
		ok, err := e.Enforce(userIdStr, c.FullPath(), strings.ToLower(c.Request.Method))
		if err != nil {
			logger.Logger.Panic("执行 e.Enforce() 出错：", err)
		}
		if !ok {
			c.JSON(403, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
