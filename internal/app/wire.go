//go:build wireinject

package app

import (
	"refty-backend/internal/app/config"
	"refty-backend/internal/app/container"
	"refty-backend/internal/controller"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func New() (*container.App, error) {
	wire.Build(
		provideGinEngine,
		config.New,

		controller.NewUserCtrl,

		container.NewCtrl,
		container.NewApp,
	)
	return nil, nil
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
