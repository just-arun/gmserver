package auth

import (
	"encoding/json"
	"fmt"
	"gmserver/database"
	"gmserver/models"
	"gmserver/pkg"
	"gmserver/pkg/password"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	APP_KEY = "golangcode.com"
)

func Login(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var userInput *models.UserModel
		err := json.NewDecoder(r.Body).Decode(&userInput)
		defer pkg.ErrCheck(w, err)

		email := userInput.Email

		// check for user
		userCountChallen := make(chan int64)
		go CheckUser(client, w, email, userCountChallen)
		count := <-userCountChallen
		if count == 0 {
			pkg.ErrWithCusMsg(w, http.StatusNotFound, "User Not Found")
			return
		}

		// get user
		var user *models.UserModel
		usrCollection, ctx := database.CollectionFun(client, database.CollectionList().Users)
		err = usrCollection.FindOne(ctx, models.UserModel{Email: userInput.Email}).Decode(&user)
		defer pkg.ErrCheck(w, err)
		fmt.Println(user)

		// compare password
		err = password.CompareHash([]byte(user.Password), []byte(userInput.Password))
		if err != nil {
			defer pkg.ErrWithCusMsg(w, http.StatusBadRequest, "username/password dosen't match")
		}

		// here, we have kept it as 5 minutes
		expirationTime := time.Now().Add(5 * time.Minute)
		token, err := GenerateToken(user.ID.Hex())
		defer pkg.ErrCheck(w, err)
		// w.Write([]byte(`{{data:{"token": ` + tokenString + ` }}}`))
		http.SetCookie(w, &http.Cookie{
			Name:    "Usr",
			Value:   token,
			Expires: expirationTime,
		})
	}
}

func CheckUser(client *mongo.Client, w http.ResponseWriter, email string, c chan int64) {
	userCllection1, ctx := database.CollectionFun(client, database.CollectionList().Users)
	count, err := userCllection1.CountDocuments(ctx, models.UserModel{Email: email})
	defer pkg.ErrCheck(w, err)
	c <- count
	close(c)
}

func GenerateToken(id string) (string, error) {
	// var user models.UserModel
	fmt.Println(id)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte("someshit"))
	panic(err)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"user": id,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
		"path": "/",
	})
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}
