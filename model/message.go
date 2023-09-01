package model

import "time"

//構造体を直接JSONにすることはできない
//外部キー制約をすることで、TalkRoomやSenderが削除されるとMessageも削除される
type Message struct {
	ID uint `json:"id" gorm:"primaryKey"`
	RoomID     uint `json:"room_id" gorm:"foreignKey:Room;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SenderID   string     `json:"sender_id" gorm:"foreignKey:Sender;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessageResponse struct {
	ID int32 `json:"id" gorm:"primaryKey"`
	SenderId string `json:"sender_id" gorm:"foreignKey:Sender;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}