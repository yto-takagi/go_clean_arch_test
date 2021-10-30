package delivery

import (
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/usecase"
	loginUsecase "go_clean_arch_test/app/article/usecase/auth"
	"go_clean_arch_test/app/domain"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// var LoginInfo auth.SessionInfo

type AuthorHandler struct {
	ArticleUsecase usecase.ArticleUsecase
	Usecase        usecase.AuthorUsecase
	LoginUsecase   loginUsecase.LoginUsecase
}

func NewAuthorHandler(db database.DB) *AuthorHandler {
	return &AuthorHandler{
		ArticleUsecase: usecase.ArticleUsecase{
			DB: &database.DBRepository{DB: db},
		},
		Usecase: usecase.AuthorUsecase{
			DB: &database.DBRepository{DB: db},
		},
	}
}

// 全件取得
func (handler *AuthorHandler) GetAllAuthor(ctx *gin.Context) {

	user := handler.LoginUsecase.GetLoginUser(ctx)
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○User.ID○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(user.Id)

	authors := handler.Usecase.GetByUser(user.Id)
	// ctx.JSON(res.StatusCode, NewH("success", articles))
	if len(authors) < 0 || authors == nil {
		ctx.JSON(200, NewH("no authors", authors))
		return
	}
	ctx.JSON(200, NewH("success", authors))
}

// 新規登録
func (handler *AuthorHandler) InputAuthor(ctx *gin.Context) {
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
		ctx.JSON(302, NewH("Bad Request", author))
		return
	}

	user := handler.LoginUsecase.GetLoginUser(ctx)
	author.UserId = user.Id

	// カテゴリー検索(カテゴリー名で)
	authorByName := handler.Usecase.GetByName(author.Name, author.UserId)
	// TODO 空チェックできてる？
	if authorByName.Id != 0 {
		ctx.JSON(200, NewH("Exists", author))
		return

	}

	// 新規登録
	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.String("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	authorByName = handler.Usecase.Input(&author)

	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■authorByName.Id■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(authorByName.Id)

	ctx.JSON(200, NewH("success", author))
}

// 更新
func (handler *AuthorHandler) UpdateAuthor(ctx *gin.Context) {

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

	user := handler.LoginUsecase.GetLoginUser(ctx)
	author.UserId = user.Id

	authorByName := handler.Usecase.GetByName(author.Name, author.UserId)

	if authorByName.Id != 0 {
		ctx.JSON(200, NewH("Exists", author))
		return
	}

	logger.Info("++++++++++++++++++++++ author_handler.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.String("■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在する場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	handler.Usecase.Update(&author)

	ctx.JSON(200, NewH("success", author))
}

// 削除(id指定)
func (handler *AuthorHandler) DeleteAuthor(ctx *gin.Context) {
	author := domain.Author{}
	err := ctx.Bind(&author)
	if err != nil {
		ctx.JSON(302, NewH("Bad Request", author))
		return
	}

	// ログインユーザーID且つ、記事ID データが存在しなければ、302で返す
	user := handler.LoginUsecase.GetLoginUser(ctx)
	articleByIdAndUserId := handler.Usecase.GetByAuthorIdAndUserId(author.Id, user.Id)
	if articleByIdAndUserId.Id == 0 {
		ctx.JSON(302, NewH("Bad Request", author))
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

	handler.Usecase.Delete(&author, user.Id)

	// エラーじゃなければ(削除件数1以上)、紐づくarticle削除
	handler.ArticleUsecase.DeleteByAuthor(author.Id)

	// if article == nil {
	// 	ctx.JSON(500, NewH("no article", article))
	// 	return
	// }
	ctx.JSON(200, NewH("success", author.Id))
}
