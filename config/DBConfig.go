package config

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/vamika-digital/wms-lib/logger"
	"sync"
)

var lock = &sync.Mutex{}
var dbClient *sqlx.DB

func GetDBInstance() *sqlx.DB {
	if dbClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbClient == nil {
			driverName, dbDataSource, err := getDataSource()
			if err != nil {
				logger.Fatal("error in database source")
			}
			dbClient, err = sqlx.Open(driverName, dbDataSource)
			if err != nil {
				panic(err)
			} else {
				logger.Info(fmt.Sprintf("Database connection success with data source %s", dbDataSource))
			}
			viper.GetInt(fmt.Sprintf("rdbms.%s.connMaxLifetime", driverName))
			viper.GetInt(fmt.Sprintf("rdbms.%s.maxOpenConns", driverName))
			viper.GetInt(fmt.Sprintf("rdbms.%s.maxIdleConns", driverName))

			dbClient.SetConnMaxLifetime(viper.GetDuration(fmt.Sprintf("rdbms.%s.connMaxLifetime", driverName)))
			dbClient.SetMaxOpenConns(viper.GetInt(fmt.Sprintf("rdbms.%s.maxOpenConns", driverName)))
			dbClient.SetMaxIdleConns(viper.GetInt(fmt.Sprintf("rdbms.%s.maxIdleConns", driverName)))
		} else {
			logger.Info("Single instance already created.")
		}
	} else {
		logger.Info("Single instance already created.")
	}
	return dbClient
}

func getDataSource() (string, string, error) {
	dbType := viper.GetString("rdbms.type")

	host := viper.GetString(fmt.Sprintf("rdbms.%s.host", dbType))
	port := viper.GetString(fmt.Sprintf("rdbms.%s.port", dbType))
	username := viper.GetString(fmt.Sprintf("rdbms.%s.username", dbType))
	password := viper.GetString(fmt.Sprintf("rdbms.%s.password", dbType))
	dbname := viper.GetString(fmt.Sprintf("rdbms.%s.dbname", dbType))
	dbPath := viper.GetString(fmt.Sprintf("rdbms.%s.dbpath", dbType))

	if dbType == "sqlite3" {
		return dbType, dbPath, nil
	} else if dbType == "mysql" {
		return dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true&sql_mode=TRADITIONAL", username, password, host, port, dbname), nil
	} else if dbType == "pgsql" {
		return dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=skip-verify&autocommit=true&sql_mode=TRADITIONAL", username, password, host, port, dbname), nil
	} else {
		return "", "", errors.New("unknown SQL drivers")
	}
}
