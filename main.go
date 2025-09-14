package main

import (
	_ "github.com/go-sql-driver/mysql"
	agai "github.com/vrianta/agai/v1"
)

func main() {
	agai.New() // creating new application
}
