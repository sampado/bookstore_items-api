package app

import (
	"net/http"

	"github.com/sampado/bookstore_items-api/src/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.ItemsController.Get).Methods(http.MethodGet)

	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet).Queries("id", "{id}")
}
