package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"webserver/models"

	"github.com/julienschmidt/httprouter"
)

// Handle with session.
type handleS func(w http.ResponseWriter, req *http.Request, p httprouter.Params, session *Session)

// Get session middleware.
func getSessionMidlleware(h handleS) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// // Get session.
		// session := getSession(req)
		// h(w, req, p, session)

		h(w, req, p, &Session{})
	}
}

// Check permission middleware.
func checkPermission(h handleS, permission string) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// // Get session.
		// session := getSession(req)
		// // Not signed.
		// if session.UserName == "" {
		// http.Redirect(w, req, "/auth/signin", http.StatusSeeOther)
		// return
		// }
		// // Have the permission.
		// if permission == "" || session.CheckPermission(permission) {
		// h(w, req, p, session)
		// return
		// }
		// // No Permission.
		// // fmt.Fprintln(w, "Not allowed")
		// data := struct {
		// Session     *Session
		// HeadMessage string
		// }{Session: session}
		// err := tmplDeniedAccess.ExecuteTemplate(w, "deniedAccess.tpl", data)
		// HandleError(w, err)

		h(w, req, p, &Session{})
	}
}

// Check if not logged.
func confirmNoLogged(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		// // Get session.
		// session := getSession(req)
		// // Not signed.
		// if session.UserName == "" {
		// h(w, req, p)
		// return
		// }
		// // fmt.Fprintln(w, "Not allowed")
		// data := struct{ Session *Session }{session}
		// err := tmplDeniedAccess.ExecuteTemplate(w, "deniedAccess.tpl", data)
		// HandleError(w, err)

		h(w, req, p)
	}
}

// Api Authorization.
func checkApiAuthorization(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		user, pass, ok := req.BasicAuth()
		if ok && user == zunkaServerUser() && pass == zunkaServerPass() {
			h(w, req, p)
			return
		}
		log.Printf("Unauthorized access, %v %v, user: %v, pass: %v, ok: %v", req.Method, req.URL.Path, user, pass, ok)
		log.Printf("authorization      , %v %v, user: %v, pass: %v", req.Method, req.URL.Path, zunkaServerUser(), zunkaServerPass())
		// Unauthorised.
		w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password for this service"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return
	}
}

/**************************************************************************************************
* Logger
**************************************************************************************************/
// Logger struct.
type logger struct {
	handler http.Handler
}

// Handle interface.
// todo - why DELETE is logging twice?
func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// log.Printf("%s %s - begin", req.Method, req.URL.Path)
	start := time.Now()
	l.handler.ServeHTTP(w, req)
	log.Printf("%s %s %v", req.Method, req.URL.Path, time.Since(start))
	// log.Printf("header: %v", req.Header)
}

// New logger.
func newLogger(h http.Handler) *logger {
	return &logger{handler: h}
}

/**************************************************************************************************
* User
**************************************************************************************************/
type sessionMiddleware struct {
	handler http.Handler
}

func (l *sessionMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, ok := loadSession(w, r)
	if !ok {
		http.Error(w, "Error loading session", http.StatusInternalServerError)
		return
	}
	// log.Printf("sessionMiddleware called")
	session.ProductCategories = models.Categories
	session.CartProductsCount = 8
	ctx := r.Context()
	ctx = context.WithValue(ctx, "session", session)
	r = r.WithContext(ctx)

	l.handler.ServeHTTP(w, r)
}

func newSessionMiddleware(h http.Handler) *sessionMiddleware {
	return &sessionMiddleware{handler: h}
}
