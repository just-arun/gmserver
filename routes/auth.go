package routes

import (
	"gmserver/httpd/auth"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Auth(r *mux.Router, client *mongo.Client) {
	r.HandleFunc("/login", auth.Login(client)).Methods("POST")
}
