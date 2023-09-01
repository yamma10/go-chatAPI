package repository

import (
	"go-chat-api/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetAllMessages(messages *[]model.Message,userId, roomId uint) error
	CreateMessage(message *model.Message,userId uint, roomId uint) error
	DeleteMessage(userId uint, messageId uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) IMessageRepository {
	return &messageRepository{db}
}

func (mr *messageRepository) GetAllMessages(messages *[]model.Message,userId uint, roomId uint) error {
	//Findで指定しているmessagesに格納される
	if err := mr.db.Where("id=? AND (user1=? OR user2=?)",roomId, userId, userId).Order("created_at").Find(messages).Error; err != nil {
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

func (mr *messageRepository) DeleteMessage(userId uint, messageId uint) error {
	//messageのポインタをdbに渡している
	//Deleteに指定した構造体を同じテーブルの要素を削除する
	if err := mr.db.Where("user_id=?", userId).Delete(&model.Message{}).Error; err != nil {
		return err
	}

	return nil
}