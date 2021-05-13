package middleware

import (
	"blog/myerr"
	"blog/pkg/logger"
	"blog/validate"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

// 错误处理
func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		if len(c.Errors) == 0  {
			return
		}

		// 参数验证错误
		if err := c.Errors.ByType(gin.ErrorTypeBind).Last(); err != nil {
			if validateErrs, ok := err.Err.(validator.ValidationErrors); ok {
				for _, e := range validateErrs {
					zhErrorMessage := e.Translate(validate.Trans)
					baseErrorMessage := validate.ErrorTitleMessage

					// 自定义错误提示
					if coverMessage, ok := err.Meta.(map[string]string); ok {
						for k, v := range coverMessage {
							baseErrorMessage[k] = v
						}
					}

					// 替换错误提示
					if res, ok := baseErrorMessage[e.Field()]; ok {
						zhErrorMessage = strings.Replace(zhErrorMessage, e.Field(), res, 1)
					}
					c.AbortWithStatusJSON(200, gin.H{
						"code": 0,
						"data": "",
						"msg":  zhErrorMessage,
					})
					return
				}
			}
			if unmarshalErrs, ok := err.Err.(*json.UnmarshalTypeError); ok {
				c.AbortWithStatusJSON(200, gin.H{
					"code": 0,
					"data": "",
					"msg":  "传参[ " + unmarshalErrs.Struct + " ]解析字段[ " + unmarshalErrs.Field + " ]异常: [ " + unmarshalErrs.Value + " ]",
				})
				logger.Logger.Error("传参解析异常", unmarshalErrs.Error())
				return
			}
		}

		// 其他异常
		if err := c.Errors.ByType(gin.ErrorTypePublic).Last(); err != nil {
			if normalValidateErrs, ok := err.Err.(*myerr.NormalValidateError); ok {
				// 普通验证错误
				c.AbortWithStatusJSON(200, gin.H{
					"code": 0,
					"data": "",
					"msg":  normalValidateErrs.Error(),
				})
			} else if dbValidateErrs, ok := err.Err.(*myerr.DbValidateError); ok {
				// 数据库验证错误
				c.AbortWithStatusJSON(200, gin.H{
					"code": 0,
					"data": "",
					"msg":  dbValidateErrs.Error(),
				})
			} else if gorm.IsRecordNotFoundError(err.Err) {
				// 数据未找到
				c.AbortWithStatusJSON(200, gin.H{
					"code": 0,
					"data": "",
					"msg":  "数据不存在",
				})
			}
			return
		}

		// 其他未知错误
		c.AbortWithStatusJSON(500, gin.H{
			"code": 0,
			"data": "",
			"msg":  c.Errors.Last().Error(),
		})
		logger.Logger.Error(c.Errors.Last().Error())
	}
}
