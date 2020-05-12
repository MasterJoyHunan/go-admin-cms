package response

type RolePage struct {
	Total       int        `json:"total"`        // 总共多少页
	PerPage     int        `json:"per_page"`     // 当前页码
	CurrentPage int        `json:"current_page"` // 每页显示多少条
	Data        []RoleList `json:"data"`
}

type RoleList struct {
	Id         int    `json:"id"`
	Pid        int    `json:"pid"`         // 上级ID
	ParentName string `json:"parent_name"` // 上级角色名
	Name       string `json:"name"`        // 角色名
	Auths      string `json:"-"`           // 所有权限str
	Auth       []int  `json:"auth"`        // 所有权限arr
	BaseAuth   []int  `json:"base_auth"`   // api权限
}

type Roles struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`      // 角色名
	Children []Roles `json:"children"` // 下级角色
}
