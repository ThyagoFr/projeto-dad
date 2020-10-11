package security

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware - Check if jwt is inside the request header
func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Rodando middleware ...")
		if err := extractToken(r.Header); err != nil {
			// fazer alguma coisa kk
			log.Println("Deu ruim...")
		}
		token, _ := generateToken(1)
		// w.Header().Add("token", token)
		next.ServeHTTP(w, r)
	})

}

func extractToken(header http.Header) error {
	authorization := header["Authorization"]
	if authorization == nil {
		log.Println("Sem o token...")
		return errors.New("Um error x")
	}
	token := strings.Split(authorization[0], " ")[1]
	log.Println(token)
	return nil
}

func generateToken(userID uint64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["expiration_time"] = time.Now().Add(time.Hour * 8).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("asdp978tqyfuh3ua8s7d"))
	if err != nil {
		return "", err
	}
	return token, nil
}
