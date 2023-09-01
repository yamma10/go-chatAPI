package repository

import (
	"go-chat-api/model"

	"gorm.io/gorm"
)

type IMessageRepository interface {
	GetAllMessages(messages *[]model.Message,userId, roomId uint) error
	CreateMessage(message *model.Message) error
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
	result := mr.db.
	Table("messages").
	Joins("JOIN talk_rooms ON talk_rooms.id = messages.room_id").
	Where("messages.room_id = ? AND (talk_rooms.user1 = ? OR talk_rooms.user2 = ?)", roomId, userId, userId).
	Find(&messages);

	if result.Error != nil {
    // エラーハンドリング
		return result.Error
	}

	
	return nil
}

func (mr *messageRepository) CreateMessage(message *model.Message) error {
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