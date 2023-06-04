package dao

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/model"
	"github.com/DATOULIN/dtservice/internal/pkg/app"
	"github.com/DATOULIN/dtservice/pkg/id"
	"github.com/DATOULIN/dtservice/pkg/util"
)

func (d *Dao) RegisterUser(username string, email string, password string) error {
	// 生成userid
	userid := id.GenId()
	// 密码加密
	pwd, _ := util.EncryptPassword(password)
	user := model.User{
		UserId:   userid,
		UserName: username,
		Email:    email,
		Password: pwd,
		Model: &model.Model{
			CreatedBy: email,
			State:     1,
		},
	}
	return user.Create(d.engine)
}

func (d *Dao) CheckUserEmailExist(email string) (int64, []*model.User) {
	user := model.User{
		Email: email,
	}
	return user.CheckUserEmailExist(d.engine)
}

func (d *Dao) CheckUserIdExist(userId int64) (int64, []*model.User, error) {
	user := model.User{UserId: userId}
	rowsAffected, users, err := user.CheckUserIdExist(d.engine)
	return rowsAffected, users, err
}

func (d *Dao) UpdateUser(userId int64, username string, state uint8) error {
	user := model.User{
		UserId: userId,
	}
	values := map[string]interface{}{
		"state": state,
	}
	if username != "" {
		values["user_name"] = username
	}

	return user.Update(d.engine, values)
}

func (d *Dao) CountUser(email string, state uint8) (int, error) {
	tag := model.User{Email: email, Model: &model.Model{State: state}}
	return tag.Count(d.engine)
}

func (d *Dao) GetUserList(email string, state uint8, page, pageSize int) ([]*model.User, error) {
	user := model.User{Email: email, Model: &model.Model{State: state}}
	pageOffset := app.GetPageOffset(page, pageSize)
	return user.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) ResetPassword(userId int64, password string) error {
	pwd, _ := util.EncryptPassword(password)
	user := model.User{
		UserId: userId,
	}
	// 密码加密
	values := map[string]interface{}{
		"password": pwd,
	}
	return user.Update(d.engine, values)
}

func (d *Dao) Login(userId int64, ip string, lastLoginOn int64) error {
	user := model.User{
		UserId: userId,
	}
	values := map[string]interface{}{
		"last_login_ip": ip,
		"last_login_on": lastLoginOn,
	}

	return user.Update(d.engine, values)
}

func (d *Dao) UploadAvatar(userId int64, avatar string) error {
	user := model.User{
		UserId: userId,
	}
	values := map[string]interface{}{
		"avatar": avatar,
	}
	return user.Update(d.engine, values)
}

func (d *Dao) DeleteUser(userId int64) error {
	user := model.User{
		UserId: userId,
	}
	return user.Delete(d.engine)
}
