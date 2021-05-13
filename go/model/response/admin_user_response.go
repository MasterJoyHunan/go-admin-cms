package response

type AdminUserPage struct {
	Total       int             `json:"total"`        // 总共多少页
	PerPage     int             `json:"per_page"`     // 当前页码
	CurrentPage int             `json:"current_page"` // 每页显示多少条
	Data        []AdminUserList `json:"data"`
}

type CasRole struct {
	Id   int    `json:"id"`   // 角色ID
	Name string `json:"name"` // 角色名
}

type AdminUserList struct {
	Id       int       `json:"id"`
	UserName string    `json:"user_name"` // 登录名
	Tel      string    `json:"tel"`       // 手机号码
	RealName string    `json:"real_name"` // 真实姓名
	Status   int8      `json:"status"`    // 用户状态
	Roles    []CasRole `json:"roles"`     // 角色信息
}
