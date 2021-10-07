package delivery

import (
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/usecase"
	"go_clean_arch_test/app/domain"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleHandler struct {
	Usecase       usecase.ArticleUsecase
	AuthorUsecase usecase.AuthorUsecase
}

func NewArticleHandler(db database.DB) *ArticleHandler {
	return &ArticleHandler{
		Usecase: usecase.ArticleUsecase{
			DB: &database.DBRepository{DB: db},
		},
		AuthorUsecase: usecase.AuthorUsecase{
			DB: &database.DBRepository{DB: db},
		},
	}
}

// 全件取得
// func GetAll(db *gorm.DB, ctx *gin.Context) {
// 	articles := usecase.GetAll(db)
// 	ctx.HTML(200, "test.html", gin.H{
// 		"articles": articles,
// 	})
// }
// 全件取得
func (handler *ArticleHandler) GetAll(ctx *gin.Context) {
	articles := handler.Usecase.GetAll()
	// ctx.JSON(res.StatusCode, NewH("success", articles))
	if len(articles) < 0 || articles == nil {
		ctx.JSON(500, NewH("no articles", articles))
		return
	}
	ctx.JSON(200, NewH("success", articles))
}

// 詳細取得(id指定)
func (handler *ArticleHandler) GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))

	// log
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
		zap.String("method", "GetById"),
		zap.Int("param id", id),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	article := handler.Usecase.GetById(id)
	// if article == nil {
	// 	ctx.JSON(500, NewH("no article", article))
	// 	return
	// }
	ctx.JSON(200, NewH("success", article))
}

// 新規登録
func (handler *ArticleHandler) Input(ctx *gin.Context) {

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
		ctx.JSON(302, NewH("Bad Request", article))
		return
	}

	handler.AuthorUsecase.Input(&article.Author)
	handler.Usecase.Input(&article)
	// if article == nil {
	// 	ctx.JSON(500, NewH("no article", article))
	// 	return
	// }
	ctx.JSON(200, NewH("success", article))
}

// 更新
func (handler *ArticleHandler) Update(ctx *gin.Context) {

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
		ctx.JSON(302, NewH("Bad Request", article))
		return
	}

	// TODO カテゴリーが変わってる場合
	// カテゴリー検索(カテゴリー名で)
	authorByName := handler.AuthorUsecase.GetByName(article.Author.Name)
	// TODO 空チェックできてる？
	if authorByName.Id == 0 {
		// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
		logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
			zap.String("method", "Update"),
			zap.String("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Input"),
			zap.Duration("elapsed", time.Now().Sub(oldTime)),
		)
		handler.AuthorUsecase.Input(&article.Author)
	} else {
		// カテゴリー存在したらそのIdで記事更新
		logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
			zap.String("method", "Update"),
			zap.String("■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在する場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Update"),
			zap.Duration("elapsed", time.Now().Sub(oldTime)),
		)
		handler.AuthorUsecase.Update(&article.Author)
	}

	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○テスト○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(&article)

	// TODO Author.ID更新されているか？
	handler.Usecase.Update(&article)
	// if article == nil {
	// 	ctx.JSON(500, NewH("no article", article))
	// 	return
	// }
	ctx.JSON(200, NewH("success", article))
}
