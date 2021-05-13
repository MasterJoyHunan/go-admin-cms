package controller

import (
	"blog/model"
	"blog/pkg/util"
	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	response := util.Response{c}
	user, hasUser := c.Get("user")
	if !hasUser {
		response.Error("用户未登录")
		return
	}
	userInfo := user.(*util.Claims)
	res := make(map[string]interface{}, 2)
	res["role"] = model.GetUserRole(userInfo.Id)
	res["auth"] = model.GetUserAuth(userInfo.Id)
	response.SuccessData(res)
}
