package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func configRouter(router *httprouter.Router) {
	router.GET("/favicon.ico", faviconHandler)
	router.GET("/", getSessionMidlleware(indexHandler))
	router.GET("/ping", getSessionMidlleware(indexPing))

	// Auth - signup.
	router.GET("/auth/signup", confirmNoLogged(authSignupHandler))
	router.POST("/auth/signup", confirmNoLogged(authSignupHandlerPost))
	router.GET("/auth/signup/confirmation/:uuid", confirmNoLogged(authSignupConfirmationHandler))

	// Auth - signin/signout.
	router.GET("/auth/signin", confirmNoLogged(authSigninHandler))
	router.POST("/auth/signin", confirmNoLogged(authSigninHandlerPost))
	router.GET("/auth/signout", authSignoutHandler)

	// Auth - password.
	router.GET("/auth/password/recovery", confirmNoLogged(passwordRecoveryHandler))
	router.POST("/auth/password/recovery", confirmNoLogged(passwordRecoveryHandlerPost))
	router.GET("/auth/password/reset", confirmNoLogged(passwordResetHandler))

	// Admin
	router.HandlerFunc("GET", "/admin/products", AdminProductListHandlerGet)

	// Test.
	router.GET("/ns/test", checkPermission(checkPageHandler, "admin"))
	router.POST("/ns/test/send-email", checkPermission(checkSendMailPost, "admin"))

	// User.
	router.GET("/ns/user/account", checkPermission(userAccountHandler, ""))
	router.GET("/ns/user/change/name", checkPermission(userChangeNameHandler, ""))
	router.POST("/ns/user/change/name", checkPermission(userChangeNameHandlerPost, ""))
	router.GET("/ns/user/change/email", checkPermission(userChangeEmailHandler, ""))
	router.POST("/ns/user/change/email", checkPermission(userChangeEmailHandlerPost, ""))
	router.GET("/ns/user/change/email-confirmation/:uuid", checkPermission(userChangeEmailConfirmationHandler, ""))
	router.GET("/ns/user/change/mobile", checkPermission(userChangeMobileHandler, ""))
	router.POST("/ns/user/change/mobile", checkPermission(userChangeMobileHandlerPost, ""))

	// Static
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
}
