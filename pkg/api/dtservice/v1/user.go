package v1

import "mime/multipart"

type RegisterRequest struct {
	UserName        string `form:"username" binding:"required,min=2,max=100"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `binding:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type UserListRequest struct {
	UserName string `form:"user_name"`
	Email    string `form:"email"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	UserId   int64  `form:"user_id" binding:"required"`
	UserName string `form:"user_name" `
	State    uint8  `form:"state" binding:"oneof=0 1"`
}

type DeleteUserRequest struct {
	UserId int64 `form:"user_id" binding:"required"`
}

type CheckUserExistRequest struct {
	UserId int64 `form:"user_id" binding:"required"`
}

type ResetPasswordRequest struct {
	UserId          int64  `form:"user_id" binding:"required"`
	OldPassword     string `binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `binding:"required,eqfield=Password"`
}

type UploadAvatarRequest struct {
	UserId int64                 `form:"user_id" binding:"required"`
	File   *multipart.FileHeader `form:"avatar"`
}

type UserInfo struct {
	UserId int64 `json:"user_id"`
}
type LoginResultResponse struct {
	AccessToken string `json:"access_token"`
	UserInfo
}
