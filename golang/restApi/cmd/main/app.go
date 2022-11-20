/*
Application entry point.
*/
package main

import (
	"restApi/internal/config"
	"restApi/internal/domain/user"
	"restApi/internal/server"
	"restApi/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Creating a new logger
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
