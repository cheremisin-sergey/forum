package context

import (
	"database/sql"
	"github.com/cheremisin-sergey/forum/config"
)

type AppContext struct {
	DB     *sql.DB
	Config *config.Config
}

