package model

import "blog/model/response"

type Auth struct {
	Id     int
	Pid    int    `gorm:"default:0;comment:'上级ID';"`
	Name   string `gorm:"default:'';comment:'节点名';"`
	IsMenu int8   `gorm:"default:0;comment:'是否是菜单栏 0：否，1：是';"`
	Api    string `gorm:"default:'';comment:'接口';"`
	Action string `gorm:"default:'';comment:'操作方法';"`
	Ext    string `gorm:"unique;default:'';comment:'前端使用';"`
}

// 查询所有权限
func GetAllAuth(where ...interface{}) (res []response.Auth) {
	Db.Model(Auth{}).Find(&res, where...)
	return
}

// 查询所有权限
func GetAllBaseAuth(where ...interface{}) (res []Auth) {
	Db.Model(Auth{}).
		Order("id asc").
		Find(&res, where...)
	return
}