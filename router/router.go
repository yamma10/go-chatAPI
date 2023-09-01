package router

import (
	"go-chat-api/controller"
	"os"
	//"net/http"
	//"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITalkRoomController) *echo.Echo {
	e := echo.New()

	//CORS
	

	//CSRFのミドルウェア
	/* e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		//CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60
	})) */


	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/rooms")
	//tに対してjwtのミドルウェアを適用
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	
	t.GET("", tc.GetAllRooms)
	t.GET("/:roomId", tc.GetRoomById)
	t.POST("", tc.CreateRoom)
	t.PUT("/:roomId", tc.UpdateRoom)
	t.DELETE("/:roomId", tc.DeleteRoom)
	return e
	
}