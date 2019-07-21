package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

type AnotherHandler struct {
}

func NewAnotherHandler() *AnotherHandler {
	ah := new(AnotherHandler)
	return ah
}

func (ah AnotherHandler) PingHandler(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(200)
	render.JSON(writer, request, "pong from antoher handler")
}
