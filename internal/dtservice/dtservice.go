package dtservice

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/router"
	"github.com/DATOULIN/dtservice/internal/pkg/helper"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewDtServiceCommand() {
	helper.SetupSetting()
	helper.SetupDBEngine()
	helper.SetUpRedis()

	// 设置 Gin 模式
	gin.SetMode(helper.ServerSettings.RunMode)
	newRouter := router.NewRouter()
	s := http.Server{
		Addr:           ":" + helper.ServerSettings.HttpPort,
		Handler:        newRouter,
		ReadTimeout:    helper.ServerSettings.ReadTimeout,
		WriteTimeout:   helper.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("init err:%v", err)
		return
	}
}
