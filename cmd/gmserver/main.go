package main

import (
	"fmt"
	"gmserver/database"
	"gmserver/httpd/posts"
	"gmserver/httpd/users"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server started at PORT 8090")
	r := mux.NewRouter()
	client := database.Init()
	r.HandleFunc("/users", users.Create(client)).Methods("POST")
	r.HandleFunc("/users", users.GetAll(client)).Methods("GET")
	r.HandleFunc("/userspost", users.GetAllWithPost(client)).Methods("GET")
	r.HandleFunc("/users/{id}", users.Update(client)).Methods("PUT")
	r.HandleFunc("/users/{id}", users.GetOne(client)).Methods("GET")
	r.HandleFunc("/users/{id}", users.DeleteOne(client)).Methods("DELETE")
	r.HandleFunc("/posts", posts.Create(client)).Methods("POST")
	r.HandleFunc("/posts", posts.GetAll(client)).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}
