package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/koobun/information/handlers"
)

func main() {
	// Router is deciding which handler to call based on the URL
	r := mux.NewRouter()

	r.HandleFunc("/main", handlers.IndexPage)

	r.HandleFunc("/users/{username}", handlers.UserPage)

	// start the server
	http.ListenAndServe(":3000", r) // :3000 is port number (http://localhost:3000)

}

// 1.create a new handler
// 2. make this handler return text
// 	 "Hello from Zen!"
// 3. connect this handler to URL /about
