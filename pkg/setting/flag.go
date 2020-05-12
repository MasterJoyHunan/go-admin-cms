package setting

import (
	"flag"
	"os"
)

var Path string

func init() {
	dir, _ := os.Getwd()
	flag.StringVar(&Path, "config", dir + "/conf/config.yaml", "配置文件路径")
}
