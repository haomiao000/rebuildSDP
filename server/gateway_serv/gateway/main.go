package main

import (
	"icode.baidu.com/baidu/personal-code/rebuildSDP/server/gateway_serv/gateway/biz/router"
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}