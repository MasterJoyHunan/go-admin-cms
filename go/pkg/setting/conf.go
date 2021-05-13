package setting

// app 配置
type Application struct {
	Env        string
	Host       string
	Port       int
	Doamin     string
	PublicPath string `mapstructure:"public_path"`
}

// mysql 数据库配置
type Mysql struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Db        string
	Prefix    string
	MaxIdle   int `mapstructure:"max_idle"`
	MaxActive int `mapstructure:"max_active"`
}

// 日志配置
type Log struct {
	Path  string
	Level string
}

// casbin配置
type Casbin struct {
	Path string
}

// 验证码配置
type Captcha struct {
	Length int
	Width  int
	Height int
	Str    string
	Font   []string
	Noise  int
	Line   int
}

// json web token 配置
type Jwt struct {
	Key     string
	ExpTime int `mapstructure:"exp_time"`
}