package main

import (
	"github.com/cheremisin-sergey/forum/config"
	"github.com/cheremisin-sergey/forum/internal/app"
)

func main() {

	config := config.NewConfig()
	application := app.NewApp(config)
	application.StartServer();

}
