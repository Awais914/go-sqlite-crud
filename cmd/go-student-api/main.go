package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Awais914/go-students-api/internal/config"
	"github.com/Awais914/go-students-api/internal/http/handlers/student"
	"github.com/Awais914/go-students-api/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /", student.Create(storage))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server started at", slog.String("address", cfg.Address))

	interruptChan := make(chan os.Signal, 1)

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func()  {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Failed to start server :(")
		}
	}()

	<-interruptChan

	slog.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfull")
}
