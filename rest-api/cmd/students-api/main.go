package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mapur2/lets_go/rest-apis/internal/config"
	"github.com/mapur2/lets_go/rest-apis/internal/http/handlers/student"
	"github.com/mapur2/lets_go/rest-apis/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	//database
	db, dbErr := sqlite.New(cfg)
	if dbErr != nil {
		log.Fatal("Db error", dbErr)
	}

	slog.Info("db connected", db)
	router := http.NewServeMux()

	router.HandleFunc("GET /", student.New())
	router.HandleFunc("POST /", student.Create(db))
	router.HandleFunc("GET /{email}", student.GetStudentByEmail(db))

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server started on ", cfg.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done //this blocks the main thread and only ends if any signal it recieves

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		log.Fatal("failed to shutdown error", err.Error())
	}
	fmt.Println("Server shutdown successfully")
}
