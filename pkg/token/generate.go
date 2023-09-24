package token

import (
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

func Generate(claims interface{}, sigKey []byte, expiredTime time.Duration) (string, error) {
	token, err := jwt.Sign(jwt.HS256, sigKey, claims, jwt.MaxAge(expiredTime))

	return string(token), err
}
