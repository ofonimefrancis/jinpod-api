package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/opiumated/jinPod/config"
	"github.com/opiumated/jinPod/handlers"
)

func main() {
	//Initialize Config, which initializes our database
	config := config.InitConfig()

	//Setup handlers {{ USing mux for this project }}
	router := mux.NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3300",
		WriteTimeout: 50 * time.Second,
		ReadTimeout:  50 * time.Second,
	}
	router.Handle("/api/podcasts", handlers.GetAllPodcast(config)).Methods(http.MethodGet)
	router.Handle("/api/podcast/{slug}", handlers.GetPodcast(config)).Methods(http.MethodGet)
	router.Handle("/api/podcast", handlers.AddPodcast(config)).Methods(http.MethodPost)
	log.Fatal(server.ListenAndServe())
}
