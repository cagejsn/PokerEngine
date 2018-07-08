package main


import "net/http"


// UserHandler is like an http.Handler but also takes a User.
type UserHandler interface {
	UserServeHTTP(User, http.ResponseWriter, *http.Request)
}

// The UserHandlerFunc type is an adapter to allow the use of
// ordinary functions as user handlers.
type UserHandlerFunc func(User, http.ResponseWriter, *http.Request)

// UserHandlerFunc implements the UserHandler interface.
func (f UserHandlerFunc) UserServeHTTP(u User, w http.ResponseWriter, r *http.Request) {
	f(u, w, r)
}





// HandlerFrom creates an http.Handler from a UserHandler. A fallback is
// provided for cases where the client is not authenticated.
func HandlerFrom(uh UserHandler, fallback http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)
		if(r.Method == http.MethodOptions){
			w.Write([]byte("good"))
			return
		}

		u, err := checkUser(r)  // undefined
		if err != nil {
			fallback.ServeHTTP(w, r)
			return
		}

		// ...maybe check some more things about the user here

		uh.UserServeHTTP(u, w, r)
	})
}