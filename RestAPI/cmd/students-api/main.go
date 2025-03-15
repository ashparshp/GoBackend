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
)

func main() {

	/* load config */
	cfg := config.MustLoad()

	/* database setup */

	/* setup router */
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wecome..."))
	})

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

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shut down the server...", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully...")
}
