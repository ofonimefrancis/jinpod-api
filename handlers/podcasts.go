package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/opiumated/jinPod/config"
	"github.com/opiumated/jinPod/models"
	"github.com/opiumated/jinPod/utils"
	"fmt"
)

//GetAllPodcast Retrieves all the podcasts in the database
func GetAllPodcast(cfg *config.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		podcasts, err := models.Podcasts{}.GetAll(cfg)
		if err != nil {
			utils.ErrorWithJSON(w, "Error Retrieving From Podcast Collection", http.StatusNotFound)
		}

		jsonBytes, err := json.Marshal(podcasts)
		fmt.Println(string(jsonBytes))
		if err != nil {
			utils.ErrorWithJSON(w, "Error Marshaling JSON", http.StatusInternalServerError)
		}
		utils.ResponseWithJSON(w, jsonBytes, http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
