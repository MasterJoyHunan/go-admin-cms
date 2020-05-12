package main

import (
	"github.com/robfig/cron"
	"log"
)

//Cron 表达式格式
//字段名	        是否必填	允许的值	允许的特殊字符
//秒（Seconds）	Yes	    0-59	* / , -
//分（Minutes）	Yes	    0-59	* / , -
//时（Hours）	Yes	    0-23	* / , -
//一个月中的某天（Day of month）	Yes	1-31	* / , - ?
//月（Month）	Yes	1-12 or JAN-DEC	* / , -
//星期几（Day of week）	Yes	0-6 or SUN-SAT	* / , - ?
func main() {
	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		log.Print("每5秒执行一次")
	})
	c.Start()
	select {}
}
