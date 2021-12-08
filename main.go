package main

import (
	"go_clean_arch_test/app/infrastructure"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.Run()
}
