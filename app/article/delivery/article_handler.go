package delivery

import (
	"go_clean_arch_test/app/article/delivery/request"
	"go_clean_arch_test/app/article/delivery/response"
	"go_clean_arch_test/app/article/usecase"
	loginUsecase "go_clean_arch_test/app/article/usecase/auth"
	"go_clean_arch_test/app/domain"
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

	setArticleModel(article request.Article) domain.Article
	setArticleResponse(domain.Article) response.Article
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

// ユーザーのデータ全件取得
func (articleHandler *articleHandler) GetAll(ctx *gin.Context) {

	user, err := articleHandler.loginUsecase.GetLoginUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), nil))
		return
	}

	articles, err := articleHandler.articleusecase.GetAll(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), articles))
		return
	}
	if len(articles) < 0 || articles == nil {
		ctx.JSON(http.StatusOK, NewH(NO_ARTICLES, articles))
		return
	}

	var articleResponseSlice []response.Article
	for _, article := range articles {
		articleResponse := articleHandler.setArticleResponse(article)
		articleResponseSlice = append(articleResponseSlice, articleResponse)
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articleResponseSlice))
}

// 詳細取得(id指定)
func (articleHandler *articleHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetById",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article, err := articleHandler.articleusecase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	articleResponse := articleHandler.setArticleResponse(article)

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articleResponse))
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
	logger.Info("GetByAuthorId",
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	articles, err := articleHandler.articleusecase.GetByAuthorIdAndUserId(id, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), articles))
		return
	}

	var articleResponseSlice []response.Article
	for _, article := range articles {
		articleResponse := articleHandler.setArticleResponse(article)
		articleResponseSlice = append(articleResponseSlice, articleResponse)
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articleResponseSlice))
}

// 検索
func (articleHandler *articleHandler) GetLikeByTitleAndContent(ctx *gin.Context) {
	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("GetLikeByTitleAndContent",
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

	if len(articles) == 0 || articles == nil || articles[0].GetId() == 0 {
		ctx.JSON(http.StatusOK, NewH(NO_ARTICLES, articles))
		return
	}

	var articleResponseSlice []response.Article
	for _, article := range articles {
		articleResponse := articleHandler.setArticleResponse(article)
		articleResponseSlice = append(articleResponseSlice, articleResponse)
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articleResponseSlice))
}

// 新規登録
func (articleHandler *articleHandler) Input(ctx *gin.Context) {
	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("Input",
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article := request.Article{}
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
	articleModel := articleHandler.setArticleModel(article)

	err = articleHandler.articleusecase.Input(ctx, &articleModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), articleModel))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), articleModel))
}

// 更新
func (articleHandler *articleHandler) Update(ctx *gin.Context) {

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("Update",
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article := request.Article{}
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
	articleModel := articleHandler.setArticleModel(article)

	err = articleHandler.articleusecase.Update(ctx, &articleModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}

	ctx.JSON(http.StatusOK, NewH(http.StatusText(http.StatusOK), article))
}

// 削除(id指定)
func (articleHandler *articleHandler) Delete(ctx *gin.Context) {
	article := request.Article{}
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
	articleByIdAndUserId, err := articleHandler.articleusecase.GetByIdAndUserId(article.Id, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewH(err.Error(), article))
		return
	}
	if articleByIdAndUserId.GetId() == 0 {
		ctx.JSON(http.StatusFound, NewH(http.StatusText(http.StatusFound), article))
		return
	}

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("Delete",
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

func (articleHandler *articleHandler) setArticleModel(article request.Article) domain.Article {
	authorModel, _ := domain.NewAuthor(article.Author.Id, article.Author.Name, article.Author.UserId, article.Author.UpdatedAt, article.Author.CreatedAt)
	expPoolModel, _ := domain.NewExpPool(article.ExpPool.Id, article.ExpPool.UserId, article.ExpPool.Exp, article.ExpPool.Lv, article.ExpPool.UpdatedAt, article.ExpPool.CreatedAt)
	articleModel, _ := domain.NewArticle(article.Id, article.Title, article.Content, article.UpdatedAt, article.CreatedAt, *authorModel, *expPoolModel)
	return *articleModel
}

func (articleHandler *articleHandler) setArticleResponse(article domain.Article) response.Article {
	authorResponse := response.NewAuthor(article.Author.GetId(), article.Author.GetName(), article.Author.GetUserId(), article.Author.GetUpdatedAt(), article.Author.GetCreatedAt())
	expPoolResponse := response.NewExpPool(article.ExpPool.GetId(), article.ExpPool.GetUserId(), article.ExpPool.GetExp(), article.ExpPool.GetLv(), article.ExpPool.GetUpdatedAt(), article.ExpPool.GetCreatedAt())
	articleResponse := response.NewArticle(article.GetId(), article.GetTitle(), article.GetContent(), article.GetUpdatedAt(), article.GetCreatedAt(), *authorResponse, *expPoolResponse)
	return *articleResponse
}
