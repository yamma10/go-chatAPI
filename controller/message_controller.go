package controller

import (
	"go-chat-api/model"
	"go-chat-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IMessageController interface {
	GetAllMessages(c echo.Context) error
	CreateMessage(c echo.Context) error
	DeleteMessage(c echo.Context) error
}

type messageController struct {
	mu usecase.IMessageUsecase
}

func NewMessageController(mu usecase.IMessageUsecase) IMessageController {
	return &messageController{mu}
}

func (mc *messageController) GetAllMessages(c echo.Context) error {
	//ユーザーから送られてくるJWTトークンからユーザーIDを取り出す
	//router.goで実装しているjwtのミドルウェアでデコードされたものがuserという名前をつけて自動的に格納してくれる
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "userがないよ")
	}
	claims := user.Claims.(jwt.MapClaims)

	//claimsの中のuser_idを取り出す
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "Invalid user_id format")
	}

	id := c.Param("roomId")
	//Atoi・・・stringからintへの変換
	roomId, _ := strconv.Atoi(id)
	//型アサーション(型推論を上書き)したのちに型変換している
	messageRes, err := mc.mu.GetAllMessages(uint(userId), uint(roomId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}



	return c.JSON(http.StatusOK, messageRes)
}

func (mc *messageController) CreateMessage(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("roomId")
	//Atoi・・・stringからintへの変換
	roomId, _ := strconv.Atoi(id)

	message:= model.Message{}
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	messageRes, err := mc.mu.CreateMessage(uint(userId), uint(roomId))

	//internalServerError・・・サーバー内部のエラー
	//トークルームに所属してない場合はunauthorizedを返すようにする(後で実装)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, messageRes)
}

func (mc *messageController) DeleteMessage(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("messageId")
	//Atoi・・・stringからintへの変換
	messageId, _ := strconv.Atoi(id)

	if err := mc.mu.DeleteMessage(uint(userId), uint(messageId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "メッセージの削除に成功しました")
}

