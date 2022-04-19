package rest

import (
	"encoding/json"
	"net/http"

	"goapi/graph/model"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{}
	json.NewEncoder(w).Encode(user)
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := model.User{}
	json.NewEncoder(w).Encode(user)
}
