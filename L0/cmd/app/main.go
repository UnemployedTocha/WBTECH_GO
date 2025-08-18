package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/docker", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from docker container!!1!1!</h1>")
	})

	http.ListenAndServe(":8089", nil)
}
