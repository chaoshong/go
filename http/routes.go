package http

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
}

func New() *App {
	return &App{
		engine: gin.Default(),
	}
}

func (a *App) Start() {
	a.Register()
	a.engine.Run("8080")
}

func (a *App) Register() {
	v1 := a.engine.Group("/v1/api")

}
