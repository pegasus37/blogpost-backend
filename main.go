package main

import (
	"github.com/pegasus37/blogpost-backend/config"
	"github.com/pegasus37/blogpost-backend/handler"
	"github.com/pegasus37/blogpost-backend/infrastructure"
)

func main() {
	config := config.GetConfig()

	handler := handler.NewHandler(config)
	// handler := handler.NewAnotherHandler()

	infrastructure.StartHttpServer(config, handler)

	// DBHandler := repository.SetUpDBConnection(config)
	// fmt.Println("Connect to database: ", DBHandler)
}
