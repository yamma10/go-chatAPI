package controller

import (
	"fmt"
	"go-chat-api/model"
	"go-chat-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITalkRoomController interface {
	GetAllRooms(c echo.Context) error
	GetRoomById(c echo.Context) error
	CreateRoom(c echo.Context) error
	UpdateRoom(c echo.Context) error
	DeleteRoom(c echo.Context) error
}

type talkroomController struct {
	tu usecase.ITalkRoomUsecase
}

func NewTalkRoomController(tu usecase.ITalkRoomUsecase) ITalkRoomController {
	return &talkroomController{tu}
}

func (tc *talkroomController) GetAllRooms(c echo.Context) error {
	// println(c)
	//tokenCookieは表示されるので、Cookieは取得できてる
	// tokenCookie, err := c.Cookie("token")
	// if err != nil {
	// 		// エラー処理
	// 		return c.JSON(http.StatusUnauthorized,"tokenがないよ")
	// }
	// fmt.Println(tokenCookie)

	//ユーザーから送られてくるJWTトークンからユーザーIDを取り出す
	//router.goで実装しているjwtのミドルウェアでデコードされたものがuserという名前をつけて自動的に格納してくれる
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "userがないよ")
	}
	//fmt.Println(user)
	//userの中のClaimsを取り出す
	claims := user.Claims.(jwt.MapClaims)

	//claimsの中のuser_idを取り出す
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Invalid user_id format")
	}

	userIdTypeInfo := fmt.Sprintf("userIdの型は、%T", userId)
	fmt.Println(userIdTypeInfo)
	//型アサーション(型推論を上書き)したのちに型変換している
	roomRes, err := tc.tu.GetAllRooms(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//fmt.Println(roomRes)
	return c.JSON(http.StatusOK, roomRes)
}

func (tc *talkroomController) GetRoomById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("roomId")
	//Atoi・・・stringからintへの変換
	roomId, _ := strconv.Atoi(id)

	roomRes, err := tc.tu.GetRoomById(uint(userId.(float64)), uint(roomId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roomRes)
}

func (tc *talkroomController) CreateRoom(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	userId  = uint(userId.(float64))
	//fmt.Println(userId)
	room := model.TalkRoom{}
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(room)
	
	fmt.Printf("room.User1 = %T\n",room.User1)
	fmt.Printf("userId = %T\n",userId)
	// fmt.Println(userId)
	//roomには、User1とUser2が入っている
	//ここは一旦
	//room.User1 = uint(userId.(float64))
	//userIdとUser1または2は等しくないとダメ
	if userId == room.User1 || userId == room.User2 {

	} else {

		return c.JSON(http.StatusUnauthorized, nil)
	}

	roomRes, err := tc.tu.CreateRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, roomRes)
}

func (tc *talkroomController) UpdateRoom(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	//リクエストパラメータからroomIdを取得
	id := c.Param("roomId")
	roomId, _ := strconv.Atoi(id)

	room := model.TalkRoom{}
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(room)

	roomRes, err := tc.tu.UpdateRoom(room, uint(userId.(float64)), uint(roomId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, roomRes)
}

func (tc *talkroomController) DeleteRoom(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("roomId")
	roomId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteRoom(uint(userId.(float64)), uint(roomId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent((http.StatusNoContent))
}
