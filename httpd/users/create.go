package users

import (
	"encoding/json"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg/password"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type outputData struct {
	ID interface{} `json:"id"`
}

func Create(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applcation/json")
		var user models.UserModel
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
		}
		defer r.Body.Close()
		user.ID = primitive.NewObjectID()
		hash, err := password.CreateHash([]byte(user.Password))
		if err != nil {
			panic(err)
		}
		user.Password = hash
		// saving user to database
		// collection := client.Database(database.Database).Collection(database.CollectionList().Users)
		// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
		}
		dat := outputData{ID: result.InsertedID}
		bi, err := json.Marshal(dat)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{ "message": ` + err.Error() + `"}`))
		}
		w.Write([]byte(bi))
	}
}
