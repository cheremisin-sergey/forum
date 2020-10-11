package app

import (
	"fmt"
	"github.com/cheremisin-sergey/forum/config"
	"net/http"
)

type App struct {
	Server *http.Server   // HTTP middleware
	config *config.Config //
}

// NewServer will create a new instance of the application
func (app *App) NewApp(config *config.Config) *App {
	app.config = config

	//err := app.modelRegistry.OpenWithConfig(config)
	//
	//if err != nil {
	//	log.Fatalf("gorm: could not connect to db %q", err)
	//}
	//
	//app.db = app.modelRegistry.DB
	//app.Server = NewRouter(app)

	routersInit := InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerPort),
		Handler:        routersInit,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	app.Server = server
	return app
}

func (app *App) startServer() {
	app.Server.ListenAndServe()
}
