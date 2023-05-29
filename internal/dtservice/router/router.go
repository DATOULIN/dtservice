package router

import (
	v1 "github.com/DATOULIN/dtservice/internal/dtservice/router/api/v1"
	"github.com/DATOULIN/dtservice/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	v1.InitUserRouter(r)

	return r
}
