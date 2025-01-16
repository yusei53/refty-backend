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

// Ping godoc
// @Summary Ping the server
// @Description Returns "pong" if server is working
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ping [get]
func (ct *UserCtrl) Ping(c *gin.Context) {
	render.JSON(c, "pong")
}
