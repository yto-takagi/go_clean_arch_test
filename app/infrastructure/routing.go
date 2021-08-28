package infrastructure

import (
	"go_clean_arch_test/app/article/delivery"

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
