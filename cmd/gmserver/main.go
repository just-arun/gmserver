package main

import (
	"fmt"
	"gmserver/database"
	"gmserver/httpd/posts"
	"gmserver/routes"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	PORT := ":8090"
	fmt.Println(`Server started at http://localhost` + PORT)
	r := mux.NewRouter()
	client := database.Init()

	// regestering routes
	user := r.PathPrefix("/users").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	routes.Users(user, client)
	routes.Auth(auth, client)
	r.HandleFunc("/posts", posts.Create(client)).Methods("POST")
	r.HandleFunc("/posts", posts.GetAll(client)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}
