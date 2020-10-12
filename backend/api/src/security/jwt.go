package security

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"ufc.com/dad/src/handler"
)

// Claims - Claims
type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

// JWTMiddleware - Check if jwt is inside the request header
func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := extractToken(r.Header); err != nil {
			handler.Handler(w, r, http.StatusUnauthorized, err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})

}

func extractToken(header http.Header) error {

	authorization := header["Authorization"]
	if authorization == nil {
		return errors.New("Authorization token not found")
	}
	token := strings.Split(authorization[0], " ")[1]
	tkn, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		jwtKey, _ := os.LookupEnv("JWT_KEY")
		return []byte(jwtKey), nil
	})
	return validateToken(tkn, err)

}

func validateToken(tkn *jwt.Token, err error) error {

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("Invalid signature")
		}
		return errors.New("Invalid token")
	}
	if !tkn.Valid {
		return errors.New("Expirated token")
	}
	return nil

}

// GenerateToken - GenerateToken
func GenerateToken(userID uint) (string, error) {

	var err error
	key, _ := os.LookupEnv("JWT_KEY")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["expiration_time"] = time.Now().Add(time.Hour * 8).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	token, err := at.SignedString([]byte(key))
	if err != nil {
		return "An error ocurred", err
	}
	return token, nil

}
