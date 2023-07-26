package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something Went Wrong")
}
