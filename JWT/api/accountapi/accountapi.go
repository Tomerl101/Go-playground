package accountapi

import (
	"encoding/json"
	"net/http"
	"time"

	"../../model"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = "mysecretkey"

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account model.Account
	// var err error

	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		//create the jwt token with user payload and expire time
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": account.UserName,
			"password": account.Password,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err2 := token.SignedString([]byte(secretKey))
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, model.JWTToken{Token: tokenString})
		}
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	respondWithJson(w, statusCode, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload) //JSON.stringify
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

//maps are like dicts / hashes
// map[<key type>]<value type>
