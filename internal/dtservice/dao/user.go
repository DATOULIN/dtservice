package dao

import (
	"github.com/DATOULIN/dtservice/internal/dtservice/model"
	"github.com/jinzhu/gorm"
)

func (d *Dao) Create(db *gorm.DB) error {
	u := model.UserM{}
	return db.Create(&u).Error
}

func (d *Dao) Update(db *gorm.DB, values interface{}) error {
	u := model.UserM{}
	return db.Model(&u).Where("user_id = ?", u.UserId).Update(values).Error
}
