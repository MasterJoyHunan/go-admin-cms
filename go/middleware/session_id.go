package middleware

import (
	"blog/pkg/setting"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const SessionIdName = "GO-SESSION-ID"

// Go语言没有自带session机制，我们手动实现一个
func MakeSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request
		sessionId, _ := c.Cookie(SessionIdName)
		if sessionId == "" {
			sessionId = uuid.NewV4().String()
			c.SetCookie(SessionIdName, sessionId, 0, "/", setting.ApplicationConf.Doamin, false, true)
		}
		c.Set("session_id", sessionId)
		c.Next()
	}
}
