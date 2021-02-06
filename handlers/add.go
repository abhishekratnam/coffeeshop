package handlers

import (
	"microservice/data"
	"net/http"
)

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handler POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
