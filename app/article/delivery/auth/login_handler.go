package delivery

import (
	"encoding/json"
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/delivery"
	usecase "go_clean_arch_test/app/article/usecase/auth"
	domain "go_clean_arch_test/app/domain/auth"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Usecase usecase.LoginUsecase
}

func NewLoginHandler(db database.DB) *LoginHandler {
	return &LoginHandler{
		Usecase: usecase.LoginUsecase{
			DB: &database.DBRepository{DB: db},
		},
	}
}

func (handler *LoginHandler) Login(ctx *gin.Context) {
	var request domain.Login
	err := ctx.BindJSON(&request)
	// request := domain.Login{}
	// err := ctx.Bind(&request)
	if err != nil {
		// ctx.Status(http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, delivery.NewH("error", http.StatusBadRequest))
	} else {
		// メールアドレスでユーザ情報取得
		user := handler.Usecase.GetByEmail(request.Email)

		// パスワードハッシュ化対応 新規登録時使用
		// hashed, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
		// log.Println(string(hashed))

		// ハッシュ値でのパスワード比較
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			// ctx.Status(http.StatusBadRequest)
			ctx.JSON(http.StatusBadRequest, delivery.NewH("error", http.StatusBadRequest))
		} else {
			session := sessions.Default(ctx)
			// セッションに格納する為にユーザ情報をJson化
			loginUser, err := json.Marshal(user)
			if err == nil {
				u, _ := uuid.NewRandom()
				accessToken := u.String()
				session.Set(accessToken, string(loginUser))
				session.Save()
				log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○Login Request.Header○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
				log.Println(ctx.Request.Header)
				log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○Login accessToken○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
				log.Println(accessToken)
				log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○Login userInfo○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
				log.Println(session.Get(accessToken))
				ctx.SetSameSite(http.SameSiteDefaultMode)
				ctx.SetCookie("testCookie", "testCookie", 3600, "/", "localhost", false, false)

				ctx.JSON(http.StatusOK, delivery.NewH("success", accessToken))
			} else {
				// ctx.Status(http.StatusInternalServerError)
				ctx.JSON(http.StatusInternalServerError, delivery.NewH("error", http.StatusInternalServerError))
			}
		}
	}
}
