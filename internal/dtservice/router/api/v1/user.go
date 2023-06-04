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
		userR.POST("/login", user.Login)

		userAuth.POST("/logout", user.Logout)
		userAuth.GET("/list", user.GetUserList)
		userAuth.PUT("/update/:user_id", user.UpdateUser)
		userAuth.POST("/resetPassword/:user_id", user.ResetPassword)
		userAuth.POST("/uploadAvatar/:user_id", user.UploadAvatar)
		userAuth.DELETE("/delete/:user_id", user.Delete)
	}
}
