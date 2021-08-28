package delivery

import (
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleHandler struct {
	Usecase usecase.ArticleUsecase
}

func NewArticleHandler(db database.DB) *ArticleHandler {
	return &ArticleHandler{
		Usecase: usecase.ArticleUsecase{
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

// id指定
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
