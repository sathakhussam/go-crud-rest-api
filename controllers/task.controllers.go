package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/sathakhussam/go-mod-crud/models"
)

var Tasks []models.Task

func TasksMethodHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		GetTasks(w, r)
	} else {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "error",
			"data":   "No routes available",
		})
	}
}

type getTaskResponseSchema struct {
	Status string        `json:"status"`
	Data   []models.Task `json:"data"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getTaskResponseSchema{
		Status: "success",
		Data:   Tasks,
	})
}
