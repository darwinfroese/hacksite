package api

import (
	"github.com/darwinfroese/hacksite/server/pkg/config"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

// Context stores values used by all handlers
type Context struct {
	Logger    *log.Logger
	DB        *database.Database
	Config    *config.EnvironmentConfig
	RequestID string
}
