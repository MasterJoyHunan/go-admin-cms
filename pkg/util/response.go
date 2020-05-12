// 统一返回
package util

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Response struct {
	C *gin.Context
}

func (g *Response) Send(code int, message string, data interface{}) {
	g.C.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  message,
	})
}

func (g *Response) SendMsg(code int, message string) {
	g.C.JSON(200, gin.H{
		"code": code,
		"data": "",
		"msg":  message,
	})
}

func (g *Response) SendData(code int, data interface{}) {
	g.C.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  "",
	})
}

func (g *Response) SuccessMsg(message string) {
	g.C.JSON(200, gin.H{
		"code": 1,
		"data": "",
		"msg":  message,
	})
}

func (g *Response) ErrorMsg(message string) {
	g.C.JSON(200, gin.H{
		"code": 0,
		"data": "",
		"msg":  message,
	})
}

func (g *Response) SuccessData(data interface{}) {
	g.C.JSON(200, gin.H{
		"code": 1,
		"data": data,
		"msg":  "",
	})
}

func (g *Response) ErrorData(data interface{}) {
	g.C.JSON(200, gin.H{
		"code": 0,
		"data": data,
		"msg":  "",
	})
}

func (g *Response) Error(msg ...string) {
	if len(msg) == 0 {
		msg = append(msg, "系统错误")
	}
	g.C.JSON(500, gin.H{
		"code": 500,
		"data": "",
		"msg":  strings.Join(msg, ""),
	})
}

// 验证错误
func (g *Response) SetValidateError(err error) *gin.Error {
	return g.C.Error(err).SetType(gin.ErrorTypeBind)
}

// 数据库异常
func (g *Response) SetOtherError(err error) *gin.Error {
	return g.C.Error(err).SetType(gin.ErrorTypePublic)
}
