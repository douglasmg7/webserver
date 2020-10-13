package main

import (
	"webserver/models"

	"log"
	"net/http"
)

// Product list
func AdminProductListHandlerGet(w http.ResponseWriter, r *http.Request) {

	// user := r.Context().Value("user").(string)
	// log.Printf("The user is: %v", user)

	session := r.Context().Value("session").(*Session)

	// Get all products
	products, err := models.GetAllProduct()
	if err != nil {
		log.Printf("err: %v", err)
	}
	// log.Printf("Products: %v", products)

	// data := struct {
	// Session  *Session
	// Products []models.Product
	// }{
	// Session: &Session{
	// UserName:          "",
	// CartProductsCount: 6,
	// ProductCategories: models.Categories,
	// },
	// Products: products,
	// }

	data := struct {
		Session  *Session
		Products []models.Product
	}{
		Session:  session,
		Products: products,
	}

	err = tmplAdminProductList.ExecuteTemplate(w, "layout.gohtml", data)
	HandleError(w, err)
}
