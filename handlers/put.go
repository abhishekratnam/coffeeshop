package handlers

import (
	"microservice/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	//fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		var b = []byte(`"Product not found in the database"`)

		// str := {Message: }
		data.ToJSON(b, rw) //error
		return
	}
	rw.WriteHeader(http.StatusNoContent)

}
