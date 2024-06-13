package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("secret_key")

type Claims struct {
    Email string `json:"email"`
    jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {
    expirationTime := time.Now().Add(72 * time.Hour)

    claims := &Claims{
        Email: email,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(
        signedToken,
        &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        },
    )

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}

func ExtractUserEmailFromJWT(c *gin.Context) (string, error) {
    token := c.Request.Header.Get("Authorization")
    claims, err := ValidateToken(token)
    if err != nil {
        return "", err
    }
    return claims.Email, nil
}
