package dtservice

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/router"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewDtServiceCommand() {
	setting.SetupSetting()
	setting.SetupDBEngine()
	setting.SetUpRedis()

	// 设置 Gin 模式
	gin.SetMode(setting.ServerSettings.RunMode)
	newRouter := router.NewRouter()
	s := http.Server{
		Addr:           ":" + setting.ServerSettings.HttpPort,
		Handler:        newRouter,
		ReadTimeout:    setting.ServerSettings.ReadTimeout,
		WriteTimeout:   setting.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("init err:%v", err)
		return
	}
}
