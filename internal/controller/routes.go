package controller

import "github.com/gin-gonic/gin"

func SetUpRoutes(r *gin.Engine, userCtrl UserCtrl) {
	r.GET("/ping", userCtrl.Ping)
}
