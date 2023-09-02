package repository

import (
	"fmt"
	"go-chat-api/model"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
type ITalkRoomRepository interface {
	GetAllRooms(rooms *[]model.TalkRoom, userId uint) error
	GetRoomById(room *model.TalkRoom, userId uint, roomId uint) error
	CreateRoom(room *model.TalkRoom) error
	UpdateRoom(room *model.TalkRoom, userId uint, roomId uint) error
	DeleteRoom( userId uint ,roomId uint) error
}

type talkroomRepository struct {
	db *gorm.DB
}

func NewTalkRoomRepository(db *gorm.DB) ITalkRoomRepository {
	return &talkroomRepository{db}
}

func (tr *talkroomRepository) GetAllRooms(rooms *[]model.TalkRoom, userId uint) error {
	//roomsという配列に格納されるので、
	//rooms[0]などの形で取り出せる
	//User1またはUser2がuserIdと等しいもの
	//findは、modelの中身と同じテーブルを探すという意味
	if err := tr.db.Where("user1 = ? OR user2 = ?", userId, userId).Order("created_at").Find(rooms).Error; err != nil {
		return err
	}

	return nil
}

func (tr *talkroomRepository) GetRoomById(room *model.TalkRoom, userId uint, roomId uint) error {
	//Firstの部分はroomの主キーがroomIdと一致するものをとってくる
	//とってきた情報はGetRoomByIdの引数として渡されたTalkRoom構造体の参照に渡される
	result := tr.db.Where("id = ? AND (user1 = ? OR user2 = ?)",roomId, userId, userId).Find(room)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("トークルームが存在しないか、権限がありません")
	}

	

	return nil
}

func (tr *talkroomRepository) CreateRoom(room *model.TalkRoom) error {
	//roomのポインタをdbに渡している
	if err := tr.db.Create(room).Error; err != nil {
		return err
	}

	return nil
}

func (tr *talkroomRepository) UpdateRoom(room *model.TalkRoom, userId uint, roomId uint) error {
	// 更新したい新しい値を指定
	newData := map[string]interface{}{
		"updated_at": time.Now(),
		"name": room.Name,
	}
	//Clauses(clause.Returning{})をしていすると更新後のレコードをModelに指定したroomオブジェクトに書き込んでくれる
	result := tr.db.Model(room).Clauses(clause.Returning{}).Where("id=? AND (user1=? OR user2=?)", roomId,userId,userId).Updates(newData)
	fmt.Println()
	if result.Error != nil {
		return result.Error
	}
	//更新されたレコードの数を取得する
	if result.RowsAffected < 1 {
		return fmt.Errorf("トークルームが存在しないか、権限がありません")
	}

	return nil
}

func (tr *talkroomRepository) DeleteRoom(userId uint, roomId uint) error {
	result := tr.db.Where("id=? AND (user1=? OR user2=?)",roomId, userId, userId).Delete(&model.TalkRoom{})
	if result.Error != nil {
		return result.Error 
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}