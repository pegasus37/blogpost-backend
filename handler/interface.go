package handler

import "net/http"

type HandlerInterface interface {
	PingHandler(writer http.ResponseWriter, request *http.Request)
	GetPegasusHandler(writer http.ResponseWriter, request *http.Request)
	PostPegasusHandler(writer http.ResponseWriter, request *http.Request)
}
