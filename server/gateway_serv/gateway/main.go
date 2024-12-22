package main

import (
	"github.com/haomiao000/rebuildSDP/server/gateway_serv/gateway/biz/router"

	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}