package helpers

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//TokenGenerate gera token jwt
func TokenGenerate() string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "d2js898ilsje6272g726g072gso",
		IssuedAt:  time.Now().Unix(),
	}

	// Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	// token -> string. Only server knows this secret (wsitesb).
	tokenstring, err := token.SignedString([]byte("wsitesb"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}
