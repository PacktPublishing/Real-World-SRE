package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", "host=localhost dbname=sretest sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO data (metric, value, created) VALUES ($1, $2, $3)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		res, err := stmt.Exec("GET /", 1, time.Now())
		if err != nil || res == nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "{\"hello\": \"world\"}")
	})

	log.Println("Starting server...")
	log.Println("Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
