package controller

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/service"
	"github.com/DATOULIN/dtservice/internal/pkg/app"
	"github.com/DATOULIN/dtservice/internal/pkg/convert"
	"github.com/DATOULIN/dtservice/internal/pkg/errno"
	v1 "github.com/DATOULIN/dtservice/pkg/api/dtservice/v1"
	"github.com/DATOULIN/dtservice/pkg/util"
	"github.com/DATOULIN/dtservice/pkg/validate"
	"github.com/gin-gonic/gin"
	"math/rand"
	"path"
	"strconv"
	"time"
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
	svc := service.New(c.Request.Context())
	err := svc.RegisterUser(&param)
	if err != nil {
		// 判断用户是否存在
		if err == errno.ErrorUserExistFail {
			response.ToErrorResponse(errno.ErrorUserExistFail)
			return
		}
		response.ToErrorResponse(errno.ErrorRegisterFail)
		return
	}

	response.ToSuccessResponse(errno.Success)
	return
}

func (u User) Login(c *gin.Context) {
	param := v1.LoginRequest{}
	response := app.NewResponse(c)
	valid, errs := validate.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	ip := util.GetRequestIP(c)
	result, err := svc.Login(&param, ip)
	if err != nil {
		// 判断用户是否存在
		if err == errno.ErrorUserNoExistFail {
			response.ToErrorResponse(errno.ErrorUserNoExistFail)
			return
		}
		// 判断密码是否正确
		if err == errno.ErrorPasswordFail {
			response.ToErrorResponse(errno.ErrorPasswordFail)
			return
		}
		// 判断密码是否被禁用
		if err == errno.ErrorUserBanFail {
			response.ToErrorResponse(errno.ErrorUserBanFail)
			return
		}
		response.ToErrorResponse(errno.ErrorLoginFail)
		return
	}
	response.ToResponseData(result, errno.Success)
	return
}

func (u User) Logout(c *gin.Context) {
	response := app.NewResponse(c)

	// 业务处理
	token := c.Request.Header.Get("Authorization")
	svc := service.New(c.Request.Context())
	err := svc.Logout(token)
	if err != nil {
		response.ToErrorResponse(errno.ErrorLogoutFail)
		return
	}
	response.ToSuccessResponse(errno.Success)
	return
}

func (u User) GetUserList(c *gin.Context) {
	param := v1.UserListRequest{}
	response := app.NewResponse(c)
	valid, errs := validate.BindQueryAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountUser(&v1.UserListRequest{Email: param.Email, State: param.State})
	if err != nil {
		response.ToErrorResponse(errno.ErrorCountUserFail)
		return
	}
	users, err := svc.GetUserList(&param, &pager)
	if err != nil {
		response.ToErrorResponse(errno.ErrorGetUserListFail)
		return
	}
	response.ToResponseList(users, totalRows, errno.Success)
}

func (u User) UpdateUser(c *gin.Context) {
	param := v1.UpdateUserRequest{UserId: convert.StrTo(c.Param("user_id")).MustInt64()}
	id := v1.CheckUserExistRequest{UserId: convert.StrTo(c.Param("user_id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := validate.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	isExist, err := svc.CheckUserExist(&id) // 检验用户是否存在
	if err != nil {
		response.ToErrorResponse(errno.ErrorUserExist)
		return
	}
	if isExist == 0 {
		response.ToErrorResponse(errno.ErrorUserNoExistFail)
		return
	}
	err = svc.UpdateUser(&param) // 更新标签
	if err != nil {
		response.ToErrorResponse(errno.ErrorUpdateUserFail)
		return
	}

	response.ToSuccessResponse(errno.Success)
	return
}

func (u User) ResetPassword(c *gin.Context) {
	param := v1.ResetPasswordRequest{UserId: convert.StrTo(c.Param("user_id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := validate.BindAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 业务处理
	svc := service.New(c.Request.Context())
	err := svc.ResetPassword(&param)
	if err != nil {
		// 判断用户是否存在
		if err == errno.ErrorUserNoExistFail {
			response.ToErrorResponse(errno.ErrorUserNoExistFail)
			return
		}
		// 密码不正确
		if err == errno.ErrorOldPasswordFail {
			response.ToErrorResponse(errno.ErrorOldPasswordFail)
			return
		}
		// 新密码不能与旧密码一样
		if err == errno.ErrorOldNowPswFail {
			response.ToErrorResponse(errno.ErrorOldNowPswFail)
			return
		}
		response.ToErrorResponse(errno.ErrorResetPasswordFail)
		return
	}
	response.ToSuccessResponse(errno.Success)
	return
}

func (u User) UploadAvatar(c *gin.Context) {
	response := app.NewResponse(c)
	file, err := c.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errno.ErrorFileIsRequiredFail)
		return
	}
	param := v1.UploadAvatarRequest{
		File:   file,
		UserId: convert.StrTo(c.Param("user_id")).MustInt64(),
	}
	// 业务处理
	svc := service.New(c.Request.Context())
	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000) + path.Ext(file.Filename)
	users, errs := svc.UploadAvatar(&param, newFileName)
	if errs != nil {
		// 判断用户是否存在
		if errs == errno.ErrorUserNoExistFail {
			response.ToErrorResponse(errno.ErrorUserNoExistFail)
			return
		}
		response.ToErrorResponse(errno.ErrorUploadFail)
		return
	}
	// 保存到服务器
	uperr := c.SaveUploadedFile(file, app.BuildSavePath(users[0].UserId, newFileName))
	if uperr != nil {
		response.ToErrorResponse(errno.ErrorUploadFail)
		return
	}
	response.ToSuccessResponse(errno.UploadSuccess)
}

func (u User) Delete(c *gin.Context) {
	param := v1.DeleteUserRequest{UserId: convert.StrTo(c.Param("user_id")).MustInt64()}
	id := v1.CheckUserExistRequest{UserId: convert.StrTo(c.Param("user_id")).MustInt64()}
	response := app.NewResponse(c)
	valid, errs := validate.BindQueryAndValid(c, &param)
	if !valid {
		response.ToErrorResponse(errno.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	isExist, err := svc.CheckUserExist(&id)
	if isExist == 0 {
		response.ToErrorResponse(errno.ErrorUserNoExistFail)
		return
	}
	if err != nil {
		response.ToErrorResponse(errno.ErrorDeleteUserFail)
		return
	}
	err = svc.DeleteUser(&param)
	if err != nil {
		response.ToErrorResponse(errno.ErrorDeleteUserFail)
		return
	}

	response.ToSuccessResponse(errno.Success)
	return
}
