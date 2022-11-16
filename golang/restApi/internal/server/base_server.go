package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"restApi/internal/config"
	"restApi/pkg/logging"
	"time"
)

type baseServer struct {
}

func New() Server {
	return &baseServer{}
}
func (bs *baseServer) GetServer(router *httprouter.Router, config *config.Config) {
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
