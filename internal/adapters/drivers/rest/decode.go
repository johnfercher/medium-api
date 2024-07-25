package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func DecodeStringIDFromURI(r *http.Request) (string, error) {
	id := chi.URLParam(r, "id")
	if id == "" {
		return "", errors.New("empty_id_error")
	}

	return id, nil
}

func DecodeTypeQueryString(r *http.Request) string {
	return r.URL.Query().Get("type")
}

func DecodeCreateProductFromBody(r *http.Request) (*CreateProduct, error) {
	createProduct := &CreateProduct{}
	err := json.NewDecoder(r.Body).Decode(&createProduct)
	if err != nil {
		return nil, err
	}

	return createProduct, nil
}

func DecodeProductFromBodyAndURI(r *http.Request) (*Product, error) {
	id, err := DecodeStringIDFromURI(r)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return nil, err
	}

	product.ID = id

	return product, nil
}
