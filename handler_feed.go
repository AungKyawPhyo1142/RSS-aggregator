package main

import (
	"encoding/json"

	"net/http"
	"time"

	"github.com/AungKyawPhyo1142/RSS-aggregator/internal/database"

	"github.com/google/uuid"
)

// function has to always like this
// response is just make sure server is running and alive
func (apiCfg *API_CONFIG) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErrror(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithErrror(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}

	respondWithJSON(w, http.StatusCreated, DatabaseFeedToFeed(feed))
}

func (apiCfg *API_CONFIG) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithErrror(w, http.StatusInternalServerError, "Failed to get feeds")
		return
	}
	respondWithJSON(w, http.StatusOK, DatabaseFeedsToFeeds(feeds))
}