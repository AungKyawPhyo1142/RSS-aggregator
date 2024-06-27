package main

import (
	"fmt"
	"net/http"

	"github.com/AungKyawPhyo1142/RSS-aggregator/internal/auth"
	"github.com/AungKyawPhyo1142/RSS-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *API_CONFIG) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
	
		fmt.Printf("API Key: %v\n", apiKey)
		if err != nil {
			respondWithErrror(w, http.StatusUnauthorized, fmt.Sprintf("no auth header: %v", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)

		if err != nil {
			respondWithErrror(w, http.StatusBadRequest, fmt.Sprintf("can't find user: %v", err))
		}

		handler(w, r, user)

	}
}
