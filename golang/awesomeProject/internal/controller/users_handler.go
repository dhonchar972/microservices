package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUsersHandler(router *mux.Router) {
	router.HandleFunc("/users", getAllUsers)
	router.HandleFunc("/users/{id}", getUserById)
}

func getUserById(writer http.ResponseWriter, reader *http.Request) {
	args := mux.Vars(reader)
	id := args["id"] // the book title slug
	writer.WriteHeader(http.StatusOK)
	//json.NewEncoder(writer).Encode(EntityThatMustBeSend)
	fmt.Fprintf(writer, "You're id is: %s", id)
}

func getAllUsers(writer http.ResponseWriter, reader *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "All users")
}
