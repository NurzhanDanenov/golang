package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ServiceBHandler(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8080/service_a")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var dataFromServiceA map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&dataFromServiceA); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"message":             "This is Service B",
		"data_from_service_a": dataFromServiceA,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/service_b", ServiceBHandler).Methods("GET")

	fmt.Println("Service B is running on :8081")
	http.Handle("/", r)
	http.ListenAndServe(":8081", nil)
}
