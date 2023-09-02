package usecase

import (
	"go-chat-api/model"
	"go-chat-api/repository"
)

type IMessageUsecase interface {
	GetAllMessages(userId, roomId uint) ([]model.MessageResponse, error)
	CreateMessage(userId uint, roomId uint) (model.MessageResponse, error)
	DeleteMessage(userId uint, messageId uint) error
}

type messageUsecase struct {
	mr repository.IMessageRepository
}

func NewMessageUsecase(mr repository.IMessageRepository) IMessageUsecase {
	return &messageUsecase{mr}
}

func (mu *messageUsecase) GetAllMessages(userId, roomId uint) ([]model.MessageResponse, error) {
	//messageの配列(スライス)
	messages := []model.Message{}

	//messagesの参照を渡している
	if err := mu.mr.GetAllMessages(&messages, userId, roomId); err != nil {
		return nil, err
	}

	resMessages := []model.MessageResponse{}

	for _, v := range resMessages {
		t := model.MessageResponse{
			ID: v.ID,
			SenderID: v.SenderID,
			Content: v.Content,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resMessages = append(resMessages, t)
	}

	return resMessages, nil
}

func (mu *messageUsecase) CreateMessage(userId uint, roomId uint) (model.MessageResponse, error) {
	message := model.Message{}

	//messageの参照を渡している
	if err := mu.mr.CreateMessage(&message, userId, roomId); err != nil {
		return model.MessageResponse{}, err
	}

	resMessage := model.MessageResponse{
		ID: message.ID,
		SenderID: message.SenderID,
		Content: message.Content,
		CreatedAt: message.CreatedAt,
	}

	return resMessage, nil
}

func (mu *messageUsecase) DeleteMessage(userId uint, messageId uint) error {


	//messageの参照を渡している
	if err := mu.mr.DeleteMessage(userId, messageId); err != nil {
		return err
	}

	return nil
}