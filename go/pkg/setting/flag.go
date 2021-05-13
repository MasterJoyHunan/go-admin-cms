package setting

import (
	"flag"
	"os"
)

var Path string
var Init bool

// 初始化设置参数 - 和 -- 等效
// 有3种方式设置参数
// 1. -init
// 2. -init=true
// 3. -init yes
// 注: 如果是布尔(bool)类型，请不要使用第3种
func init() {
	dir, _ := os.Getwd()
	flag.StringVar(&Path, "config", dir + "/conf/config.yaml", "配置文件路径")
	flag.BoolVar(&Init, "init", false, "初始化 -- 清空数据表，并且插入初始数据")
}
