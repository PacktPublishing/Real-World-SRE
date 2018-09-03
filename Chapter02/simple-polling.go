package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]float64{
			"hello": 1.0,
			"now":   float64(time.Now().Unix()),
			"rand":  rand.Float64(),
		}

		json, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "%s", json)
	})

	log.Println("Starting server...")
	log.Println("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
