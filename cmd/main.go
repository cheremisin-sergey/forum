package cmd

import (

	"github.com/cheremisin-sergey/forum/config"
	"github.com/cheremisin-sergey/forum/internal/app"
)

func main() {

	config := config.NewConfig()
	app
	NewApp(config)
	app.N



}
