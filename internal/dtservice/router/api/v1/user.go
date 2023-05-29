package v1

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/controller"
	"github.com/DATOULIN/dtservice/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {
	userR := r.Group("/api/v1/user")
	userAuth := r.Group("/api/v1/user")
	userAuth.Use(middleware.JwtAuth())
	user := controller.NewUser()
	{
		userR.POST("/register", user.Register)
	}
}
