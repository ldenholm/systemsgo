package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Package level properties
// Errors
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product struct for API resource
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// Serialization consider extracting
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// Deserialization
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

// Getter
func GetProducts() Products {
	return productList
}

// Add Product
func AddProduct(p *Product) {
	// primary key
	p.ID = GetNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := GetById(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

// Get Next ID
func GetNextId() int {
	list := productList[len(productList)-1]
	return list.ID + 1
}

// Get By Id
func GetById(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// Sample data store
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
