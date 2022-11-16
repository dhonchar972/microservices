package main

import (
	"awesomeProject/internal/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	controller.RegisterUsersHandler(router)

	log.Println("Service started at port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
