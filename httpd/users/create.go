package users

import (
	"encoding/json"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg"
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
		err := json.NewDecoder(r.Body).Decode(&user)
		defer pkg.ErrCheck(w, err)

		// checking for existing user
		collection1, ctx1 := database.CollectionFun(client, database.CollectionList().Users)
		existUser, err := collection1.CountDocuments(ctx1, models.UserModel{Email: user.Email})
		defer pkg.ErrCheck(w, err)
		if existUser > 0 {
			defer pkg.ErrWithCusMsg(w, http.StatusBadRequest, "user already exist")
			return
		}

		defer r.Body.Close()
		user.ID = primitive.NewObjectID()
		hash, err := password.CreateHash([]byte(user.Password))
		defer pkg.ErrCheck(w, err)
		user.Password = hash
		collection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		result, err := collection.InsertOne(ctx, user)
		defer pkg.ErrCheck(w, err)
		dat := outputData{ID: result.InsertedID}
		bi, err := json.Marshal(dat)
		defer pkg.ErrCheck(w, err)
		w.Write([]byte(bi))
	}
}
