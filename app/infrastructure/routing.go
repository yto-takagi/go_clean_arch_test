package infrastructure

import (
	"go_clean_arch_test/app/article/delivery"
	authDelivery "go_clean_arch_test/app/article/delivery/auth"
	middleware "go_clean_arch_test/app/infrastructure/auth"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
			"http://localhost:54386",
		},
		// AllowAllOrigins: true,
		// アクセス許可HTTPメソッド(以下PUT,DELETEアクセス不可)
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
			"accessToken",
			"Set-Cookie",
			"Cookie",
		},
		// AllowOrigins: []string{"*"},
		// AllowMethods: []string{"*"},
		// AllowHeaders: []string{"*"},
		// cookie必要許可
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// セッションCookieの設定
	// secure属性がtrueになっているため、httpホストでcookie情報を取得できていない？
	// cookieのsamesiteをnoneにする必要がありそう？
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Secure:   false,
		HttpOnly: false})
	r.Gin.Use(sessions.Sessions("bookMarkAppSessKey", store))

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	// router.GET("/", func(ctx *gin.Context) {
	// 	delivery.GetAll(ctx)
	// })

	signupHandler := authDelivery.NewSignUpHandler(r.DB)
	loginHandler := authDelivery.NewLoginHandler(r.DB)
	articleHandler := delivery.NewArticleHandler(r.DB)
	r.Gin.POST("/signup", func(ctx *gin.Context) { signupHandler.SignUp(ctx) })
	r.Gin.POST("/login", func(ctx *gin.Context) { loginHandler.Login(ctx) })
	r.Gin.POST("/logout", func(ctx *gin.Context) { authDelivery.Logout(ctx) })
	// 認証済のみアクセス可能なグループ
	authUserGroup := r.Gin.Group("/auth")
	authUserGroup.Use(middleware.LoginCheckMiddleware())
	{
		r.Gin.GET("/", func(ctx *gin.Context) { articleHandler.GetAll(ctx) })
		r.Gin.GET("/article", func(ctx *gin.Context) { articleHandler.GetById(ctx) })
		r.Gin.GET("/article/search", func(ctx *gin.Context) { articleHandler.GetLikeByTitleAndContent(ctx) })
		r.Gin.POST("/article/input", func(ctx *gin.Context) { articleHandler.Input(ctx) })
		r.Gin.POST("/article/update", func(ctx *gin.Context) { articleHandler.Update(ctx) })
		r.Gin.POST("/article/delete", func(ctx *gin.Context) { articleHandler.Delete(ctx) })
	}
}

func (r *Routing) Run() {
	r.Gin.Run()
}
