package users

import (
	"encoding/json"
	"fmt"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func GetAll(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		users := GetAllUserService(client, w)
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(users)
		pkg.ErrCheck(w, err)
		w.Write(jsonData)
	}
}

func GetAllUserService(client *mongo.Client, w http.ResponseWriter) []*models.UserModel {
	var users []*models.UserModel
	filter := bson.M{}
	collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
	cursor, err := collection.Find(ctx, filter)
	pkg.ErrCheck(w, err)
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
	return users
}
