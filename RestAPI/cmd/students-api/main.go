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

	"github.com/ashparshp/students-api/internal/config"
	"github.com/ashparshp/students-api/internal/http/handlers/student"
	"github.com/ashparshp/students-api/internal/storage/sqllite"
)

func main() {

	/* load config */
	cfg := config.MustLoad()

	/* database setup */
	storage, err := sqllite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized...", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	/* setup router */
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))

	/* setup server */
	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started... %s", slog.String("Address: ", cfg.Addr))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start...")
		}
	}()

	<- done

	slog.Info("Shtting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shut down the server...", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully...")
}
