package routes

import (
	"gmserver/httpd/users"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Users(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("", users.GetAll(client)).Methods("GET")
	r.HandleFunc("", users.Create(client)).Methods("POST")
	r.HandleFunc("/all/posts", users.GetAllWithPost(client)).Methods("GET")
	r.HandleFunc("/{id}", users.Update(client)).Methods("PUT")
	r.HandleFunc("/{id}", users.GetOne(client)).Methods("GET")
	r.HandleFunc("/{id}", users.DeleteOne(client)).Methods("DELETE")
}
