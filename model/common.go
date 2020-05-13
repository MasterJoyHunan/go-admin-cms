package model

import (
	"blog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var Db *gorm.DB

type Base struct{}
type FieldTrans map[string]string

func Setup() {
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.MysqlConf.User,
		setting.MysqlConf.Pwd,
		setting.MysqlConf.Host,
		setting.MysqlConf.Port,
		setting.MysqlConf.Db))
	if err != nil {
		log.Panicf("连接数据库错误 ：%s", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.MysqlConf.Prefix + defaultTableName
	}

	Db.SingularTable(true)
	Db.LogMode(true)
	Db.SetLogger(&GormLogger{})
	Db.DB().SetMaxIdleConns(setting.MysqlConf.MaxIdle)
	Db.DB().SetMaxOpenConns(setting.MysqlConf.MaxActive)
	AutoMigrate()

	// 设置程序启动参数 -init | -init=true
	if setting.Init {
		InitSql()
	}
}

// 自动创建修改表
func AutoMigrate() {
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '后台用户'").AutoMigrate(&AdminUser{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '角色'").AutoMigrate(&Role{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '权限'").AutoMigrate(&Auth{})
	Db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'casbin policy 配置'").AutoMigrate(&CasbinRule{})
}

func InitSql() {
	// 清空
	Db.Exec("truncate admin_user")
	Db.Exec("truncate role")
	Db.Exec("truncate casbin_rule")
	Db.Exec("truncate auth")

	// 初始化
	Db.Exec("insert into admin_user (id, user_name, password, real_name, tel, status) values (1, 'admin', '$2a$10$057uuCLoKja2J04GLuWl1eNnwQtS7HxvookpbBa0thTHq7/fIaNF6', 'joy', '13054174174', 1)")

	Db.Exec("insert into role (id, pid, name, auth) values (1, 0, '超级管理员', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18')")
	Db.Exec("insert into role (id, pid, name, auth) values (2, 1, '系统维护管理员', '1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18')")

	Db.Exec("insert into casbin_rule (p_type, v0, v1) values ('g', 'user:1', 'role:1')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/role', 'get')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/role/:id', 'get')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/role', 'post')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/role/:id', 'put')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/role/:id', 'delete')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/adminUser', 'get')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/adminUser/:id', 'get')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/adminUser', 'post')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/adminUser/:id', 'put')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/adminUser/:id', 'delete')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/auth/role', 'get')")
	Db.Exec("insert into casbin_rule (p_type, v0, v1, v2) values ('p', 'role:1', '/auth/tree', 'get')")

	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (1, 0, '首页', 1, '', '', 'index')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (2, 0, '后台管理', 1, '', '', 'admin')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (3, 0, '项目管理', 1, '', '', 'project')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (4, 0, '文章管理', 1, '', '', 'article')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (5, 2, '角色', 1, '', '', 'admin-role')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (6, 2, '用户', 1, '', '', 'admin-user')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (7, 5, '角色列表', 0, '/role', 'get', 'admin-role-list')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (8, 5, '角色详情', 0, '/role/:id', 'get', 'admin-role-detail')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (9, 5, '角色添加', 0, '/role', 'post', 'admin-role-add')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (10, 5, '角色修改', 0, '/role/:id', 'put', 'admin-role-edit')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (11, 5, '角色删除', 0, '/role/:id', 'delete', 'admin-role-del')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (12, 6, '用户列表', 0, '/adminUser', 'get', 'admin-user-list')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (13, 6, '用户详情', 0, '/adminUser/:id', 'get', 'admin-user-detail')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (14, 6, '用户添加', 0, '/adminUser', 'post', 'admin-user-add')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (15, 6, '用户修改', 0, '/adminUser/:id', 'put', 'admin-user-edit')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (16, 6, '用户删除', 0, '/adminUser/:id', 'delete', 'admin-user-del')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (17, 6, '获取角色树', 0, '/auth/role', 'get', 'admin-user-role-tree')")
	Db.Exec("insert into auth (id, pid, name, is_menu, api, action, ext) values (18, 5, '获取权限树', 0, '/auth/tree', 'get', 'admin-user-auth-tree')")
}

// 通用分页获取偏移量
func GetOffset(page, pageSize int) int {
	if page <= 1 {
		return 0
	}
	return (page - 1) * pageSize
}
