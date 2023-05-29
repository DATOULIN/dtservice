package setting

import (
	"github.com/DATOULIN/dtservice/pkg/db"
	"github.com/DATOULIN/dtservice/pkg/vp"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var (
	Redis            *redis.Client
	DBEngine         *gorm.DB
	ServerSettings   *ServerSetting
	DatabaseSettings *DatabaseSetting
	AppSettings      *AppSetting
	JwtSettings      *JwtSetting
	RedisSettings    *RedisSetting
)

// SetupSetting 初始化配置
func SetupSetting() {
	opt := vp.Options{
		ConfigName: "dtservice",
		ConfigPath: "configs/",
		ConfigType: "yaml",
	}
	newSetting, err := vp.NewViper(&opt)
	err = newSetting.ReadSection("Server", &ServerSettings)
	err = newSetting.ReadSection("Database", &DatabaseSettings)
	err = newSetting.ReadSection("App", &AppSettings)
	err = newSetting.ReadSection("Jwt", &JwtSettings)
	err = newSetting.ReadSection("Redis", &RedisSettings)
	ServerSettings.ReadTimeout *= time.Second
	ServerSettings.WriteTimeout *= time.Second
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
}

// SetupDBEngine 初始化数据库
func SetupDBEngine() {
	dbOptions := &db.MySQLOptions{
		DBType:       DatabaseSettings.DBType,
		Username:     DatabaseSettings.Username,
		Password:     DatabaseSettings.Password,
		Host:         DatabaseSettings.Host,
		DBName:       DatabaseSettings.DBName,
		TablePrefix:  DatabaseSettings.TablePrefix,
		Charset:      DatabaseSettings.Charset,
		ParseTime:    DatabaseSettings.ParseTime,
		MaxIdleConns: DatabaseSettings.MaxIdleConns,
		MaxOpenConns: DatabaseSettings.MaxOpenConns,
	}
	ins, err := db.NewDBEngine(dbOptions)
	DBEngine = ins
	if ServerSettings.RunMode == "debug" {
		ins.LogMode(true)
	}
	if err != nil {
		log.Fatalf("init.SetupDBEngine err:%v", err)
	}
}

// SetUpRedis 初始化redis
func SetUpRedis() {
	redisOptions := &db.RedisOptions{
		Host:     RedisSettings.Host,
		Port:     RedisSettings.Port,
		DB:       RedisSettings.DB,
		Password: RedisSettings.Password,
	}
	ins, err := db.NewRedis(redisOptions)
	Redis = ins
	if err != nil {
		log.Fatalf("init.SetUpRedis err:%v", err)
	}
}
