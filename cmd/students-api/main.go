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

	"github.com/igpriyanshunegi/students-api/internal/config"
	"github.com/igpriyanshunegi/students-api/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.MustLoad()
	// database setip
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started %s\n", slog.String("address", cfg.Addr))
	//gracefull shutdown
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//for the production gracefull shutdown is needed
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	//structure login
	slog.Info("Shuting down the server")

	//we give the timerer to finish ongoing requests if there are any before shutting down the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown Successfully")
}
