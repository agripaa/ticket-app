package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jeypc/go-jwt-mux/config"
	"github.com/jeypc/go-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				res := map[string]string{"status": "401", "message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, res)
				return
			}
		}
		tokenString := c.Value
		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				res := map[string]string{"status": "401", "message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, res)
				return
			case jwt.ValidationErrorExpired:
				res := map[string]string{"status": "401", "message": "Token expired"}
				helper.ResponseJSON(w, http.StatusUnauthorized, res)
				return
			default:
				res := map[string]string{"status": "401", "message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, res)
				return
			}
		}
		if !token.Valid {
			res := map[string]string{"status": "401", "message": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, res)
			return
		}
		next.ServeHTTP(w, r)
	})
}
