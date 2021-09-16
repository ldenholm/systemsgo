package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ldenholm/systemsgo/data"
)

type Products struct {
	logger *log.Logger
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

	if req.Method == http.MethodPost {
		p.addProduct(rw, req)
		return
	}

	if req.Method == http.MethodPut {
		p.logger.Println("PUT", req.URL.Path)
		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(req.URL.Path, -1)

		if len(g) != 1 {
			p.logger.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.logger.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.logger.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, req)
		return
	}

	// disallow other request types
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// Convert list to json using json encoding
func (p *Products) getProducts(rw http.ResponseWriter, req *http.Request) {
	p.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to deserialize json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
