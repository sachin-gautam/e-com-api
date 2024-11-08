package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sachin-gautam/go-crud-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//setup database
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Crud API"))
	})
	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("Sever Started %s", cfg.HttpServer.Address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server")
	}

}
