package main

import (
	"github.com/julienschmidt/httprouter"
	"restApi/internal/config"
	"restApi/internal/domain/user"
	"restApi/internal/server"
	"restApi/pkg/logging"
)

func main() {
	log := logging.GetLogger()
	router := httprouter.New()

	cfg, err := config.GetConfig()
	if err != nil {
		log.Info("file not found!!!")
	}

	handler := user.GetHandler(log)
	handler.Register(router)

	sv := server.New()
	sv.GetServer(router, cfg)
}
