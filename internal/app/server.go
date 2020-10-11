package app

import (
	"github.com/cheremisin-sergey/forum/config"
	"net/http"
)

type App struct {
	Server *http.Server // HTTP middleware
	config *config.Config //
}


// NewServer will create a new instance of the application
func NewApp(config *config.Config) *App {
	app := &App{}
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
		Addr:           config.
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}


	return app
}