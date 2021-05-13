// 权限一般都是程序员添加的，无需写增删改的逻辑
// 只需查的逻辑
package controller

import (
	"blog/pkg/util"
	"blog/service"
	"github.com/gin-gonic/gin"
)

// @Tags 角色
// @Summary 获取权限树
// @Produce json
// @Success 200 {object} response.RolePage
// @Router /auth/tree [get]
func AuthTree(c *gin.Context) {
	response := util.Response{c}
	res := service.AuthTreeCache()
	response.SuccessData(res)
}