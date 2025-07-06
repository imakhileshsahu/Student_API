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

	"github.com/imakhileshsahu/students-api/internal/config"
	"github.com/imakhileshsahu/students-api/internal/http/handlers/student"
	"github.com/imakhileshsahu/students-api/internal/storage/mysql"
)

func main() {

	// 🔷 Read config file path from args
	if len(os.Args) < 3 || os.Args[1] != "--config" {
		log.Fatal("Usage: go run main.go --config <path_to_config>")
	}
	cfg := config.MustLoad(os.Args[2])

	// database setup
	storage, err := mysql.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized",
		slog.String("env", cfg.Env),
		slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address, // 🔷 fixed
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.HTTPServer.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
