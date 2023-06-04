package router

import (
	v1 "github.com/DATOULIN/dtservice/internal/dtservice/router/api/v1"
	"github.com/DATOULIN/dtservice/internal/pkg/middleware"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.StaticFS("/static", http.Dir(setting.AppSettings.StaticDir))
	v1.InitUserRouter(r)
	return r
}
