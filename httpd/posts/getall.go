package posts

import (
	"encoding/json"
	"gmserver/database"
	"gmserver/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func GetAll(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var posts []*models.PostModel
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
			var post models.PostModel
			err := cursor.Decode(&post)
			if err != nil {
				panic(err)
			}
			posts = append(posts, &post)
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(posts)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
			panic(err)
		}
		// w.Write(jsonData)
	}
}
