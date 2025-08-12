package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/vrianta/agai/v1/config"
	log "github.com/vrianta/agai/v1/log"
)

// Function to init the Database with the Database/sql object and store it in the program
func Init() {

	if config.GetDatabaseConfig().Host == "" {
		log.WriteLog("DataBase Config do not have any host in it so We are skipping all database connections")
		return
	}
	var err error
	if database, err = sql.Open(config.GetDatabaseDriver(), config.GetDSN()); err != nil {
		if config.ShowDsn {
			fmt.Println(config.GetDSN())
		}
		panic("[ERROR] - DB Connection Failed Due to: " + err.Error())
	}

	// TODO : Create a Config Element to get the Desired Detail
	database.SetConnMaxIdleTime(time.Second * 10)
	database.SetMaxOpenConns(10000)
	if err := database.Ping(); err != nil {
		if config.ShowDsn {
			log.Info("DNS of the Database Server : %s", config.GetDSN())
		}
		panic("[ERROR] - DB Connection Failed Due to: " + err.Error())
	}

	Initialized = true
	fmt.Println("[Info] DataBase Connection Established Successfully")
}

func GetDatabase() (*sql.DB, error) {
	if !Initialized {
		panic("[ERROR] - You are calling for the Database but Connection with Database is not established properly")
	}
	if database == nil {
		panic("[ERROR] - You are calling for the Database but Connection with Database is not established properly")
	}
	if err := database.Ping(); err != nil {
		log.Error("Failed to ping the DB Server: %s", err.Error())
		return nil, err
	}
	return database, nil
}
