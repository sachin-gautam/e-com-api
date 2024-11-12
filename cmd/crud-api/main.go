package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sachin-gautam/go-crud-api/internal/config"
	student "github.com/sachin-gautam/go-crud-api/internal/http/handlers"
	"github.com/sachin-gautam/go-crud-api/internal/middleware"
	"github.com/sachin-gautam/go-crud-api/internal/storage/mysql"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//setup database
	storage, err := mysql.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initialized", slog.String("env", cfg.Env))

	//setup router
	router := http.NewServeMux()
	studentHandler := student.NewStudentHandler(storage)

	router.HandleFunc("POST /api/login", studentHandler.Login)

	router.HandleFunc("POST /api/students", middleware.AuthMiddleware(studentHandler.Create))
	router.HandleFunc("GET /api/students/{id}", studentHandler.Get)
	router.HandleFunc("GET /api/students", studentHandler.GetList)
	router.HandleFunc("PUT /api/students/update/{id}", middleware.AuthMiddleware(studentHandler.Update))
	router.HandleFunc("DELETE /api/students/delete/{id}", middleware.AuthMiddleware(studentHandler.Delete))

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address:", cfg.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown sucessfully")

}
