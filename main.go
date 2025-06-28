package main

import (
	Server "github.com/vrianta/Server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// err, _sql := sql.Open(Config.GetDatabaseDriver(), Config.GetDSN())
	Server.New().InitDatabase().Start()
}
