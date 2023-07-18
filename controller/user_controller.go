package controller

import (
	"fmt"
	"go-chat-api/model"
	"go-chat-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}

	//リクエストボディを構造体にバインドする
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(user.Name)
	//userUsecaseのSignUpを呼び出す
	userRes, err := uc.uu.SignUp(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uc.uu.Login(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//httpパッケージに定義されているCookie構造体を作成する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	//user_usecaseのloginで作成したJWTトークンを設定する
	cookie.Value = tokenString
	//有効期限を24時間に設定する
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//Postmanでテストするため一旦コメントアウト
	//cookie.Secure = true
	
	//クライアントサイドのJSからいじれないようにする
	cookie.HttpOnly = true
	//クロスドメイン間でのcookieの送受信を許可
	cookie.SameSite = http.SameSiteNoneMode

	fmt.Println(cookie)
	//cookieをHTTPResponseに加えるように設定する
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//Postmanでテストするため、一旦コメント
	//cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	fmt.Println(cookie)
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}