package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ldenholm/systemsgo/data"
)

type Product struct {
	l *log.Logger
}

// Contructor
func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()

	// Convert list to json using json encoding
	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(data)
}
