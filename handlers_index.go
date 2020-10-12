package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type valueMsg struct {
	Value string
	Msg   string
}

// Template message data.
type messageTplData struct {
	Session    *Session
	TitleMsg   string
	WarnMsg    string
	SuccessMsg string
}

// Handler error.
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		// http.Error(w, "Some thing wrong", 404)
		if production {
			http.Error(w, "Alguma coisa deu errado", http.StatusInternalServerError)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		log.Println(err.Error())
		return
	}
}

// Favicon handler.
func faviconHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "./static/img/favicon.ico")
}

// Index handler.
func indexHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session     *Session
		HeadMessage string
	}{session, "Aviso de regatta na Lagoa dos Ingleses, dia 18/03/2019"}
	// fmt.Println("session: ", data.Session)
	err := tmplIndex.ExecuteTemplate(w, "index.tpl", data)
	HandleError(w, err)
}

// Index handler.
func indexPing(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	w.Write([]byte("Pong"))
}

/**************************************************************************************************
* To organizer
**************************************************************************************************/

func userHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER, %s!\n", ps.ByName("name"))
}

func userAddHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tmplUserAdd.ExecuteTemplate(w, "user_add.tpl", nil)
	HandleError(w, err)
}
