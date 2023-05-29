package auth

import (
	"context"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	"github.com/DATOULIN/dtservice/pkg/token"
	"github.com/DATOULIN/dtservice/pkg/util"
	"strconv"
	"time"
)

func ParseToken(tokenStr string) (*token.Claims, int64, error) {
	opt := &token.Option{JwtSecret: setting.JwtSettings.JwtSecret}
	clams, err := token.Parse(opt, tokenStr)
	userid := clams.UserId
	return clams, userid, err
}

func GenerateToken(username string, userid int64) string {
	opt := &token.Option{
		JwtSecret: setting.JwtSettings.JwtSecret,
		Expires:   setting.JwtSettings.Expires,
		Issuer:    setting.JwtSettings.Issuer,
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
	clams, _, errs := ParseToken(tokenStr)
	//var claims token.Claims
	//tokenClaims, errs := jwt.ParseWithClaims(tokenStr, &claims, func(tokens *jwt.Token) (interface{}, error) {
	//	return setting.JwtSettings.JwtSecret, nil
	//})
	if errs != nil {
		return errs
	}
	timer := time.Duration(clams.ExpiresAt-nowUnix) * time.Second
	//timer := time.Duration(tokenClaims.Claims.(*token.Claims).ExpiresAt-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	errs = setting.Redis.Set(context.Background(), getBlackListKey(tokenStr), nowUnix, timer).Err()
	return errs
}

func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := setting.Redis.Get(context.Background(), getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	if time.Now().Unix()-joinUnix < setting.JwtSettings.JwtBlacklistGracePeriod {
		return false
	}
	return true
}

// 获取黑名单缓存 key
func getBlackListKey(token string) string {
	return "jwt_black_list:" + util.MD5([]byte(token))
}
