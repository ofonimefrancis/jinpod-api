package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"github.com/opiumated/jinPod/config"
	"github.com/opiumated/jinPod/models"
	"github.com/opiumated/jinPod/utils"
)

//AddAuthor adds an author to the collection
func AddAuthor(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var author models.Author
		author.ID = bson.NewObjectId()
		author.Name = r.FormValue("name")
		author.AvatarURL = r.FormValue("avatar_url")
		author.Country = r.FormValue("country")
		author.DateCreated = time.Now()
		author.DateUpdated = time.Now()

		err := author.Add(cfg)
		if err != nil {
			res := &utils.Response{
				Message: "Failed to add a new Author",
			}

			responseJSON, err := json.Marshal(res)
			if err != nil {
				utils.ErrorWithJSON(w, "Cannot Marshal JSON", http.StatusInternalServerError)
			}
			utils.ResponseWithJSON(w, responseJSON, http.StatusOK)
		}

		successResponse := &utils.Response{
			Message: "Successfully added a new author",
		}
		successResponseJSON, err := json.Marshal(successResponse)
		if err != nil {
			utils.ErrorWithJSON(w, "Cannot Marshal JSON", http.StatusInternalServerError)
		}
		utils.ResponseWithJSON(w, successResponseJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//GetAllAuthors Get all the authors in our collection
func GetAllAuthors(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authors, err := models.Author{}.GetAll(cfg)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Retrieving Authors", http.StatusInternalServerError)
		}
		authorJSON, err := json.Marshal(authors)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Marshalling JSON", http.StatusNotImplemented)
		}
		utils.ResponseWithJSON(w, authorJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//GetAuthor Gets an author by their id in our collection
func GetAuthor(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		author, err := models.Author{}.Get(cfg, mux.Vars(r)["id"])
		if err != nil {
			utils.ErrorWithJSON(w, "Error Retrieving Author", http.StatusInternalServerError)
		}
		authorJSON, _ := json.Marshal(author)
		utils.ResponseWithJSON(w, authorJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}

//RemoveAuthor Remove an Author by ID
func RemoveAuthor(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		err := models.Author{}.Remove(cfg, mux.Vars(r)["id"])
		if err != nil {
			utils.ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
		}
		res := &utils.Response{
			Message: "Successfully Removed Author",
		}
		responseJSON, err := json.Marshal(res)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Marshalling JSON", http.StatusNotImplemented)
		}
		utils.ResponseWithJSON(w, responseJSON, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
