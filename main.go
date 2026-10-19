package main

import (
	// 只执行init函数
	"scaffold-demo/config"
	// _ "scaffold-demo/config"
	"scaffold-demo/utils/logs"

	"github.com/gin-gonic/gin"
)

func main() {
	// Your code here
	r := gin.Default()
	logs.Info(nil, "服务启动成功")
	r.Run(config.Port)

}
