// @title MasterJoy博客
package main

import (
	"blog/model"
	"blog/pkg/logger"
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
)

func main() {
	// 初始化操作 (因为 init 方法无法保证我们想要的顺序)
	setting.Setup()
	logger.Setup()
	model.Setup()

	router := routers.InitRouter()
	panic(router.Run(fmt.Sprintf("%s:%d", setting.ApplicationConf.Host, setting.ApplicationConf.Port)))
}
