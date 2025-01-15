package controller

import (
	"refty-backend/internal/controller/render"

	"github.com/gin-gonic/gin"
)

type UserCtrl struct {
}

func NewUserCtrl() UserCtrl {
	return UserCtrl{}
}

func (ct *UserCtrl) Ping(c *gin.Context) {
	render.JSON(c, "pong")
}
