package model

import "strconv"

type CasbinRule struct {
	PType string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

// 获取所有用户和角色对应表
// [用户ID] => [角色ID_1, 角色ID_2, 角色ID_3]
func GetUserRoleMapping() map[int][]int {
	var all []CasbinRule
	Db.Select([]string{"substring(v0, 6) as v0", "substring(v1, 6) as v1"}).
		Where("p_type = 'g' and v0 like ?",  "user:%").
		Find(&all)
	hash := make(map[int][]int)
	for _, v := range all {
		userId, _ := strconv.Atoi(v.V0)
		roleId, _ := strconv.Atoi(v.V1)
		_, ok := hash[userId]
		if ok {
			hash[userId] = append(hash[userId], roleId)
		} else {
			hash[userId] = []int{roleId}
		}
	}
	return hash
}

// 根据条件查数据
func GetCasbinByWhere(where ...interface{}) (res CasbinRule, err error) {
	err = Db.First(&res, where...).Error
	return
}
