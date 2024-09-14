package utils

import (
	"os"
)

const (
	EnvPort       string = "PORT"
	MysqlUsername string = "MYSQL_USERNAME"
	MysqlPassword string = "MYSQL_PASSWORD"
	MysqlHost     string = "MYSQL_HOST"
	MysqlPort     string = "MYSQL_PORT"
	MysqlDatabase string = "MYSQL_DATABASE"
)

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
