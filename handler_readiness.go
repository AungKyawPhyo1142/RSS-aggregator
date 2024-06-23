package main

import "net/http"

// function has to always like this
// response is just make sure server is running and alive
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}