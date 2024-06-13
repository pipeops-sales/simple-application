package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `json:"message"`
		}{Message: "Hello, World!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	log.Println("Starting server on :3000...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
