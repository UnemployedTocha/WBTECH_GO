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
	_, err := repository.NewDb(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("SSL_MODE"),
	})

	if err != nil {
		log.Error("Failed to init db: ", err.Error())
	}

	//rep := repository.NewRepository(db)

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
