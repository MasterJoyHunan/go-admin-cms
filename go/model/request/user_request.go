package request

type LoginUser struct {
	Username   string `binding:"required,max=255" example:"admin"`                 // 用户名
	Password   string `binding:"required,max=255" example:"admin"`                 // 密码
	VerifyCode string `json:"verify_code" binding:"required,len=4" example:"9527"` // 验证码
}

type UserAdd struct {
	CommonUser
	Password string `binding:"required,max=255" example:"test"` // 密码

}

type UserEdit struct {
	Id int `json:"-"`
	CommonUser
	Password string `example:"test"` // 密码（非必填）

}

type CommonUser struct {
	UserName string `json:"user_name" binding:"required,max=255" example:"test"` // 账号
	RealName string `json:"real_name" binding:"required,max=255" example:"test"` // 真实姓名
	Status   int8   `binding:"oneof=0 1" example:"1"`                            // 状态
	Tel      string `binding:"required,max=12" example:"13054174174"`            // 电话号码
	Roles    [] int `binding:"required,min=1" example:"1,2"`                     // 所属角色
}
