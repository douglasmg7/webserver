package main

import (
	"log"
	"webserver/models"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Product list
func AdminProductListHandlerGet(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get all products
	products, err := models.GetAllProduct()
	if err != nil {
		log.Printf("err: %v", err)
	}
	// log.Printf("Products: %v", products)

	data := struct {
		Session  Session
		Products []models.Product
	}{
		Session: Session{
			UserName:          "Lucas",
			CartProductsCount: 6,
			Categories:        []string{"Notebook", "Monitor"},
		},
		Products: products,
	}

	err = tmplAdminProductList.ExecuteTemplate(w, "layout.gohtml", data)
	HandleError(w, err)
}
