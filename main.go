package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/pritam-is-next/resume/models"
	Server "github.com/vrianta/Server"
)

func main() {
	Server.New().InitDatabase().Start()
}
