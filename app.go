package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {	
	decoder := json.NewDecoder(r.Body)
	var data map[string]interface{}
	decoder.Decode(&data)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message Received")
	fmt.Println(data)
}

func main() {
	godotenv.Load()
	r := mux.NewRouter()

	r.HandleFunc("/airdata", Handler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started at port 8000")
	log.Fatal(srv.ListenAndServe())
}
