//Package clissification of product API
//Documentation for Product API
package handlers

import (
	"log"
	"microservice/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Alist of products
type productsResponse struct {
	//All products in the system
	// in:
	Body []data.Product
}

//swagger:response noContent
type productIDParameterWrapper struct {
	//The id of the product to delete from the database
	//in:path
	// required: true
	ID int `json:"id"`
}
type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
// 	p.l.Println("Handle GET Product")
// 	lp := data.GetProducts()
// 	err := lp.ToJSON(rw)
// 	if err != nil {
// 		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
// 	}
// }
// func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
// 	p.l.Println("Handle POST Product")

// 	prod := r.Context().Value(KeyProduct{}).(data.Product)
// 	data.AddProduct(prod)
// }

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
