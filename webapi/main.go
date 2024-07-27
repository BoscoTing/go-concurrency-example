package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", newGuidHandler)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func newGuidHandler(w http.ResponseWriter, r *http.Request) {
	count := r.URL.Query().Get("count") // get 'count' param in URL
	if count == "" {
		count = "1"
	}
	times, err := strconv.Atoi(count) // str to int
	if err != nil {
		http.Error(w, "Invalid count value", http.StatusBadRequest) 
		return
	}
	uuids := make([]string, times) // make a []string with 'times' length
	for i := 0; i < times; i++ {
		uuids[i] = generateGUID()
	}

	jsonResponse, err := json.Marshal(uuids)
	if err != nil {
		http.Error(w, "Failed to generate JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	time.Sleep(time.Second) // delay 1 second and return response
	w.Write(jsonResponse)
}

func generateGUID() string {
	return uuid.New().String()
}
