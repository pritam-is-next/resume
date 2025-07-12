package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/pritam-is-next/resume/models"
	"github.com/vrianta/agai/v1/server"
)

func main() {

	server.Start()
}
