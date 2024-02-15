package main

import (
	"app/config"
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	controllers.ServeAutor(r)
	controllers.ServeSong(r)

	r.Run(config.Config.Port)
}
