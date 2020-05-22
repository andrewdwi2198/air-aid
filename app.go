package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andrewdwi2198/air-aid/drivers/db"
	"github.com/andrewdwi2198/air-aid/helpers/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := db.InitDB()
	if err != nil {
		log.Panicf("[InitDB]unable to connect to db, err: %s\n", err.Error())
	}

	service := handler.NewHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("/airdata", service.MutateHandler).Methods("POST")

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

	fmt.Println("Server started at port ", port)
	log.Fatal(srv.ListenAndServe())
}
