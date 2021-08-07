package main

import (
	"go_clean_arch_test/app/infrastructure"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// type User struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// }

func main() {
	// db := db_con.SqlConnect()
	// defer db.Close()

	// router := gin.Default()
	// router.LoadHTMLGlob("templates/*.html")
	// ログテスト
	log.Println("++++++++++++++++++++++ test ++++++++++++++++++++++")

	// ログテスト
	oldTime := time.Now()
	logger, _ := zap.NewProduction()
	logger.Info("++++++++++++++++++++++ test222 ++++++++++++++++++++++",
		zap.String("path", "test"),
		zap.String("Ua", "test"),
		zap.Int("status", 111),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)

	// router := gin.Default()
	// router.LoadHTMLGlob("templates/*.html")
	// infrastructure.SetRouting(router)

	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "index.html", gin.H{
	// 		// "users": users,
	// 	})
	// })
	// router.Run()

	// db := infrastructure.NewDB()
	// r := infrastructure.NewRouting(db)
	// r.Run()

	// router.GET("/", func(ctx *gin.Context) {
	// 	db := db_connect.sqlConnect()
	// 	var users []User
	// 	db.Order("created_at asc").Find(&users)
	// 	defer db.Close()

	// 	ctx.HTML(200, "index.html", gin.H{
	// 		"users": users,
	// 	})
	// })

	// router.POST("/new", func(ctx *gin.Context) {
	// 	db := sqlConnect()
	// 	name := ctx.PostForm("name")
	// 	email := ctx.PostForm("email")
	// 	fmt.Println("create user " + name + " with email " + email)
	// 	db.Create(&User{Name: name, Email: email})
	// 	defer db.Close()

	// 	ctx.Redirect(302, "/")
	// })

	// router.POST("/delete/:id", func(ctx *gin.Context) {
	// 	db := sqlConnect()
	// 	n := ctx.Param("id")
	// 	id, err := strconv.Atoi(n)
	// 	if err != nil {
	// 		panic("id is not a number")
	// 	}
	// 	var user User
	// 	db.First(&user, id)
	// 	db.Delete(&user)
	// 	defer db.Close()

	// 	ctx.Redirect(302, "/")
	// })

	// router.Run()
}

// func sqlConnect() (database *gorm.DB) {
// 	DBMS := "mysql"
// 	USER := "go_clean_arch_test"
// 	PASS := "password"
// 	PROTOCOL := "tcp(db:3306)"
// 	DBNAME := "go_clean_arch"

// 	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

// 	count := 0
// 	db, err := gorm.Open(DBMS, CONNECT)
// 	if err != nil {
// 		for {
// 			if err == nil {
// 				fmt.Println("")
// 				break
// 			}
// 			fmt.Print(".")
// 			time.Sleep(time.Second)
// 			count++
// 			if count > 180 {
// 				fmt.Println("")
// 				panic(err)
// 			}
// 			db, err = gorm.Open(DBMS, CONNECT)
// 		}
// 	}

// 	return db
// }
