package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErrror(w, http.StatusNotFound, "not found")
}