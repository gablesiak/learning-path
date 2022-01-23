package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gablesiak/services"
)

type status struct {
	Confirmation string
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		statusCheck := status{Confirmation: "positive"}
		statusCheckJson, err := json.Marshal(statusCheck)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("Write failed: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(statusCheckJson)

	case http.MethodPost:
		newInput, err := services.ValidateRequestBody(r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		services.SaveUser(newInput)
		w.WriteHeader(http.StatusCreated)
		return

	case http.MethodPut:
		newInput, err := services.ValidateRequestBody(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		azureAccess := services.SetStorageAccess()
		services.SaveUser(newInput)
		services.UploadFile(azureAccess)
		return
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":5000", nil)
}
