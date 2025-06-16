package main

import (
	"errors"
	"net/http"
)

var AuthError = errors.New("Auth Error")

func Authorize(r *http.Request) error {
	//username := r.FormValue("username")
	//user, ok := users[username]
	//if !ok {
	//	//panic("User not found " + " " + username)
	//	return AuthError
	//}
	//
	//st, err := r.Cookie("funkycms_session")
	//if err != nil || st.Value == "" || st.Value != user.SessionToken {
	//	//panic("Session token issue " + user.SessionToken)
	//	return AuthError
	//}
	//
	//csrf := r.Header.Get("X-CSRF-Token")
	//if csrf != user.CSRFToken || csrf == "" {
	//	//panic("CSRF token issue " + user.CSRFToken)
	//	return AuthError
	//}

	return nil
}
