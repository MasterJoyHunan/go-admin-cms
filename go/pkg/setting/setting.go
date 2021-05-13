package setting

import (
	"github.com/spf13/viper"
	"log"
)

var (
	ApplicationConf   = Application{}
	MysqlConf         = Mysql{}
	LogConf           = Log{}
	CasbinConf        = Casbin{}
	JwtConf           = Jwt{}
	CaptchaConf       = Captcha{}
)

func Setup() {
	viper.SetConfigFile(Path)
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("读取配置文件错误", err)
	}

	if err := viper.UnmarshalKey("application", &ApplicationConf); err != nil {
		log.Panic("APP配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("mysql", &MysqlConf); err != nil {
		log.Panic("mysql配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("log", &LogConf); err != nil {
		log.Panic("log配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("casbin", &CasbinConf); err != nil {
		log.Panic("casbin配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("jwt", &JwtConf); err != nil {
		log.Panic("jwt配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("captcha", &CaptchaConf); err != nil {
		log.Panic("captcha配置文件格式错误", err)
	}
}
