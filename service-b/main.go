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
		log.Printf("get a request: %s %s", r.Method, r.URL.Path)

		response := Response{
			Message: "hi from svc B",
			Service: "svc B",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Handler for health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "hi from svc b :)")
	})

	log.Printf("svc B on %s port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("err to start : %v", err)
	}
}
