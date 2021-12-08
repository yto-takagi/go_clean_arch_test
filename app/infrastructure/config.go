package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DB struct {
		Production struct {
			Dbms     string
			User     string
			Pass     string
			Protocol string
			Dbname   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Dbms = "mysql"
	c.DB.Production.User = "go_clean_arch_test"
	c.DB.Production.Pass = "password"
	c.DB.Production.Protocol = "tcp(db:3306)"
	c.DB.Production.Dbname = "go_clean_arch"

	c.Routing.Port = ":5000"

	return c
}
