package services

import (
	"github.com/sampado/bookstore_items-api/src/domain/items"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

var (
	ItemsService ItemsServiceInterface = &itemsService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestError)
	Get(string) (*items.Item, rest_errors.RestError)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestError) {
	item := items.Item{Id: id}
	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}
