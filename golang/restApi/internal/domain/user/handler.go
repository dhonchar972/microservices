package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	ihandler "restApi/internal/handler"
	"restApi/pkg/logging"
)

const (
	user     = "/users"
	userById = "/users/:id"
)

type handler struct {
	logger *logging.Logger
}

func GetHandler(logger *logging.Logger) ihandler.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(user, h.GetAllUsers)
	router.GET(userById, h.GetUserById)
	router.POST(user, h.CreateUser)
	router.PUT(user, h.UpdateUser)
	router.DELETE(user, h.DeleteUser)
}

func (h *handler) GetAllUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(200)
}

func (h *handler) GetUserById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(200)
}

func (h *handler) CreateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

func (h *handler) UpdateUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

func (h *handler) DeleteUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

//func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	name := params.ByName("name")
//	fmt.Fprintf(w, "Hello %s", name) // w.Write([]byte(fmt.Sprintf("Hello %s", name)))
//}
