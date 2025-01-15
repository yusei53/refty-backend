package container

import (
	"fmt"
	"refty-backend/internal/app/config"
	"refty-backend/internal/controller"
	"refty-backend/internal/domain/infra/logging"

	"github.com/gin-gonic/gin"
)

type container struct {
	userCtrl controller.UserCtrl
}

func NewCtrl(userCtrl controller.UserCtrl) container {
	return container{
		userCtrl: userCtrl,
	}
}

type App struct {
	r   *gin.Engine
	cfg *config.Config
}

func NewApp(r *gin.Engine, container container, cfg *config.Config) *App {
	logging.Init()

	controller.SetUpRoutes(r, container.userCtrl)

	return &App{
		r:   r,
		cfg: cfg,
	}
}

func (a *App) Run() error {
	return a.r.Run(fmt.Sprintf(":%d", a.cfg.Port))
}
