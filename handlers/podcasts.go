package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	slugify "github.com/metal3d/go-slugify"
	"github.com/opiumated/jinPod/config"
	"github.com/opiumated/jinPod/models"
	"github.com/opiumated/jinPod/utils"
)

//GetAllPodcast Retrieves all the podcasts in the database
func GetAllPodcast(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		podcasts, err := models.Podcasts{}.GetAll(cfg)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Retrieving From Podcast Collection", http.StatusNotFound)
		}

		jsonBytes, err := json.Marshal(podcasts)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Marshaling JSON", http.StatusInternalServerError)
		}
		utils.ResponseWithJSON(w, jsonBytes, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//GetPodcast Gets a single podcast using the passed slug param
func GetPodcast(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		podcast, err := models.Podcasts{}.GetBySlug(cfg, mux.Vars(r)["slug"])
		if err != nil {
			utils.ErrorWithJSON(w, "Error Retrieving Podcast", http.StatusInternalServerError)
		}
		json, err := json.Marshal(podcast)
		utils.ResponseWithJSON(w, json, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//AddPodcast Adds a podcast to the collection
func AddPodcast(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var newPodcast models.Podcasts
		newPodcast.ID = bson.NewObjectId()
		newPodcast.Title = r.FormValue("title")
		newPodcast.Slug = slugify.Marshal(strings.ToLower(r.FormValue("title")))
		newPodcast.Body = r.FormValue("body")
		newPodcast.Description = r.FormValue("description")
		newPodcast.PodcastsURL = r.FormValue("podcast_url")
		newPodcast.DateCreated = time.Now()
		newPodcast.DateUpdated = time.Now()

		err := newPodcast.Add(cfg)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Adding Podcast", http.StatusInternalServerError)
		}
		marshalledJSON, err := json.Marshal(newPodcast)
		if err != nil {
			log.Fatal(err)
			utils.ErrorWithJSON(w, "Error Decoding JSON", http.StatusNoContent)
		}
		utils.ResponseWithJSON(w, marshalledJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//RemovePodcast Removes a podcast from the collection
func RemovePodcast(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := models.Podcasts{}.Remove(cfg, mux.Vars(r)["slug"])
		if err != nil {
			log.Fatal("Error Removing podcast ", err)
			utils.ErrorWithJSON(w, "Error Removing podcast", http.StatusNotImplemented)
		}
		res := &utils.Response{
			Message: "Successfully Removed Podcast",
		}
		responseJSON, err := json.Marshal(res)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Marshalling JSON", http.StatusNotImplemented)
		}
		utils.ResponseWithJSON(w, responseJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
