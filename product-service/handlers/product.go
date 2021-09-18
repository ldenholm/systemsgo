package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ldenholm/systemsgo/data"
)

type Products struct {
	logger *log.Logger
}

// Contructor
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Convert list to json using json encoding
func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request) {
	p.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)

	// Add to dynamodb

}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	// Pass id from request using vars
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

// Key for context
type KeyProduct struct {
}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

// func SaveToDynamo() {

// 	//-----------------------------------------------
// 	svc := dynamodb.New(session.New())
// 	input := &dynamodb.PutItemInput{
// 		Item: map[string]*dynamodb.AttributeValue{
// 			"AlbumTitle": {
// 				S: aws.String("Somewhat Famous"),
// 			},
// 			"Artist": {
// 				S: aws.String("No One You Know"),
// 			},
// 			"SongTitle": {
// 				S: aws.String("Call Me Today"),
// 			},
// 		},
// 		ReturnConsumedCapacity: aws.String("TOTAL"),
// 		TableName:              aws.String("Music"),
// 	}

// 	result, err := svc.PutItem(input)
// }
