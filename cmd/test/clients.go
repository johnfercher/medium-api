package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/johnfercher/medium-api/internal/core/models"
)

type Client struct {
	port int
}

// nolint:gomnd // magic number
func NewClient() *Client {
	return &Client{
		port: 8081,
	}
}

func (c *Client) Create(product *models.Product) (*models.Product, error) {
	productBytes, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	requestURL := fmt.Sprintf("http://localhost:%d/products", c.port)

	request, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(productBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Update(product *models.Product) (*models.Product, error) {
	productBytes, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, product.ID)

	request, err := http.NewRequest(http.MethodPut, requestURL, bytes.NewBuffer(productBytes))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) GetByID(id string) (*models.Product, error) {
	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, id)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	response, err := DecodeProduct(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Search(productType string) ([]*models.Product, error) {
	requestURL := fmt.Sprintf("http://localhost:%d/products?type=%s", c.port, productType)

	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Body.Close()

	response, err := DecodeProducts(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return response, nil
}

func (c *Client) Delete(id string) error {
	requestURL := fmt.Sprintf("http://localhost:%d/products/%s", c.port, id)

	request, err := http.NewRequest(http.MethodDelete, requestURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer res.Body.Close()

	return nil
}

func DecodeProduct(r *http.Response) (*models.Product, error) {
	product := &models.Product{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func DecodeProducts(r *http.Response) ([]*models.Product, error) {
	products := []*models.Product{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
