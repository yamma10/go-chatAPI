package router

import (
	"go-chat-api/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()

	//CORS
	

	//CSRFのミドルウェア
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		//CookieSameSite: http.SameSiteNoneMode,
		//CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge: 60
	}))


	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/scrf", uc.CsrfToken)

	//t := e.Group("/tasks")
	//tに対してjwtのミドルウェアを適用
	
	return e
}