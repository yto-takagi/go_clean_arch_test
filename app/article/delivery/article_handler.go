package delivery

import (
	"encoding/json"
	"go_clean_arch_test/app/article/database"
	"go_clean_arch_test/app/article/usecase"
	"go_clean_arch_test/app/domain"
	"go_clean_arch_test/app/domain/auth"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"go.uber.org/zap"
)

var LoginInfo auth.SessionInfo

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

	user := getLoginUser(ctx)
	log.Println("○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○User.ID○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○○")
	log.Println(user.Id)

	articles := handler.Usecase.GetAll(user.Id)
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

	user := getLoginUser(ctx)
	article.Author.UserId = user.Id

	// カテゴリー検索(カテゴリー名で)
	authorByName := handler.AuthorUsecase.GetByName(article.Author.Name, article.Author.UserId)
	// TODO 空チェックできてる？
	if authorByName.Id == 0 {
		// カテゴリー存在しなければ、カテゴリー新規登録してそのIdで記事更新
		logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
			zap.String("method", "Update"),
			zap.String("■■■■■■■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在しない場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Input"),
			zap.Duration("elapsed", time.Now().Sub(oldTime)),
		)
		authorByName = handler.AuthorUsecase.Input(&article.Author)
	} else {
		// カテゴリー存在したらそのIdで記事更新
		logger.Info("++++++++++++++++++++++ article_handler.go ++++++++++++++++++++++",
			zap.String("method", "Update"),
			zap.String("■■■■■■■■■■■■■■■■■■■■■■カテゴリー存在する場合■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■", "Update"),
			zap.Duration("elapsed", time.Now().Sub(oldTime)),
		)
		// handler.AuthorUsecase.Update(&article.Author)
	}

	log.Println("■■■■■■■■■■■■■■■■■■■■■■■■■■■■authorByName.Id■■■■■■■■■■■■■■■■■■■■■■■■■■■■")
	log.Println(authorByName.Id)

	article.Author.Id = authorByName.Id
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

	user := getLoginUser(ctx)
	article.Author.UserId = user.Id

	// TODO カテゴリーが変わってる場合
	// カテゴリー検索(カテゴリー名で)
	authorByName := handler.AuthorUsecase.GetByName(article.Author.Name, article.Author.UserId)
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

// 削除(id指定)
func (handler *ArticleHandler) Delete(ctx *gin.Context) {
	article := domain.Article{}
	err := ctx.Bind(&article)
	if err != nil {
		ctx.JSON(302, NewH("Bad Request", article))
		return
	}

	// ログインユーザーID且つ、記事ID データが存在しなければ、302で返す
	user := getLoginUser(ctx)
	articleByIdAndUserId := handler.Usecase.GetByIdAndUserId(article.Id, user.Id)
	if articleByIdAndUserId.Id == 0 {
		ctx.JSON(302, NewH("Bad Request", article))
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

	handler.Usecase.Delete(article.Id)
	// if article == nil {
	// 	ctx.JSON(500, NewH("no article", article))
	// 	return
	// }
	ctx.JSON(200, NewH("success", article.Id))
}

func getLoginUser(ctx *gin.Context) domain.User {
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
