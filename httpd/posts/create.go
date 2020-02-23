package posts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg/postutil"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type outputData struct {
	ID interface{} `json:"id"`
}

func Create(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applcation/json")
		var post models.PostModel
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
		}
		defer r.Body.Close()
		post.ID = primitive.NewObjectID()
		post.Link = postutil.StringToLink(post.Title)

		// saving user to database
		collection, ctx := database.CollectionFun(client, database.CollectionList().Posts)
		result, err := collection.InsertOne(ctx, post)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		dat := outputData{ID: result.InsertedID}
		bi, err := json.Marshal(dat)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		fmt.Println(bi)
		userFilter := bson.M{"_id": post.Author}
		updateUser := bson.M{"$push": bson.M{"postsId": post.ID}}
		var user models.UserModel
		userCollection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		err = userCollection.FindOneAndUpdate(ctx, userFilter, updateUser).Decode(&user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		}
		fmt.Println(user)
		w.Write([]byte(bi))
	}
}
