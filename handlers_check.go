package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

// Test page.
func checkPageHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params, session *Session) {
	data := struct {
		Session *Session
	}{session}

	err := tmplTest.ExecuteTemplate(w, "test.gohtml", data)
	HandleError(w, err)
}

// Test send mail.
func checkSendMailPost(w http.ResponseWriter, req *http.Request, ps httprouter.Params, _ *Session) {
	msg := time.Now().String()
	err := sendMail([]string{"douglasmg7@gmail.com"}, "Teste (zunkasrv).", msg)
	if err == nil {
		w.WriteHeader(200)
		return
	}
	HandleError(w, err)
}
