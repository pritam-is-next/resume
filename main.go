package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/pritam-is-next/resume/models"
	server "github.com/vrianta/Server"
)

func main() {

	server.Start()
}
