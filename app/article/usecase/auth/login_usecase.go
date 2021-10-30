package auth

import (
	"encoding/json"
	sql "go_clean_arch_test/app/article/repository/sql/auth"
	"go_clean_arch_test/app/article/usecase"
	"go_clean_arch_test/app/domain"
	auth "go_clean_arch_test/app/domain/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"go.uber.org/zap"
)

type LoginUsecase struct {
	DB usecase.DBRepository
}

// ログイン
func (usecase *LoginUsecase) GetByEmail(email string) domain.User {
	db := usecase.DB.Connect()
	// defer db.Close()

	var user domain.User
	var login auth.Login
	login.Email = email
	userInfo := sql.GetByEmail(db, login.Email, user)

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_usecase.go ++++++++++++++++++++++",
		zap.String("method", "GetByName"),
		zap.String("param email", email),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	log.Println(userInfo)

	return userInfo
}

func (usecase *LoginUsecase) GetLoginUser(ctx *gin.Context) domain.User {
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
	return loginInfo

}
