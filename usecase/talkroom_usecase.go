package usecase

import (
	"go-chat-api/model"
	"go-chat-api/repository"
)

type ITalkRoomUsecase interface {
	//[]model.
	GetAllRooms(userId uint) (
		[]model.TalkRoomResponse, error)
	GetRoomById(userId uint, roomId uint) (
		model.TalkRoomResponse, error)
	CreateRoom(room model.TalkRoom) (model.TalkRoomResponse, error)
	UpdateRoom(room model.TalkRoom, userId uint, roomId uint) (model.TalkRoomResponse, error)
	DeleteRoom(userId uint, roomId uint) error 
}

type talkroomUsecase struct {
	tr repository.ITalkRoomRepository
}

func NewTalkRoomUsecase(tr repository.ITalkRoomRepository) ITalkRoomUsecase {
	return &talkroomUsecase{tr}
}

func (tu *talkroomUsecase) GetAllRooms(userId uint) ([]model.TalkRoomResponse, error) {
	//Talkroomの配列
	rooms := []model.TalkRoom{}

	if err := tu.tr.GetAllRooms(&rooms, userId); err != nil {
		return nil, err
	}

	resRooms := []model.TalkRoomResponse{}

	for _, v := range rooms {
		t := model.TalkRoomResponse {
			ID: v.ID,
			Name: v.Name,
			User1: v.User1,
			User2: v.User2,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resRooms = append(resRooms, t)
	}

	return resRooms, nil
}

func (tu *talkroomUsecase) GetRoomById(userId uint, roomId uint) (model.TalkRoomResponse, error) {
	room := model.TalkRoom{}
	if err := tu.tr.GetRoomById(&room, userId, roomId);
	err != nil {
		return model.TalkRoomResponse{}, err
	}

	resRoom := model.TalkRoomResponse {
		ID: room.ID,
		User1: room.User1,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}

	return resRoom, nil
}

func (tu *talkroomUsecase) CreateRoom(room model.TalkRoom) (model.TalkRoomResponse, error) {
	if err := tu.tr.CreateRoom(&room); err != nil {
		return model.TalkRoomResponse{}, err
	}

	resRoom := model.TalkRoomResponse {
		ID: room.ID,
		Name: room.Name,
		User1: room.User1,
		User2: room.User2,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}

	return resRoom, nil
}

func (tu *talkroomUsecase) UpdateRoom(room model.TalkRoom, userId uint, roomId uint) (model.TalkRoomResponse, error) {

	if err := tu.tr.UpdateRoom(&room, userId, roomId); err != nil {
		return model.TalkRoomResponse{}, err
	}

	resRoom := model.TalkRoomResponse {
		ID: room.ID,
		Name: room.Name,
		User1: room.User1,
		User2: room.User2,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
	return resRoom, nil
}

func (tu *talkroomUsecase) DeleteRoom(userId uint, roomId uint) error {
	if err := tu.tr.DeleteRoom(userId, roomId); err != nil {
		return err
	}

	return nil
}