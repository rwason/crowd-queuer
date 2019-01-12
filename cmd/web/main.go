package main

import (
	"fmt"
	"mux"
	"net/http"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"alive": true, "current_time": %s}`, time.Now())
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Spotify Queuer")
}

func registerRoutes(rt *mux.Router) {
	rt.HandleFunc("/health", health)
	rt.HandleFunc("/", root)
}

func main() {
	fmt.Println("Server starting up...")

	router := mux.NewRouter()
	registerRoutes(router)

	fmt.Println("Server started")
	http.ListenAndServe("localhost:8080", router)
}
