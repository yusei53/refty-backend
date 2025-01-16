package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRoutes(r *gin.Engine, userCtrl UserCtrl) {
	var name, password string
	if name = os.Getenv("BASIC_AUTH_USERNAME"); name == "" {
		panic("BASIC_AUTH_USERNAME is not set")
	}
	if password = os.Getenv("BASIC_AUTH_PASSWORD"); password == "" {
		panic("BASIC_AUTH_PASSWORD is not set")
	}
	r.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{name: password}), ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", userCtrl.Ping)
}
