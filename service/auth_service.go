package service

import (
	"blog/model"
	"blog/model/response"
)

var (
	// 缓存权限树结构
	authTreeCache []response.Auth

	// 缓存权限map结构
	authMapCache map[int]model.Auth
)

// 返回无极限分类方式的权限
func AuthTreeCache() []response.Auth {
	if len(authTreeCache) == 0 {
		authTreeCache = authTree(0)
	}
	return authTreeCache
}

func authTree(pid int) []response.Auth {
	res := model.GetAllAuth("pid = ?", pid)
	for i, v := range res {
		res[i].Children = authTree(v.Id)
	}
	return res
}

// 缓存权限
func AuthMapCache() map[int]model.Auth {
	if len(authMapCache) == 0 {
		authMapCache = make(map[int]model.Auth)
		base := model.GetAllBaseAuth()
		for _, v := range base {
			authMapCache[v.Id] = v
		}
	}
	return authMapCache
}
