package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jdrada/go-auth-v1/api/model"
	"github.com/jdrada/go-auth-v1/db"

	"github.com/jdrada/go-auth-v1/utils"
)


func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	var errUser model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	db.DB.Where("email = ?", user.Email).First(&errUser)

	if errUser.ID == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	if !utils.CheckPasswordHash(user.Password, errUser.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Incorrect password")
		return
	}

	token, err := utils.GenerateJWT(errUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error generating token")
		return
	}

	json.NewEncoder(w).Encode(token)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello Protected")
}

