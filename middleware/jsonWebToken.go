package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
)
const (
	SecretKey = "hsn wallet"
)
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		//token := jwt.New(jwt.SigningMethodHS256)
		//claims := make(jwt.MapClaims)
		//
		//claims["exp"] = time.Now().Add(time.Minute * time.Duration(10)).Unix()
		//claims["iat"] = time.Now().Unix()
		//claims["username"] = "wyl"
		//token.Claims = claims
		//tokenString, err := token.SignedString([]byte(SecretKey))
		//fmt.Println(tokenString)
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})
		if err == nil {
			if token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusOK)
				_,_=w.Write([]byte("NO AUTH"))
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_,_=w.Write([]byte("NO AUTH !"))
		}
		// Call the next handler, which can be another middleware in the chain, or the final handler.
	})
}
