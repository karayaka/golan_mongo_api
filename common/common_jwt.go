package common

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type CustomJWT struct {
	Name    string
	Surname string
	Email   string
	jwt.RegisteredClaims
}

func GetSession(c echo.Context) (*CustomJWT, error) {
	c.Response().Header().Add("Vary", "Authorization")
	authHeader := c.Request().Header.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer") {
		return nil, nil
	}
	tokens := strings.Split(authHeader, " ")
	if len(tokens) < 2 {
		return nil, nil
	}
	return ParseJWT(tokens[1])
}

func ParseJWT(accesToken string) (*CustomJWT, error) {

	parsedAccesToken, err := jwt.ParseWithClaims(accesToken, &CustomJWT{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET_KEY")), nil
	}, jwt.WithLeeway(5*time.Second))
	if !parsedAccesToken.Valid {
		return nil, nil
	}
	if err != nil {
		return nil, err
	} else if claims, ok := parsedAccesToken.Claims.(*CustomJWT); ok {
		return claims, nil
	} else {
		return nil, nil
	}
}
