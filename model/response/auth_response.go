package response

type Auth struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`      // 上级ID
	Name     string `json:"name"`     // 节点名
	IsMenu   int8   `json:"is_menu"`  // 是否是菜单 0:否 1:是
	Api      string `json:"api"`      // 接口
	Action   string `json:"action"`   // 方法
	Ext      string `json:"ext"`      // 前端使用
	Children []Auth `json:"children"` // 下级
}
