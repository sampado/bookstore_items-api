package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sampado/bookstore_items-api/domain/items"
	"github.com/sampado/bookstore_items-api/services"
	"github.com/sampado/bookstore_items-api/utils/http_utils"
	"github.com/sampado/bookstore_oauth-go/oauth"
	"github.com/sampado/bookstore_utils-go/logger"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c itemsController) Create(w http.ResponseWriter, r *http.Request) {

	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondJsonError(w, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondJsonError(w, resErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		resErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondJsonError(w, resErr)
		return
	}

	itemRequest.Seller = oauth.GetClientId(r)

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondJsonError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c itemsController) Get(w http.ResponseWriter, r *http.Request) {
	itemId := strings.TrimSpace(r.FormValue("id"))
	if itemId == "" {
		resErr := rest_errors.NewBadRequestError("invalid request ID")
		logger.Error(fmt.Sprintf("fetching item id [%s]", itemId), resErr)
		http_utils.RespondJsonError(w, resErr)
		return
	}

	result, err := services.ItemsService.Get(itemId)
	if err != nil {
		http_utils.RespondJsonError(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, result)
}
