package handlers

import (
	"log"
	"net/http"

	"github.com/ldenholm/systemsgo/data"
)

type Products struct {
	l *log.Logger
}

// Contructor
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ensure request == GET
	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}

	// disallow other request types
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Convert list to json using json encoding
func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
