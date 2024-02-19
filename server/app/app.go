package app

import (
	"app/config"
	"app/controllers"
	"app/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Server *gin.Engine
}

func (o *Application) Initialize() {
	db.Setup()

	o.Server = gin.Default()
	o.Server.MaxMultipartMemory = 8 << 20

	corsConfig := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://0.0.0.0"}
	corsConfig.AllowAllOrigins = true

	o.Server.Use(cors.New(corsConfig))

	controllers.ServeAutor(o.Server)
	controllers.ServeSong(o.Server)

	o.Server.Run(config.Config.Port)
}