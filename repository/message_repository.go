package repository

import (
	"errors"
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

	result := mr.db.Where("id = ? AND(user1 = ? OR user2 = ?)",message.RoomID, message.SenderID, message.SenderID).First(&model.TalkRoom{})

	if result.Error != nil {
		// エラーハンドリング
		return errors.New("トークルームに所属していません")
	}
	
	//Senderがトークルームに所属している場合のみメッセージを送信できる
	if result.RowsAffected > 0 {
		//messageのポインタをdbに渡している
		if err := mr.db.Create(message).Error; err != nil {
			return err
		}
	} 
	

	return nil
}

func (mr *messageRepository) DeleteMessage(userId uint, messageId uint) error {

	result := mr.db.Where("id = ? AND sender_id=? ", messageId,userId).First(&model.Message{})

	if result.RowsAffected == 0 {
		// エラーハンドリング
		return errors.New("メッセージが存在しないか、削除権限がありません")
	}

	//messageのポインタをdbに渡している
	//Deleteに指定した構造体を同じテーブルの要素を削除する
	if err := mr.db.Where("id = ? AND sender_id=? ", messageId,userId).Delete(&model.Message{}).Error; err != nil {
		return err
	}

	return nil
}