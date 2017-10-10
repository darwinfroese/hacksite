package api

import (
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// Context stores values used by all handlers
type Context struct {
	DB        *database.Database
	Config    *models.ServerConfig
	RequestID string
}
