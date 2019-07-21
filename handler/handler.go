package handler

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/go-chi/render"
	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/entity"
	"github.com/pegasus37/blogpost-backend/service"
)

type Handler struct {
	Config *config.Configuration
}

func NewHandler(config *config.Configuration) *Handler {
	h := new(Handler)
	h.Config = config
	return h
}

func (h Handler) PingHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[LOG: Handler: PingHandler, Message: Call Ping]")

	writer.WriteHeader(200)
	render.JSON(writer, request, "pong from handler")
}

func (h Handler) GetPegasusHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[LOG: Handler: GetPegasusHandler, Message: Get Pegasus]")
	pegasus := service.GetPegasus(h.Config)

	writer.WriteHeader(200)
	render.JSON(writer, request, pegasus)
}

func (h Handler) PostPegasusHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("[LOG: Handler: PostPegasusHandler, Message: Post Pegasus]")

	var data []entity.Pegasus
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println("[LOG - Error: Handler: PostPegasusHandler, Error: ", err)
	}

	service.InsertPegasus(h.Config, data)

	writer.WriteHeader(200)
	render.JSON(writer, request, "Success insert data")
}
