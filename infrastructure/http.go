package infrastructure

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/handler"
)

const (
	ping = "/ping"

	get_pegasus  = "/pegasus"
	post_pegasus = "/pegasus"
)

func StartHttpServer(config *config.Configuration, handler handler.HandlerInterface) {
	router := mux.NewRouter()

	router.HandleFunc(ping, handler.PingHandler).Methods("GET")
	router.HandleFunc(get_pegasus, handler.GetPegasusHandler).Methods("GET")
	router.HandleFunc(post_pegasus, handler.PostPegasusHandler).Methods("POST")

	startServer(config, router)
}

func startServer(config *config.Configuration, router *mux.Router) {
	fmt.Println("💚💚💚 Pegasus37 is now online")

	err := http.ListenAndServe(config.HTTP_PORT, router)
	if err != nil {
		log.Fatal("ListenAndServe Error:", err)
	}
}
