package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Define your JWT secret key
var jwtSecret = []byte("your-secret-key")

// GenerateJWT generates a new JWT token for the provided user ID
func GenerateJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	return token.SignedString(jwtSecret)
}

// AuthenticateMiddleware is the middleware function to validate the JWT token
func AuthenticateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GenerateJWT("101")
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authorizationHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		fmt.Println("error:", err)

		if err != nil || !token.Valid {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// You can retrieve the user ID from the token claims and perform further authorization checks if needed

		// Pass the request to the next handler if the token is valid
		next.ServeHTTP(w, r)
	})
}
