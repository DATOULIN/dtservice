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
	userR := r.Group("/api/v1/")
	userRAuth := r.Group("/api/v1/")
	userRAuth.Use(middleware.JwtAuth())

	user := v1.NewUser()

	userR.POST("/register", user.Register)

	return r
}
