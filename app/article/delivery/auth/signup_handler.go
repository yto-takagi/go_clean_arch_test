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
)

type SignUpHandler struct {
	Usecase usecase.SignUpUsecase
}

func NewSignUpHandler(db database.DB) *SignUpHandler {
	return &SignUpHandler{
		Usecase: usecase.SignUpUsecase{
			DB: &database.DBRepository{DB: db},
		},
	}
}

func (handler *SignUpHandler) SignUp(ctx *gin.Context) {
	var request domain.SignUp
	err := ctx.BindJSON(&request)
	// request := domain.Login{}
	// err := ctx.Bind(&request)
	if err != nil {
		// ctx.Status(http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, delivery.NewH("error", http.StatusBadRequest))
	} else {
		// 会員登録処理
		user, err := handler.Usecase.SignUp(request.Email, request.Password)
		log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■会員登録 user■■■■■■■■■■■■■■■■■■■■■■■■■■")
		log.Println(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, delivery.NewH(err.Error(), http.StatusInternalServerError))
		} else {
			session := sessions.Default(ctx)
			// セッションに格納する為にユーザ情報をJson化
			loginUser, err := json.Marshal(user)
			if err == nil {
				u, _ := uuid.NewRandom()
				accessToken := u.String()
				session.Set(accessToken, string(loginUser))
				session.Save()

				ctx.JSON(http.StatusOK, delivery.NewH("success", accessToken))
			} else {
				ctx.JSON(http.StatusInternalServerError, delivery.NewH("error", http.StatusInternalServerError))
			}
		}
	}
}
