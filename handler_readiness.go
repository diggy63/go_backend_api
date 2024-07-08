package main

import (
    "net/http"
    "runtime"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
    response := struct {
        Status    string `json:"status"`
        GoVersion string `json:"goVersion"`
    }{
        Status:    "Server is Running",
        GoVersion: runtime.Version(),
    }
    respondWithJSON(w, 200, response)
}