package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ServiceAHandler(w http.ResponseWriter, r *http.Request) {
	// Запрос данных у сервиса B
	response, err := http.Get("http://localhost:8081/service_b")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var dataFromServiceB map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&dataFromServiceB); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Обработка данных и формирование ответа
	data := map[string]interface{}{
		"message":             "This is Service A",
		"data_from_service_b": dataFromServiceB,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/service_a", ServiceAHandler).Methods("GET")

	fmt.Println("Service A is running on :8080")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
