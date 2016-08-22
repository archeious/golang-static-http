package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./content/"))))
	logged := handlers.LoggingHandler(os.Stdout, r)
	srv := &http.Server{
		Handler:	logged,
		Addr:		"198.154.238.237:80",
		WriteTimeout:	15 * time.Second,
		ReadTimeout:	15 * time.Second,
	}
	
	log.Fatal(srv.ListenAndServe())
}
