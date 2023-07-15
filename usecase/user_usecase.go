package usecase

import (
	"go-chat-api/model"
	"go-chat-api/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	
	if err != nil {
		return model.UserResponse{}, err
	}

	//新たなUserオブジェクトを作成
	newUser := model.User{Name: user.Name, Password: string(hash)}

	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse {
		ID: newUser.ID,
		Name: newUser.Name,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {

	//Nameが登録されていないかを確認する
	storedUser := model.User{}
	if err := uu.ur.GetUserByName(&storedUser, user.Name); err != nil {
		return "", err
	}
	//同名のユーザーが存在する場合、渡ってきたパスワードをハッシュ化し、取得したパスワードと、storedUserのパスワードを比較する
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))

	if err != nil {
		return "", err
	}

	//JWTトークンを作成する
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	//指定されたキーで署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}