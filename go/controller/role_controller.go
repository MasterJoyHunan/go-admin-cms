package controller

import (
	"blog/model/request"
	"blog/pkg/util"
	"blog/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags 角色
// @Summary 获取所有角色
// @Produce json
// @Param page query int false "页码"
// @Param pageSize query int false "每页显示多少条"
// @Success 200 {object} response.RolePage
// @Router /role [get]
func RoleIndex(c *gin.Context) {
	response := util.Response{c}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	res, err := service.RoleList(page, pageSize)
	if err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessData(res)
}

// @Tags 角色
// @Summary 获取角色详情
// @Produce json
// @Param id path int true "角色ID"
// @Success 200 {object} response.RoleList
// @Router /role/{id} [get]
func RoleDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	response := util.Response{c}
	detail, err := service.RoleDetail(id)
	if err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessData(detail)
}

// @Tags 角色
// @Summary 添加角色
// @Produce json
// @Param body body request.RoleAdd true "角色"
// @Success 200
// @Router /role [post]
func RoleAdd(c *gin.Context) {
	response := util.Response{c}
	var role request.RoleAdd
	if err := c.ShouldBindJSON(&role); err != nil {
		response.SetValidateError(err).SetMeta(map[string]string{"Name": "角色名"})
		return
	}
	if err := service.RoleAdd(role); err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessMsg("添加成功")
}

// @Tags 角色
// @Summary 编辑角色
// @Produce json
// @Param id path int true "角色ID"
// @Param body body request.RoleEdit true "角色"
// @Router /role/{id} [put]
func RoleEdit(c *gin.Context) {
	response := util.Response{c}
	var role request.RoleEdit
	role.Id, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&role); err != nil {
		response.SetValidateError(err).SetMeta(map[string]string{"Name": "角色名"})
		return
	}
	if err := service.RoleEdit(role); err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessMsg("编辑成功")
}

// @Tags 角色
// @Summary 删除角色
// @Produce json
// @Param id path int true "角色ID"
// @Success 200
// @Router /role/{id} [delete]
func RoleDel(c *gin.Context) {
	response := util.Response{c}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.RoleDel(id); err != nil {
		response.SetOtherError(err)
		return
	}
	response.SuccessMsg("删除完成")
}
