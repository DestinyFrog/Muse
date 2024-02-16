package main

import (
	"app/config"
	"app/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	corsConfig := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://0.0.0.0"}
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))

	controllers.ServeAutor(r)
	controllers.ServeSong(r)

	r.Run(config.Config.Port)
}
