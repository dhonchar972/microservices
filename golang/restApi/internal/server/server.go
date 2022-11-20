/*
Package contains methods for creating new web server.
*/
package server

import (
	"fmt"
	"net"
	"net/http"
	"restApi/internal/config"
	"restApi/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Server.
type server struct {
}

// Return new server instance.
func New() *server {
	return &server{}
}

// Register router and configurations and start server.
func (bs *server) GetServer(router *httprouter.Router, config *config.Config) {
	log := logging.GetLogger()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Listen.BindIp, config.Listen.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("App run on %s:%s", config.Listen.BindIp, config.Listen.Port)
	log.Fatalln(server.Serve(listener))
}
