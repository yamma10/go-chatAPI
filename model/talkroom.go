package model

import "time"

type TalkRoom struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	User1     uint     `json:"user_1" gorm:"foreignKey:User"`
	User2     uint      `json:"user_2" gorm:"foreignKey:User"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TalkRoomResponse struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	User1     uint      `json:"user_1" gorm:"foreignKey:User1ID"`
	User2     uint      `json:"user_2" gorm:"foreignKey:User2ID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}