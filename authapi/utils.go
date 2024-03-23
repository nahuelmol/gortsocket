package auth

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

func GenerateAccessKey() (string, error){
    claims := jwt.MapClaims{}
    claims["sub"] = "subject"
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    secret := []byte("secret-key")
    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "not signed token generated", err
    }

    return tokenString, nil
}
