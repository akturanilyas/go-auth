package database

import (
	"auth-go/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func SetDBConnection(conn *gorm.DB) {
	DB = conn
}

func getMySQLDatabaseConnection() (*gorm.DB, error) {
	// TODO: Move to config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		utils.GetEnvValue(utils.MysqlUsername),
		utils.GetEnvValue(utils.MysqlPassword),
		utils.GetEnvValue(utils.MysqlHost),
		utils.GetEnvValue(utils.MysqlPort),
		utils.GetEnvValue(utils.MysqlDatabase),
		"parseTime=true",
	)

	fmt.Println(dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return conn, err
}

func Connect() *gorm.DB {
	conn, err := getMySQLDatabaseConnection()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully")

	return conn
}
