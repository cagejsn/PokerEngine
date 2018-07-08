package main

import (
	"net/http"
	"github.com/gorilla/securecookie"
	"time"
	"strings"
)

const userCookie = "something-obscure"

var s = securecookie.New([]byte("very-secret"),  // hash key
	nil) // block key


func checkUser(r *http.Request) (User, error) {

	//u := User{}

	//cookie, err := r.Cookie(userCookie)

	token := r.Header.Get("Authorization")
	sections := strings.Split(token," ")


	session, err := retrieveUserSession(DBCon,sections[1])
	if err != nil {

	}

	//if err != nil {
	//	return u, err
	//}
	//fmt.Print(cookie)
	//err = s.Decode(userCookie, cookie.Value, &u)


	return session.user, err
}


func setUserSessionCookie(session UserSession, w http.ResponseWriter){

	if encoded, err := s.Encode(userCookie, session.sessionKey ); err == nil {

		expiration := time.Now().Add(365 * 24 * time.Hour)

		cookie := &http.Cookie{
			Name:  userCookie,
			Value: encoded,
			Path:  "/",
			Domain: "127.0.0.1",
			HttpOnly:true,
			Expires:expiration,
		}

		http.SetCookie(w, cookie)
	}
}


func checkUserSessionCookie(r *http.Request) (UserSession, error){

	var userSessionKey string

	cookie, err := r.Cookie(userCookie)
	err = s.Decode(userCookie, cookie.Value, &userSessionKey)
	u, err := retrieveUserSession(DBCon,userSessionKey)

	return u, err
}

