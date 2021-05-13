package model

import (
	"blog/model/response"
	"strconv"
)

type Role struct {
	Id   int
	Pid  int    `gorm:"default:0;comment:'上级ID'"`
	Name string `gorm:"unique;default:'';comment:'角色名'"`
	Auth string `gorm:"type:text;comment:'权限ID'"`
}

// 获取所有角色
func (r *Role) GetAll(page, pageSize int) (response.RolePage, error) {
	all := response.RolePage{
		Total:       r.getCount(),
		PerPage:     page,
		CurrentPage: pageSize,
		Data:        []response.RoleList{},
	}
	offset := GetOffset(page, pageSize)
	scope := Db.Table("role").
		Select([]string{"role.id", "role.pid", "role.name", "p_role.name as parent_name"}).
		Joins("left join role as p_role on role.pid = p_role.id").
		Order("role.id desc")
	err := scope.Offset(offset).Limit(pageSize).Find(&all.Data).Error
	return all, err
}

func GetAllRole() (mapping map[int]response.CasRole) {
	var roles []response.CasRole
	mapping = make(map[int]response.CasRole)
	Db.Table("role").Select("id,name").Order("id desc").Find(&roles)
	for _, role := range roles {
		mapping[role.Id] = role
	}
	return
}

// 获取角色详情
func (r *Role) Detail(id ...int) (res response.RoleList, err error) {
	searchId := r.Id
	if len(id) > 0 {
		searchId = id[0]
	}
	err = Db.Table("role").
		Select([]string{"role.id", "role.pid", "role.name", "p_role.name as parent_name", "role.auth as auths"}).
		Joins("left join role as p_role on role.pid = p_role.id").
		Where("role.id = ?", searchId).
		First(&res).
		Error
	return
}

// 创建角色
func (r *Role) Create(all []Auth) error {
	// 执行事务处理
	tx := Db.Begin()
	if err := tx.Create(&r).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 同步更新 casbin
	for _, v := range all {
		if v.IsMenu != 0 {
			continue
		}
		if err := tx.Create(&CasbinRule{
			PType: "p",
			V0:    "role:" + strconv.Itoa(r.Id),
			V1:    v.Api,
			V2:    v.Action,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// 编辑角色
func (r *Role) Edit(all []Auth) error {
	// 执行事务处理
	tx := Db.Begin()

	// 1.更新角色表
	if err := tx.Model(&Role{Id: r.Id}).
		Updates(map[string]interface{}{
			"name": r.Name,
			"pid":  r.Pid,
			"auth": r.Auth,
		}).Error;
		err != nil {
		tx.Rollback()
		return err
	}

	// 2.删除casbin表
	if err := tx.Delete(CasbinRule{}, "p_type = 'p' and v0 = ?", "role:"+strconv.Itoa(r.Id)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 3.重新加入casbin表
	for _, v := range all {
		if v.IsMenu != 0 {
			continue
		}
		if err := tx.Create(&CasbinRule{
			PType: "p",
			V0:    "role:" + strconv.Itoa(r.Id),
			V1:    v.Api,
			V2:    v.Action,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// 删除角色
func (r *Role) Del() error {
	tx := Db.Begin()

	// 1.删除角色表
	if err := tx.Where("id = ?", r.Id).Delete(r).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2.删除casbin表
	key := "role:" + strconv.Itoa(r.Id)
	if err := tx.Where("p_type = 'g' and (v0 = ? or v1 = ?)", key, key).Delete(CasbinRule{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 获取角色条数
func (r *Role) getCount() (total int) {
	Db.Model(&r).Count(&total)
	return
}

// 根据条件获取角色
func GetRoleByWhere(where ...interface{}) (res Role, err error) {
	err = Db.First(&res, where...).Error
	return
}

// 根据条件获取多个角色
func GetRolesByWhere(where ...interface{}) (res []Role, err error) {
	err = Db.Find(&res, where...).Error
	return
}

// 根据条件获取多个角色
func GetRoleTreeByWhere(where ...interface{}) (res []response.Roles, err error) {
	err = Db.Table("role").Find(&res, where...).Error
	return
}
