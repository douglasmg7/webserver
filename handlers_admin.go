package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
// FILTERS
///////////////////////////////////////////////////////////////////////////////////////////////////
// Get filters.
func AdminProductListHandlerGet(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	data := struct {
		Session Session
	}{
		Session: Session{
			UserName:          "Lucas",
			CartProductsCount: 6,
			Categories:        []string{"Notebook", "Monitor"},
		},
	}

	err := tmplAdminProductList.ExecuteTemplate(w, "layout.gohtml", data)
	HandleError(w, err)
}
