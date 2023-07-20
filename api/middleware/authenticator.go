package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jdrada/go-auth-v1/api/model"
)

// JwtVerify verifies the token client provided
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token") // Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Not Authorized")
			return
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "Not Authorized")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			user := &model.User{}
			user.ID = uint(claims["user_id"].(float64))

			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx)) //proceed in the middleware chain!
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Not Authorized")
		}
	})
}
