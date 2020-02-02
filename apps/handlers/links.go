package handlers

import "net/http"

func GetLinkListHandler(w http.ResponseWriter, r *http.Request) {
	result := "result"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
