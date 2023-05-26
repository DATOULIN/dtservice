package router

import (
	"github.com/DATOULIN/dtservice/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	//userR := r.Group("/api/v1/")
	userRAuth := r.Group("/api/v1/")
	userRAuth.Use(middleware.JwtAuth())
	userRAuth.GET("/test", func(c *gin.Context) {
		response := gin.H{"code": 0, "msg": "msg", "result": nil}
		c.JSON(http.StatusOK, response)
	})
	return r
}
