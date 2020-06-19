package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/sampado/bookstore_items-api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	mapUrls()

	address := "127.0.0.1:8080"
	srv := &http.Server{
		Handler: router,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("server starting at:", address)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
