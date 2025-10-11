package main

import (
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	agai "github.com/vrianta/agai/v1"
)

func main() {
	runtime.GOMAXPROCS(1)
	app := agai.New() // creating new application

	app.Start()
}
