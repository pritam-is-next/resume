package main

import (
	Server "github.com/vrianta/Server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	Server.New().InitDatabase().Start()
}
