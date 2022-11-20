package handlers

import "github.com/julienschmidt/httprouter"

// Base inteface for all controllers(handlers).
type Handler interface {
	Register(router *httprouter.Router)
}
