package model

import "github.com/jinzhu/gorm"

type User struct {
	*Model
	UserId           int64  `json:"user_id"`
	UserName         string `json:"username"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Password         string `json:"password"`
	Avatar           string `json:"avatar"` // 头像
	Sex              uint8  `json:"sex"`
	SelfIntroduction string `json:"self_introduction"` // 自我介绍
	LastLoginOn      int64  `json:"last_login_on"`     // 上次登录时间
	LastLoginIP      string `json:"last_login_ip"`     //登录IP
	UserType         uint8  `json:"user_type"`         // 用户类型
}

func (u User) TableName() string {
	return "dt_user"
}

func (u User) CheckUserEmailExist(db *gorm.DB) (int64, []*User) {
	var user []*User
	return db.Where("email = ? AND is_del = ?", u.Email, 0).Find(&user).RowsAffected, user
}

func (u User) CheckUserIdExist(db *gorm.DB) (int64, []*User, error) {
	var user []*User
	db = db.Where("user_id = ? AND is_del = ? ", u.UserId, 0).Find(&user)
	return db.RowsAffected, user, db.Error
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&u).Where("user_id = ? AND is_del = ?", u.UserId, 0).Update(values).Error
}

func (u User) Count(db *gorm.DB) (int, error) {
	var count int
	var err error
	if u.Email != "" {
		db = db.Where("email = ?", u.Email)
	}
	db = db.Where("state = ?", u.State)
	if err = db.Model(&u).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (u User) List(db *gorm.DB, pageOffset, pageSize int) ([]*User, error) {
	var users []*User
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if u.Email != "" {
		db = db.Where("email = ?", u.Email)
	}
	db = db.Where("state = ?", u.State)
	if err = db.Where("is_del = ?", 0).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u User) Delete(db *gorm.DB) error {
	return db.Where("user_id = ? AND is_del = ?", u.UserId, 0).Delete(&u).Error
}
