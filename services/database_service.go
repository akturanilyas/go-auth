package databaseService

import (
	"fmt"
	commonUtils "go-test/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getMySQLDatabaseConnection() (*gorm.DB, error) {
	// TODO: Move to config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		commonUtils.GetEnvValue(commonUtils.MysqlUsername),
		commonUtils.GetEnvValue(commonUtils.MysqlPassword),
		commonUtils.GetEnvValue(commonUtils.MysqlHost),
		commonUtils.GetEnvValue(commonUtils.MysqlPort),
		commonUtils.GetEnvValue(commonUtils.MysqlDatabase),
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
