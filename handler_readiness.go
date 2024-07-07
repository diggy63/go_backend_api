package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{"Server is Running"}
	respondWithJSON(w, 200, response)
}
