package repository

import (
	"go-chat-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByName(user *model.User, name string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

//コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByName(user *model.User, name string) error {
	if err := ur.db.Where("name=?", name).First(user).Error; err!= nil {
		return err
	}

	return nil
}

//userモデルの
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}