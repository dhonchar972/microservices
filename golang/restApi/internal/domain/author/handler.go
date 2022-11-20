package author

import (
	"fmt"
	"net/http"
	"restApi/internal/handler"
	"restApi/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

/*
Author controller.
*/

// Routing path.
const (
	author     = "/authors"
	authorById = "/authors/:id"
)

type authorHandler struct {
	logger *logging.Logger
}

// Return new author controller(authorHandler).
func GetHandler(logger *logging.Logger) handler.Handler {
	return &authorHandler{
		logger: logger,
	}
}

// Register REST methods in http router.
func (h *authorHandler) Register(router *httprouter.Router) {
	router.GET(author, h.GetAllAuthors)
	router.GET(authorById, h.GetAuthorById)
	router.POST(author, h.CreateAuthor)
	router.PUT(author, h.UpdateAuthor)
	router.DELETE(author, h.DeleteAuthor)
}

// Return list of all users.
func (h *authorHandler) GetAllAuthors(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(200)
}

// Return user with specified ID.
func (h *authorHandler) GetAuthorById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(200)
}

// Create new user from request parameters.
func (h *authorHandler) CreateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

// Updates an existing user with request parameters.
func (h *authorHandler) UpdateAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

// Delete an existing user with request parameters.
func (h *authorHandler) DeleteAuthor(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "This is list of users")
	writer.WriteHeader(204)
}

//func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	name := params.ByName("name")
//	fmt.Fprintf(w, "Hello %s", name) // w.Write([]byte(fmt.Sprintf("Hello %s", name)))
//}
