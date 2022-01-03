package auth

import (
	"encoding/json"
	auth "go_clean_arch_test/app/domain/auth"
	"go_clean_arch_test/app/interfaces/database/repository/entity"
	"go_clean_arch_test/app/interfaces/delivery/request"
	repository "go_clean_arch_test/app/usecase/repository/auth"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

// LoginUsecase interface
type LoginUsecase interface {
	GetByEmail(email string) (request.User, error)
	GetLoginUser(ctx *gin.Context) (request.User, error)
}

type loginUsecase struct {
	loginUpRepository repository.LoginRepository
}

// NewSignUpUsecase constructor
func NewLoginUsecase(loginUpRepository repository.LoginRepository) LoginUsecase {
	return &loginUsecase{loginUpRepository: loginUpRepository}
}

// ログイン
func (loginUsecase *loginUsecase) GetByEmail(email string) (request.User, error) {

	var user entity.User
	var login auth.Login
	login.Email = email
	user, err := loginUsecase.loginUpRepository.GetByEmail(login.Email, user)

	var userInfo request.User
	userInfo.Id = user.Id
	userInfo.Email = user.Email
	userInfo.Password = user.Password

	if err != nil {
		return userInfo, err
	}

	return userInfo, err
}

// ログインユーザー情報取得
func (loginUsecase *loginUsecase) GetLoginUser(ctx *gin.Context) (request.User, error) {
	accessToken := ctx.Request.Header.Get("accessToken")

	session := sessions.Default(ctx)
	// Json文字列がinterdace型で格納されている。dproxyのライブラリを使用して値を取り出す
	loginUserJson, err := dproxy.New(session.Get(accessToken)).String()

	var loginInfo request.User
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		ctx.Abort()
	} else {
		// Json文字列のアンマーシャル
		err := json.Unmarshal([]byte(loginUserJson), &loginInfo)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
	return loginInfo, nil

}
