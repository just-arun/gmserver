package users

import (
	"encoding/json"
	"net/http"

	"gmserver/database"
	"gmserver/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOne(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var user models.UserModel
		id, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		// creating filter
		filter := models.UserModel{ID: id}
		filterOptions := options.FindOne()
		filterOptions.SetAllowPartialResults(false)
		//search for data in database
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		err = collection.FindOne(ctx, filter, filterOptions).Decode(&user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		user.Password = ""
		resu, err := json.Marshal(user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resu)
	}
}
