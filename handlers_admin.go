package main

import (
	"webserver/models"

	"log"
	"net/http"
)

// Product list
func AdminProductListHandlerGet(w http.ResponseWriter, req *http.Request) {

	// Get all products
	products, err := models.GetAllProduct()
	if err != nil {
		log.Printf("err: %v", err)
	}
	// log.Printf("Products: %v", products)

	data := struct {
		Session  *Session
		Products []models.Product
	}{
		Session: &Session{
			UserName:          "",
			CartProductsCount: 6,
			ProductCategories: models.Categories,
		},
		Products: products,
	}

	err = tmplAdminProductList.ExecuteTemplate(w, "layout.gohtml", data)
	HandleError(w, err)
}
