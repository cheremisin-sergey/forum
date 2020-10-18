package app

import (
	"database/sql"
	"fmt"
	"github.com/cheremisin-sergey/forum/config"
	pg_db "github.com/cheremisin-sergey/forum/service/pg-db"
	_ "github.com/lib/pq"
	"net/http"
)

type App struct {
	Server   *http.Server   // HTTP middleware
	config   *config.Config //
	db *sql.DB
}

// NewServer will create a new instance of the application
func NewApp(config *config.Config) *App {
	app := App{}
	app.config = config

	pgClient, err := pg_db.New(config)
	if err != nil {
		panic(err)
	}

	app.db = pgClient

	router := InitRouter(app)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.ServerPort),
		Handler:        router,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	app.Server = server
	return &app
}

func (app *App) StartServer() {
	err := app.Server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
