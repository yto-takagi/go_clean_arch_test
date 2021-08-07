package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Dbms       string
	User       string
	Pass       string
	Protocol   string
	Dbname     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Dbms:     c.DB.Production.Dbms,
		User:     c.DB.Production.User,
		Pass:     c.DB.Production.Pass,
		Protocol: c.DB.Production.Protocol,
		Dbname:   c.DB.Production.Dbname,
	})
}

func newDB(d *DB) *DB {
	DBMS := "mysql"
	USER := "go_clean_arch_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_clean_arch"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	d.Connection = db
	return d
}

func SqlConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := "go_clean_arch_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_clean_arch"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}

	return db
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
