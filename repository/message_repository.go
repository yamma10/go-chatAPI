package repository

import (
	"go-chat-api/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetAllMessages(messages *[]model.Message, roomId uint) error
	CreateMessage(message *model.Message,userId uint, roomId uint) error
	DeleteMessage(message *model.Message, userId uint, messageId uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &messageRepository{db}
}

func (mr *messageRepository) GetAllMessages(messages *[]model.Message, roomId uint) error {
	//Findで、引数と同じテーブルの中身を探す
	if err := mr.db.Where("room_id=?", roomId).Order("created_at").Find(messages).Error; err != nil {
		return err
	}

	return nil
}

func (mr *messageRepository) CreateMessage(message *model.Message, userId uint, roomId uint) error {
	//messageのポインタをdbに渡している
	if err := mr.db.Create(message).Error; err != nil {
		return err
	}

	return nil
}

func (mr *messageRepository) DeleteMessage(message *model.Message, userId uint, messageId uint) error {
	//messageのポインタをdbに渡している
	if err := mr.db.Where("user_id=?", userId).Delete(message, messageId).Error; err != nil {
		return err
	}

	return nil
}