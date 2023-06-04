package service

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/model"
	"github.com/DATOULIN/dtservice/internal/pkg/app"
	"github.com/DATOULIN/dtservice/internal/pkg/auth"
	"github.com/DATOULIN/dtservice/internal/pkg/errno"
	"github.com/DATOULIN/dtservice/internal/pkg/setting"
	v1 "github.com/DATOULIN/dtservice/pkg/api/dtservice/v1"
	"github.com/DATOULIN/dtservice/pkg/util"
	"time"
)

func (svc *Service) RegisterUser(params *v1.RegisterRequest) error {
	isExist, _ := svc.dao.CheckUserEmailExist(params.Email)
	if isExist != 0 {
		return errno.ErrorUserExistFail
	}
	return svc.dao.RegisterUser(params.UserName, params.Email, params.Password)

}

func (svc *Service) Login(params *v1.LoginRequest, ip string) (interface{}, error) {
	isExist, users := svc.dao.CheckUserEmailExist(params.Email)
	// 用户不存在
	if isExist == 0 {
		return "", errno.ErrorUserNoExistFail
	}
	// 密码不正确
	if !util.EqualsPassword(params.Password, users[0].Password) {
		return "", errno.ErrorPasswordFail
	}
	// 判断密码是否被禁用
	if users[0].State == 0 {
		return "", errno.ErrorUserBanFail
	}
	// 登录成功后将token返回
	token := auth.GenerateToken(params.Email, users[0].UserId)

	result := v1.LoginResultResponse{
		AccessToken: token,
		UserInfo: v1.UserInfo{
			UserId: users[0].UserId,
		},
	}
	err := svc.dao.Login(users[0].UserId, ip, time.Now().Unix())
	if err != nil {
		return "nil", err
	}
	return result, nil
}

func (svc *Service) Logout(token string) error {
	return auth.JoinBlackList(token)
}

func (svc *Service) CountUser(params *v1.UserListRequest) (int, error) {
	return svc.dao.CountUser(params.Email, params.State)
}

func (svc *Service) GetUserList(params *v1.UserListRequest, pager *app.Pager) ([]*model.User, error) {
	users, err := svc.dao.GetUserList(params.Email, params.State, pager.Page, pager.PageSize)

	return users, err
}

// CheckUserExist 根据userId查找是否存在用户
func (svc *Service) CheckUserExist(params *v1.CheckUserExistRequest) (int64, error) {
	isExist, _, err := svc.dao.CheckUserIdExist(params.UserId)
	return isExist, err
}

// UpdateUser 更新用户
func (svc *Service) UpdateUser(params *v1.UpdateUserRequest) error {
	return svc.dao.UpdateUser(params.UserId, params.UserName, params.State)
}

// ResetPassword 重置密码
func (svc *Service) ResetPassword(params *v1.ResetPasswordRequest) error {
	isExist, users, err := svc.dao.CheckUserIdExist(params.UserId)
	// 用户不存在
	if isExist == 0 {
		return errno.ErrorUserNoExistFail
	}
	// 密码不正确
	if !util.EqualsPassword(params.OldPassword, users[0].Password) {
		return errno.ErrorOldPasswordFail
	}
	// 新密码不能与旧密码一样
	if util.EqualsPassword(params.Password, users[0].Password) {
		return errno.ErrorOldNowPswFail
	}
	err = svc.dao.ResetPassword(params.UserId, params.Password)
	return err
}

// UploadAvatar 上传头像
func (svc *Service) UploadAvatar(params *v1.UploadAvatarRequest, newFileName string) ([]*model.User, error) {
	isExist, users, err := svc.dao.CheckUserIdExist(params.UserId)
	// 用户不存在
	if isExist == 0 {
		return users, errno.ErrorUserNoExistFail
	}
	ipAddrList, _ := util.GetLocalIPv4s()
	// 本机port
	port := setting.ServerSettings.HttpPort

	uploadPath := ipAddrList[3] + ":" + port + app.BuildSavePath(users[0].UserId, newFileName)
	err = svc.dao.UploadAvatar(params.UserId, uploadPath)
	return users, err
}

func (svc *Service) DeleteUser(param *v1.DeleteUserRequest) error {
	return svc.dao.DeleteUser(param.UserId)
}
