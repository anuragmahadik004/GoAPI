package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/anuragmahadik004/hr_api/datalayer"
	encryptdecrypt "github.com/anuragmahadik004/hr_api/encrypt_decrypt"
	"github.com/anuragmahadik004/hr_api/interfaces"
	"github.com/anuragmahadik004/hr_api/models"
	"github.com/gorilla/mux"
)

var _UserRepository interfaces.UserRepository

func GetUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	User := _UserRepository.GetUser(params["userId"])

	json.NewEncoder(w).Encode(User)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var User models.User

	_ = json.NewDecoder(r.Body).Decode(&User)

	UserRetrieved := _UserRepository.GetUser(User.UserName)

	if len(UserRetrieved.UserId) > 0 {

		json.NewEncoder(w).Encode("Username already exist choose different username")

	} else {

		UserSaved := _UserRepository.SaveUser(User)

		json.NewEncoder(w).Encode(UserSaved)

	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var UserLogin models.UserLogin

	_ = json.NewDecoder(r.Body).Decode(&UserLogin)

	User := _UserRepository.GetUser(UserLogin.UserName)

	if len(User.UserId) == 0 {

		json.NewEncoder(w).Encode("User Does Not Exist")

	} else {

		if UserLogin.Password == User.Password {

			User.Password, _ = encryptdecrypt.Encrypt(User.Password)

			jsonData, _ := json.Marshal(User)

			fmt.Println(string(jsonData))

			encryptedJWT, _ := encryptdecrypt.Encrypt(string(jsonData))

			w.Header().Set("JWTToken", encryptedJWT)

		} else {
			json.NewEncoder(w).Encode("Invalid Password")
		}

	}
}

func VerifyJWTToken(JWTToken string) bool {

	var User models.User

	jsonDecrypted, err := encryptdecrypt.Decrypt(JWTToken)

	if err != nil {
		return false
	}

	err = json.Unmarshal([]byte(jsonDecrypted), &User)

	if err != nil {
		return false
	}

	User.Password, _ = encryptdecrypt.Decrypt(User.Password)

	UserRetrieved := _UserRepository.GetUser(User.UserName)

	if len(UserRetrieved.UserId) == 0 {
		return false
	} else {
		if User.UserName == UserRetrieved.UserName && User.Password == UserRetrieved.Password {
			return true
		} else {
			return false
		}
	}

}
