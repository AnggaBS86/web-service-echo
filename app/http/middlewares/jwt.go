package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user_id").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	username := claims.Username
	userId := claims.UserId
	return c.String(http.StatusOK, fmt.Sprintf("Welcome %s:%d", username, userId))
}

func CustomClaims() echo.MiddlewareFunc {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtClaims)
		},
		SigningKey: []byte(secretKey),
	}

	return echojwt.WithConfig(config)
}

func GetToken(userId int, username string) (string, error) {
	claims := &JwtClaims{
		userId,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
