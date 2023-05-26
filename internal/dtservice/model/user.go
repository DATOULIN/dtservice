package model

type UserM struct {
	ID
	UserId           int64  `json:"user_id"`
	UserName         string `json:"username"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Password         string `json:"password"`
	Avatar           string `json:"avatar"` // 头像
	Sex              uint8  `json:"sex"`
	SelfIntroduction string `json:"self_introduction"` // 自我介绍
	State            uint8  `json:"state"`             // 用户状态
	CreatedOn               // 创建时间
	ModifiedOn              // 修改时间
	LastLoginOn      int64  `json:"last_login_on"` // 上次登录时间
	LastLoginIP      string `json:"last_login_ip"` //登录IP
	UserType         uint8  `json:"user_type"`     // 用户类型
}

func (u UserM) TableName() string {
	return "dt_user"
}
