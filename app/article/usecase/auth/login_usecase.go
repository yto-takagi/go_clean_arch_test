package auth

import (
	"encoding/json"
	"go_clean_arch_test/app/domain"
	auth "go_clean_arch_test/app/domain/auth"
	repository "go_clean_arch_test/app/domain/repository/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"go.uber.org/zap"
)

// LoginUsecase interface
type LoginUsecase interface {
	GetByEmail(email string) (domain.User, error)
	GetLoginUser(ctx *gin.Context) (domain.User, error)
}

type loginUsecase struct {
	loginUpRepository repository.LoginRepository
}

// NewSignUpUsecase constructor
func NewLoginUsecase(loginUpRepository repository.LoginRepository) LoginUsecase {
	return &loginUsecase{loginUpRepository: loginUpRepository}
}

// ログイン
func (loginUsecase *loginUsecase) GetByEmail(email string) (domain.User, error) {

	var user domain.User
	var login auth.Login
	login.Email = email
	userInfo, err := loginUsecase.loginUpRepository.GetByEmail(login.Email, user)
	if err != nil {
		return userInfo, err
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param email", email),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(userInfo)

	return userInfo, err
}

// ログインユーザー情報取得
func (loginUsecase *loginUsecase) GetLoginUser(ctx *gin.Context) (domain.User, error) {
	accessToken := ctx.Request.Header.Get("accessToken")
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○Request.Header○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(ctx.Request.Header)
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○accessToken○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(accessToken)

	testCookie, _ := ctx.Cookie("testCookie")
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○testCookie○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(testCookie)

	session := sessions.Default(ctx)
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○userInfo○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(session.Get(accessToken))
	// Json文字列がinterdace型で格納されている。dproxyのライブラリを使用して値を取り出す
	loginUserJson, err := dproxy.New(session.Get(accessToken)).String()

	var loginInfo domain.User
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
