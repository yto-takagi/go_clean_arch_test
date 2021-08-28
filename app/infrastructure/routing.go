package infrastructure

import (
	"go_clean_arch_test/app/article/delivery"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB  *DB
	Gin *gin.Engine
}

func NewRouting(db *DB) *Routing {
	r := &Routing{
		DB:  db,
		Gin: gin.Default(),
	}
	// Corsの設定
	r.Gin.Use(cors.New(cors.Config{
		// 許可アクセス元
		AllowOrigins: []string{
			"http://localhost:61092",
		},
		// アクセス許可HTTPメソッド(以下PUTDELETEアクセス不可)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可HTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookie必要許可
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	// router.GET("/", func(ctx *gin.Context) {
	// 	delivery.GetAll(ctx)
	// })

	articleHandler := delivery.NewArticleHandler(r.DB)
	r.Gin.GET("/", func(ctx *gin.Context) { articleHandler.GetAll(ctx) })
	r.Gin.GET("/article", func(ctx *gin.Context) { articleHandler.GetById(ctx) })
}

func (r *Routing) Run() {
	r.Gin.Run()
}
