package middleware

import (
	"blog/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const JwtName = "Authorization"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"data": "",
				"msg":  JwtName + "不存在",
			})
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"data": "",
				"msg":  JwtName + "验证错误",
			})
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"data": "",
				"msg":  "登录过期",
			})
			return
		}

		// 设置用户对象在上下文中，方便后续使用
		c.Set("user", claims)
		c.Next()
	}
}

// 各种方法获取 token
// 为了防范 CSRF 攻击,不获取 query 和 from 里的 token
func getToken(c *gin.Context) (string, error) {
	if token := c.GetHeader(JwtName); token != "" {
		return token, nil
	}

	if token, _ := c.Cookie(JwtName); token != "" {
		return token, nil
	}
	return "", errors.New("没有找到" + JwtName)
}
