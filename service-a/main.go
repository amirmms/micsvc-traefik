package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Service string `json:"service"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)

		response := Response{
			Message: "hi from svc A",
			Service: "svc A",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hi form svc a :)")
	})

	http.HandleFunc("/call-b", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("let wake up & call svc B ...")

		serviceBURL := os.Getenv("SERVICE_B_URL")
		if serviceBURL == "" {
			serviceBURL = "http://service-b:8080"
		}

		resp, err := http.Get(serviceBURL)
		if err != nil {
			log.Printf("err calling svc B: %v", err)
			http.Error(w, "err calling svc B", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var serviceBResp Response
		if err := json.NewDecoder(resp.Body).Decode(&serviceBResp); err != nil {
			log.Printf("Error decoding response from svc B: %v", err)
			http.Error(w, "Failed to decode response from svc B", http.StatusInternalServerError)
			return
		}

		combinedResp := Response{
			Message: "svc A received from svc B: " + serviceBResp.Message,
			Service: "svc A -> svc B",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(combinedResp)
	})

	log.Printf("svc A starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
