package main

import (
	"log"
	"net/http"
	"os"
	"main/pkg/handler"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-retryablehttp"
)

func main() {
    r := mux.NewRouter()

	client := retryablehttp.NewClient()
	l := log.New(os.Stdout,"--RETRY--",log.Ldate|log.Ltime|log.Lshortfile)

	h := handler.NewRetry(l,client)

    r.HandleFunc("/vault", h.Get).Methods(http.MethodGet)

	l.Println("[INFO] Starting server at :8080")

    http.ListenAndServe(":8080", r)
}