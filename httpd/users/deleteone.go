package users

import (
	"fmt"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOne(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		pkg.ErrCheck(w, err)
		fmt.Println(id)
		filter := models.UserModel{ID: id}
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		result, err := collection.DeleteOne(ctx, filter)
		pkg.ErrCheck(w, err)
		fmt.Println(result.DeletedCount)
		w.WriteHeader(http.StatusOK)
	}
}
