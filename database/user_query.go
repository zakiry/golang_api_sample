package database

import (
	"github.com/jinzhu/gorm"
	"github.com/zakiry/golang_api_sample/entity"
)

type User interface {
	GetById(id int64) entity.User
}

func NewUser(db *gorm.DB) User {
	return &userImpl{db: db}
}

type userImpl struct {
	db *gorm.DB
}

func (impl *userImpl) GetById(id int64) entity.User {
	record := entity.UserRecord{}
	impl.db.
		Select("id, name, age").
		Table("users").
		Where("id = ?", id).
		Find(&record)
	return record.Build()
}
