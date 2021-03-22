package main

import (
	"log"
	"main/pkg/handler"
	"main/pkg/send"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-retryablehttp"
)

func main() {
    r := mux.NewRouter()

	client := retryablehttp.NewClient()
	l := log.New(os.Stdout,"--RETRY--",log.Ldate|log.Ltime|log.Lshortfile)

	c := send.NewRetryableRequest(client,l)
	h := handler.NewRetry(l,c)

    r.HandleFunc("/vault", h.Get).Methods(http.MethodGet)
	r.HandleFunc("/vault",h.MakeKV).Methods(http.MethodPost)

	l.Println("[INFO] Starting server at :8080")

    http.ListenAndServe(":8080", r)
}