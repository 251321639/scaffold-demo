package main

import (
	"fmt"
	"scaffold-demo/config"

	// _ "scaffold-demo/config"
	"scaffold-demo/utils/jwtutil"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	logs.Info(nil, "服务启动成功")
	ss, _ := jwtutil.GenToken("ddd")
	_, err := jwtutil.ParseToken(ss)
	if err != nil {
		logs.Error(nil, "解析token失败")
	} else {
		logs.Info(nil, "解析token成功")
	}
	fmt.Println(ss)
	r.Run(config.Port)
}
