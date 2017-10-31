package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/opiumated/jinPod/config"
	"github.com/opiumated/jinPod/handlers"
	"github.com/rs/cors"
)

func main() {
	//Initialize Config, which initializes our database
	config := config.InitConfig()

	//Setup handlers {{ USing mux for this project }}
	router := mux.NewRouter()
	//Add default CORS
	handler := cors.Default().Handler(router)

	server := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:3400",
		WriteTimeout: 50 * time.Second,
		ReadTimeout:  50 * time.Second,
	}

	//Routes
	router.Handle("/api/podcasts", handlers.GetAllPodcast(config)).Methods(http.MethodGet)
	router.Handle("/api/podcast/{slug}", handlers.GetPodcast(config)).Methods(http.MethodGet)
	router.Handle("/api/podcast", handlers.AddPodcast(config)).Methods(http.MethodPost)
	router.Handle("/api/podcast/{slug}", handlers.RemovePodcast(config)).Methods(http.MethodDelete)

	router.Handle("/api/author/{name}", handlers.GetAuthor(config)).Methods(http.MethodGet)
	router.Handle("/api/authors", handlers.GetAllAuthors(config)).Methods(http.MethodGet)
	//TODO: AddAuthor works but does not return a restful JSON to the user
	router.Handle("/api/author", handlers.AddAuthor(config)).Methods(http.MethodPost)
	router.Handle("/api/author/{id}", handlers.RemoveAuthor(config)).Methods(http.MethodDelete)
	log.Fatal(server.ListenAndServe())
}
