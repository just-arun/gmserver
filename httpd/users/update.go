package users

import (
	"encoding/json"
	"fmt"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var err error
var filter bson.D

func Update(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(params["id"])
		pkg.ErrCheck(w, err)
		var user models.UserModel
		err = json.NewDecoder(r.Body).Decode(&user)
		pkg.ErrCheck(w, err)
		user.Password = ""
		user.ID = id

		//preparing filter, fields and option

		// updating to database
		var updatedUser models.UserModel
		filter := bson.M{"_id": id}
		query := bson.M{"$set": user}
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		err = collection.FindOneAndUpdate(ctx, filter, query).Decode(&updatedUser)
		pkg.ErrCheck(w, err)
		fmt.Println(updatedUser)
		w.WriteHeader(http.StatusOK)
	}
}
