package delivery

import (
	"go_clean_arch_test/app/article/usecase"
	loginUsecase "go_clean_arch_test/app/article/usecase/auth"
	"go_clean_arch_test/app/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const NO_AUTHORS = "no authors"
const EXISTS = "exists"

// AuthorHandler interface
type AuthorHandler interface {
	GetAllAuthor(ctx *gin.Context)
	InputAuthor(ctx *gin.Context)
	UpdateAuthor(ctx *gin.Context)
	DeleteAuthor(ctx *gin.Context)
}

type authorHandler struct {
	articleusecase usecase.ArticleUsecase
	authorUsecase  usecase.AuthorUsecase
	loginUsecase   loginUsecase.LoginUsecase
}

// NewAuthorHandler constructor
func NewAuthorHandler(articleusecase usecase.ArticleUsecase, authorUsecase usecase.AuthorUsecase, loginUsecase loginUsecase.LoginUsecase) AuthorHandler {
	return &authorHandler{articleusecase: articleusecase, authorUsecase: authorUsecase, loginUsecase: loginUsecase}
}

// 全件取得
func (authorHandler *authorHandler) GetAllAuthor(ctx *gin.Context) {

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○User.ID○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(user.Id)

	authors, err := authorHandler.authorUsecase.GetByUser(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if len(authors) < 0 || authors == nil {
		ctx.JSON(http.StatusOK, NewH(NO_AUTHORS, authors))
		return
	}
	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), authors))
}

// 新規登録
func (authorHandler *authorHandler) InputAuthor(ctx *gin.Context) {
	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	author := domain.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	author.UserId = user.Id

	// カテゴリー検索(カテゴリー名で)
	authorByName, err := authorHandler.authorUsecase.GetByName(author.Name, author.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if authorByName.Id != 0 {
		ctx.JSON(http.StatusOK, NewH(EXISTS, author))
		return

	}

	// 新規登録
	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.String("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	authorByName, err = authorHandler.authorUsecase.Input(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■authorByName.Id■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(authorByName.Id)

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), author))
}

// 更新
func (authorHandler *authorHandler) UpdateAuthor(ctx *gin.Context) {

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	author := domain.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(302, NewH("Bad Request", author))
		return
	}

	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	author.UserId = user.Id

	authorByName, err := authorHandler.authorUsecase.GetByName(author.Name, author.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if authorByName.Id != 0 {
		ctx.JSON(http.StatusOK, NewH(EXISTS, author))
		return
	}

	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.String("■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在する場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	err = authorHandler.authorUsecase.Update(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), author))
}

// 削除(id指定)
func (authorHandler *authorHandler) DeleteAuthor(ctx *gin.Context) {
	author := domain.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	// ログインユーザーID且つ、記事ID データが存在しなければ、302で返す
	user, err := authorHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	articleByIdAndUserId, err := authorHandler.authorUsecase.GetByAuthorIdAndUserId(author.Id, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	if articleByIdAndUserId.Id == 0 {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), author))
		return
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Int("param id", author.Id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	err = authorHandler.authorUsecase.Delete(&author, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	// エラーじゃなければ(削除件数1以上)、紐づくarticle削除
	err = authorHandler.articleusecase.DeleteByAuthor(author.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), author.Id))
}
