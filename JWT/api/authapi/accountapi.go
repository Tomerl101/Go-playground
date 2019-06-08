package authapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../../model"
	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "mysecretkey"
)

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account model.Account

	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		//we can also check if the struct fields are empty...
	} else {
		fmt.Println(account.UserName)
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

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//validate the alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//cb send the secret key to Parser()
		return []byte(secretKey), nil
	})
	return token, err
}

//claims are the token payload
func GetJwtClaims(response http.ResponseWriter, request *http.Request) {
	tokenString := request.Header.Get("token")

	if tokenString == "" {
		respondWithError(response, http.StatusOK, "token haeder is empty")
		return
	}

	token, err := ValidateToken(tokenString)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["password"])
		respondWithJson(response, http.StatusOK, claims)
	} else {
		fmt.Println(err)
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
