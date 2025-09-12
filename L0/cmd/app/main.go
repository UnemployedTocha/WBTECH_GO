package main

import (
	"demo_service/internal/http-server/handlers"
	k "demo_service/internal/kafka"
	"demo_service/internal/repository"
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

//go:embed static/*
var page embed.FS

func main() {

	// TODO: logger init
	log := setupLogger()
	log.Info("Starting service")

	// TODO: init storage
	repo, err := repository.NewRepository(log)

	if err != nil {
		log.Error("Failed to connect to database")
		os.Exit(1)
	}
	_ = repo

	h := k.NewHandler(repo, log)
	consumer, err := k.NewConsumer("kafka:29091", h, "some_topic", "123", log)

	if err != nil {
		log.Error("AAAAAAA", err)
	}

	go consumer.Start()

	// TODO: init router
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	content, _ := fs.Sub(page, "static")
	router.Handle("/*", http.FileServer(http.FS(content)))

	router.Route("/orders", func(r chi.Router) {
		r.Get("/{order_uid}", handlers.New(repo, log)) // GET /orders/{order_uid} - конкретный заказ
	})

	server := &http.Server{
		Addr:         ":" + os.Getenv("SERVICE_INTERNAL_PORT"),
		Handler:      router,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err = server.ListenAndServe(); err != nil {
		log.Error("server start failed :(")
	}
}

func setupLogger() (log *slog.Logger) {
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return
}
