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

func GetAllWithPost(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var users []*models.UserModel
		lookingup := bson.M{"$lookup": bson.M{"from": database.CollectionList().Posts, "localField": "postsId", "foreignField": "_id", "as": "postss"}}
		// unwinding := bson.M{"$unwind": bson.M{"path": "$postsId", "preserveNullAndEmptyArrays": false}}
		grouping := bson.M{"$group": bson.M{"_id": "$_id", "email": bson.M{"$first": "$email"}, "name": bson.M{"$first": "$name"}, "posts": bson.M{"$first": "$postss"}}}

		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		cursor, err := collection.Aggregate(ctx, []bson.M{lookingup, grouping})
		pkg.ErrCheck(w, err)
		for cursor.Next(ctx) {
			var user models.UserModel
			err := cursor.Decode(&user)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
			}
			user.Password = ""
			users = append(users, &user)
		}
		cursor.Close(ctx)
		fmt.Println(users)
		w.WriteHeader(http.StatusOK)
		jsonData, err := json.Marshal(users)
		pkg.ErrCheck(w, err)
		w.Write(jsonData)
	}
}
