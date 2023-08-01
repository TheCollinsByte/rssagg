package main

import "net/http"

func handerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{}) // This Response will marshal to an empty JSON Object
}
