package setting

import "time"

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseSetting struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileExt      string
	CaptchaType     string
	StaticDir       string
	FileDir         string
}

type JwtSetting struct {
	JwtSecret               string
	Expires                 int64
	Issuer                  string
	JwtBlacklistGracePeriod int64
}

type RedisSetting struct {
	Host     string
	Port     string
	DB       int
	Password string
}
