package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gmserver/database"
	"gmserver/models"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func GetAll(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []*models.UserModel
		filter := bson.M{}
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
			panic(err)
		}
		for cursor.Next(ctx) {
			var user models.UserModel
			err := cursor.Decode(&user)
			if err != nil {
				panic(err)
			}
			user.Password = ""
			users = append(users, &user)
		}
		cursor.Close(ctx)
		fmt.Println(users)
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(users)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
			panic(err)
		}
		w.Write(jsonData)
	}
}
