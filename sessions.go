package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const SESSION_ID_COOKIE_NAME = "session_id"

type Session struct {
	UserID            int    `json:"userId"`
	UserName          string `json:"userName"`
	Permission        uint64 `json:"userPermission"`
	CartProductsCount uint   `json:"cartProductsCount"`
	ProductCategories []string
}

// CheckPermission return if session has the permission.
func (s *Session) CheckPermission(p string) bool {
	// Admin.
	if s.Permission&1 == 1 {
		return true
	}
	// Permissions.
	switch p {
	case "write":
		return s.Permission&2 == 2
	case "read":
		return s.Permission&4 == 4
	default:
		return false
	}
}

// SetPermission grand permission.
func (s *Session) SetPermission(p string) {
	switch p {
	case "admin":
		s.Permission = s.Permission | 1
	case "write":
		s.Permission = s.Permission | 2
	case "read":
		s.Permission = s.Permission | 4
	}
}

// UnsetPermission revoke permission.
func (s *Session) UnsetPermission(p string) {
	switch p {
	case "admin":
		s.Permission = s.Permission ^ 1
	case "write":
		s.Permission = s.Permission ^ 2
	case "read":
		s.Permission = s.Permission ^ 4
	}
}

func createSession(w http.ResponseWriter, userID string) {
	// create cookie
	sUUID := uuid.NewV4()
	sessionID := sUUID.String()
	// Save cookie.
	http.SetCookie(w, &http.Cookie{
		Name:  SESSION_ID_COOKIE_NAME,
		Value: sessionID,
		Path:  "/",
		// Secure: true, // to use only in https
		// HttpOnly: true, // Can't be used into js client
	})
	redisSetUserID(sessionID, userID)
}

func saveSession(userID string, session *Session) {
	redisSetSession(userID, session)
}

func removeSession(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(SESSION_ID_COOKIE_NAME)
	if err == http.ErrNoCookie {
		// No cookie.
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else if err != nil {
		// Some error.
		log.Printf("[error] Getting cookie. %v", err)
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		// Remove cookie.
		c.MaxAge = -1
		c.Path = "/"
		// log.Println("changed cookie:", c)
		http.SetCookie(w, c)
		http.Redirect(w, req, "/ns/auth/signin", http.StatusSeeOther)
		redisDelUserID(c.Value)
	}
}

func getSession(req *http.Request) *Session {
	// timeToGetSession = time.Now()
	userID := getUserIDFromSessionID(req)
	if userID == "" {
		return &Session{}
	}
	return redisGetSession(userID)
}

// Return user id from session uuid.
// Try the cache first.
func getUserIDFromSessionID(req *http.Request) string {
	cookie, err := req.Cookie(SESSION_ID_COOKIE_NAME)
	// log.Println("Cookie:", cookie.Value)
	// log.Println("Cookie-err:", err)
	// No cookie.
	if err == http.ErrNoCookie {
		return ""
		// some error
	} else if err != nil {
		return ""
	}
	// Have a cookie.
	if cookie != nil {
		return redisGetUserID(cookie.Value)
	}
	// No cookie
	return ""
}
