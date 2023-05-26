package v1

import (
	"github.com/DATOULIN/dtservice/internal/pkg/app"
	"github.com/DATOULIN/dtservice/internal/pkg/errno"
	v1 "github.com/DATOULIN/dtservice/pkg/api/dtservice/v1"
	"github.com/DATOULIN/dtservice/pkg/validate"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// Register 注册
func (u User) Register(c *gin.Context) {
	param := v1.RegisterRequest{}
	response := app.NewResponse(c)
	valid, errs := validate.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 业务处理
	//bz := biz.New(c.Request.Context())
	//err := bz.RegisterUser(&param)
	//if err != nil {
	//	// 判断用户是否存在
	//	if err == errno.ErrorUserExistFail {
	//		response.ToErrorResponse(errno.ErrorUserExistFail)
	//		return
	//	}
	//	response.ToErrorResponse(errno.ErrorRegisterFail)
	//	return
	//}

	response.ToSuccessResponse(errno.Success)
	return
}
