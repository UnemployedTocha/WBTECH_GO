package main

import (
	"demo_service/internal/repository"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {

	// TODO: logger init
	log := setupLogger()
	log.Info("Starting service")

	fmt.Println("AAaaaaaaaaaaaa")

	// TODO: init storage
	repo, err := repository.NewRepository()

	if err != nil {
		log.Error("Failed to connect to database")
		os.Exit(1)
	}
	_ = repo

	// TODO: init router
	// chi?

	// init start server
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from docker container!!1!1!</h1>")
	})

	http.ListenAndServe(":"+os.Getenv("SERVICE_INTERNAL_PORT"), nil)
}

func setupLogger() (log *slog.Logger) {
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return
}
