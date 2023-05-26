package auth

import (
	"context"
	"github.com/DATOULIN/dtservice/internal/dtservice/helper"
	"github.com/DATOULIN/dtservice/pkg/token"
	"github.com/DATOULIN/dtservice/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func ParseToken(tokenStr string) (*token.Claims, error) {
	opt := &token.Option{JwtSecret: helper.JwtSettings.JwtSecret}
	clams, err := token.Parse(opt, tokenStr)
	return clams, err
}

func GenerateToken(username string, userid int64) string {
	opt := &token.Option{
		JwtSecret: helper.JwtSettings.JwtSecret,
		Expires:   helper.JwtSettings.Expires,
		Issuer:    helper.JwtSettings.Issuer,
	}
	tokenStr, err := token.Sign(opt, username, userid)
	if err != nil {
		return ""
	}
	return tokenStr
}

// JoinBlackList token 加入黑名单
func JoinBlackList(tokenStr string) (err error) {
	nowUnix := time.Now().Unix()
	var claims token.Claims
	tokenClaims, errs := jwt.ParseWithClaims(tokenStr, &claims, func(tokens *jwt.Token) (interface{}, error) {
		return helper.JwtSettings.JwtSecret, nil
	})
	if errs != nil {
		return errs
	}
	timer := time.Duration(tokenClaims.Claims.(*token.Claims).ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	errs = helper.Redis.Set(context.Background(), getBlackListKey(tokenStr), nowUnix, timer).Err()
	return errs
}

func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := helper.Redis.Get(context.Background(), getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	if time.Now().Unix()-joinUnix < helper.JwtSettings.JwtBlacklistGracePeriod {
		return false
	}
	return true
}

// 获取黑名单缓存 key
func getBlackListKey(token string) string {
	return "jwt_black_list:" + util.MD5([]byte(token))
}
