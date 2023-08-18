package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jeypc/go-jwt-mux/config"
	"github.com/jeypc/go-jwt-mux/helper"
	"github.com/jeypc/go-jwt-mux/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		res := map[string]string{"status": "400", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}
	defer r.Body.Close()
	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res := map[string]string{"status": "401", "message": "user not found"}
			helper.ResponseJSON(w, http.StatusUnauthorized, res)
			return
		default:
			res := map[string]string{"status": "500", "message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, res)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		res := map[string]string{"status": "401", "message": "user not found"}
		helper.ResponseJSON(w, http.StatusUnauthorized, res)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claim := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := tokenAgo.SignedString(config.JWT_KEY)
	if err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, res)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	res := map[string]string{"status": "200", "message": "login successfully"}
	helper.ResponseJSON(w, http.StatusOK, res)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		res := map[string]string{"status": "400", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := models.DB.Create(&userInput).Error; err != nil {
		res := map[string]string{"status": "500", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, res)
		return
	}

	res := map[string]string{"status": "200", "message": "success"}
	helper.ResponseJSON(w, http.StatusOK, res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	res := map[string]string{"status": "200", "message": "logout successfully"}
	helper.ResponseJSON(w, http.StatusOK, res)
}
