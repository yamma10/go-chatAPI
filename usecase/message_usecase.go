package usecase

import "go-chat-api/model"

type IMessageUsecase interface {
	GetAllMessages(roomId uint) ([]model.MessageResponse, error)

}