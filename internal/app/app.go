package app

import (
	"github.com/cheremisin-sergey/forum/config"
	"net/http"
)

type App struct {
	Server *http.Server // HTTP middleware
	config *config.Config
}
