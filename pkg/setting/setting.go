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
	ImageUploadConf   = ImageUpload{}
	RedisConf         = Redis{}
	QrcodeConf        = Qrcode{}
	ElasticSearchConf = ElasticSearch{}
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

	if err := viper.UnmarshalKey("image_upload", &ImageUploadConf); err != nil {
		log.Panic("image_upload配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("redis", &RedisConf); err != nil {
		log.Panic("redis配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("qrcode", &QrcodeConf); err != nil {
		log.Panic("qrcode配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("elastic_search", &ElasticSearchConf); err != nil {
		log.Panic("elastic_search配置文件格式错误", err)
	}

	if err := viper.UnmarshalKey("captcha", &CaptchaConf); err != nil {
		log.Panic("captcha配置文件格式错误", err)
	}
}
