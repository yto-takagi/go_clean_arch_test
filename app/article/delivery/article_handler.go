package delivery

import (
	"go_clean_arch_test/app/article/usecase"
	loginUsecase "go_clean_arch_test/app/article/usecase/auth"
	"go_clean_arch_test/app/domain"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const NO_ARTICLES = "no articles"

// ArticleHandler interface
type ArticleHandler interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetByAuthorId(ctx *gin.Context)
	GetLikeByTitleAndContent(ctx *gin.Context)
	Input(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type articleHandler struct {
	articleusecase usecase.ArticleUsecase
	authorUsecase  usecase.AuthorUsecase
	loginUsecase   loginUsecase.LoginUsecase
}

// NewArticleHandler constructor
func NewArticleHandler(articleusecase usecase.ArticleUsecase, authorUsecase usecase.AuthorUsecase, loginUsecase loginUsecase.LoginUsecase) ArticleHandler {
	return &articleHandler{articleusecase: articleusecase, authorUsecase: authorUsecase, loginUsecase: loginUsecase}
}

func (articleHandler *articleHandler) GetAll(ctx *gin.Context) {

	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○User.ID○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(user.Id)

	articles, err := articleHandler.articleusecase.GetAll(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), articles))
		return
	}
	if len(articles) < 0 || articles == nil {
		ctx.JSON(http.StatusOK, NewH(NO_ARTICLES, articles))
		return
	}
	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articles))
}

// 詳細取得(id指定)
func (articleHandler *articleHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article, err := articleHandler.articleusecase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article))
}

// 詳細取得(authorId,ユーザーid指定)
func (articleHandler *articleHandler) GetByAuthorId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "GetByAuthorId"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article, err := articleHandler.articleusecase.GetByAuthorIdAndUserId(id, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}
	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article))
}

// 検索
func (articleHandler *articleHandler) GetLikeByTitleAndContent(ctx *gin.Context) {
	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "GetLikeByTitleAndContent"),
		zap.String("param searchContent", ctx.Query("content")),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	articles, err := articleHandler.articleusecase.GetLikeByTitleAndContent(ctx.Query("content"), user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), articles))
		return
	}

	if len(articles) == 0 || articles == nil || articles[0].Id == 0 {
		ctx.JSON(http.StatusOK, NewH(NO_ARTICLES, articles))
		return
	}
	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articles))
}

// 新規登録
func (articleHandler *articleHandler) Input(ctx *gin.Context) {
	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "Input"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article := domain.Article{}
	err := ctx.Bind(&article)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), article))
		return
	}

	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	article.Author.UserId = user.Id

	err = articleHandler.articleusecase.Input(ctx, &article)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article))
}

// 更新
func (articleHandler *articleHandler) Update(ctx *gin.Context) {

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "Update"),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article := domain.Article{}
	err := ctx.Bind(&article)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), article))
		return
	}

	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	article.Author.UserId = user.Id

	// // TODO カテゴリーが変わってる場合
	// // カテゴリー検索(カテゴリー名で)
	// authorByName, err := articleHandler.authorUsecase.GetByName(article.Author.Name, article.Author.UserId)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
	// 	return
	// }
	// // TODO 空チェックできてる？
	// if authorByName.Id == 0 {
	// 	// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
	// 	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
	// 		zap.String("method", "Update"),
	// 		zap.String("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Input"),
	// 		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	// 	)
	// 	articleHandler.authorUsecase.Input(&article.Author)
	// } else {
	// 	// カテゴリー存在したらそのIdで記事更新
	// 	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
	// 		zap.String("method", "Update"),
	// 		zap.String("■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在する場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Update"),
	// 		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	// 	)
	// 	articleHandler.authorUsecase.Update(&article.Author)
	// }

	// log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○テスト○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	// log.Println(&article)

	// TODO Author.ID更新されているか？
	err = articleHandler.articleusecase.Update(ctx, &article)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article))
}

// 削除(id指定)
func (articleHandler *articleHandler) Delete(ctx *gin.Context) {
	article := domain.Article{}
	err := ctx.Bind(&article)
	if err != nil {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), article))
		return
	}

	// ログインユーザーID且つ、記事ID データが存在しなければ、302で返す
	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}
	articleByIdAndUserId, err := articleHandler.articleusecase.GetByIdAndUserId(article.Id, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}
	if articleByIdAndUserId.Id == 0 {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), article))
		return
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "Delete"),
		zap.Int("param id", article.Id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	err = articleHandler.articleusecase.Delete(article.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article.Id))
}
