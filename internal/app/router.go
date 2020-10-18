package app

import (
	api "github.com/cheremisin-sergey/forum/internal/api"
	"github.com/cheremisin-sergey/forum/internal/app/middleware"
	"github.com/cheremisin-sergey/forum/internal/context"
	"github.com/gin-gonic/gin"
)

func InitRouter(app App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	appContext := context.AppContext{
		DB:     app.db,
		Config: app.config,
	}

	r.Use(middleware.AppContext(&appContext))

	routerGroup := r.Group("db/api")
	routerGroup.POST("/user/create", api.CreateUser)
	routerGroup.GET("/user/get", api.GetUser)

	return r
}
