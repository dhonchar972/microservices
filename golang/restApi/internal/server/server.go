package server

import (
	"github.com/julienschmidt/httprouter"
	"restApi/internal/config"
)

type Server interface {
	GetServer(router *httprouter.Router, config *config.Config)
}
