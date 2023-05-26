package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserName string `json:"user_name"`
	UserId   int64  `json:"user_id"`
	jwt.StandardClaims
}

type Option struct {
	key       interface{}
	JwtSecret string
	Expires   int64
	Issuer    string
}

var SECRET = "1234567890"

// Sign   签发 token
func Sign(opts *Option, userName string, userId int64) (string, error) {
	claims := Claims{
		userName,
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + opts.Expires, //过期时间
			Issuer:    opts.Issuer,                      //签发人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 用签名来 签名这个token
	token, err := tokenClaims.SignedString([]byte(opts.JwtSecret))
	return token, err
}

// Parse 验证token
func Parse(opts *Option, token string) (*Claims, error) {
	var claims Claims
	tokenClaims, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(opts.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}
