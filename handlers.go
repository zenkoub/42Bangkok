package handlers

import (
	"net/http"

	"log"

	"github.com/gorilla/mux"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

var users = map[string]map[string]any{
	"zen": {
		"Name": "zen",
		"City": "Bangkok",
	},

	"john": {
		"Name": "john",
		"City": "Tokyo",
	},

	"david": {
		"Name": "david",
		"City": "Alabama",
	},
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, exists := users[username]
	log.Printf("User: %v, exists: %v", user, exists)

	if !exists {
		http.NotFound(w, r)
		return
	}

	RenderTemplate(w, "user.html", user)
}
