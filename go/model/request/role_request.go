package request

type RoleAdd struct {
	Role
}

type RoleEdit struct {
	Id int `json:"-"`
	Role
}
type Role struct {
	Pid  int    `example:"1"`                              // 上级ID
	Name string `binding:"required,max=255" example:"程序员"` // 角色名
	Auth []int  `binding:"required,unique,min=1" example:"1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19"`
}
