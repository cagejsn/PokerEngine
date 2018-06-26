package main

import (
	"net/http"
	"fmt"
	"github.com/gorilla/securecookie"
)

const userCookie = "something-obscure"

var s = securecookie.New([]byte("very-secret"),  // hash key
	[]byte("a-lot-secret")) // block key


func checkUser(r *http.Request) (User, error) {
	u := User{}
	cookie, err := r.Cookie(userCookie)
	if err != nil {
		return u, err
	}
	fmt.Print(cookie)
	err = s.Decode(userCookie, cookie.Value, &u)
	return u, err
}


func setUserCookie(user User, w http.ResponseWriter){

	if encoded, err := s.Encode(userCookie, user); err == nil {

		cookie := &http.Cookie{
			Name:  userCookie,
			Value: encoded,
			Path:  "/",
			Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}
}


