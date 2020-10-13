package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const SESSION_ID_COOKIE_NAME = "sessionid"

type Session struct {
	SessionID         string
	UserID            string `json:"userId"`
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

func saveSession(session *Session) {
	redisSetUserSession(session.UserID, session)
}

func removeSessionID(session *Session) {
	redisDelSessionIDUserID(session.SessionID)
}

func loadSession(w http.ResponseWriter, r *http.Request) (*Session, bool) {
	// Get user id from session id
	cookie, err := r.Cookie(SESSION_ID_COOKIE_NAME)
	// log.Println("Cookie:", cookie.Value)
	if err != nil && err != http.ErrNoCookie {
		log.Printf("[error] Getting cookie. %v", err)
		return &Session{}, false
	}
	userID := ""
	if cookie != nil {
		userID = redisGetSessionIDUserID(cookie.Value)
	}
	if userID == "" {
		// create new anonymous and new cookie
		sUUID := uuid.NewV4()
		sessionID := sUUID.String()
		// save cookie.
		http.SetCookie(w, &http.Cookie{
			Name:  SESSION_ID_COOKIE_NAME,
			Value: sessionID,
			Path:  "/",
			// Secure: true, // to use only in https
			// HttpOnly: true, // Can't be used into js client
		})
		// anonymous user
		userID := "@@" + sessionID
		redisSetSessionIDUserID(sessionID, userID)
		session := &Session{}
		redisSetUserSession(userID, session)
		return &Session{}, true
	}
	return redisGetUserSession(userID), true
}
