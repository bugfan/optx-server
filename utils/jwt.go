package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"optx-server/settings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret string

func init() {
	secret := settings.Get("jwt_secret")
	if secret != "" {
		return
	}
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	secret = string(b)
}

func GenJWT(attr map[string]string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":       "bar",
		"nbf":       time.Now().Unix(),
		"attribute": attr,
	})
	jwtString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtString, nil
}
func VerifyJWT(jwtString string) (map[string]string, error) {
	token, err := ParseToken(jwtString, secret)
	if err != nil {
		return nil, err
	}
	if token["attribute"] == nil {
		return nil, err
	}

	if attr, ok := token["attribute"].(map[string]interface{}); ok {
		res := make(map[string]string)
		for k, v := range attr {
			if vv, ok := v.(string); ok {
				res[k] = vv
			}
		}
		return res, nil
	}
	return nil, errors.New("cant parse attribute")
}

func ParseToken(jwtString, t string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(t), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("jwt validate fail")
}
