package commonUtils

import "os"

const (
	EnvPort       string = "PORT"
	MysqlUsername        = "MYSQL_USERNAME"
	MysqlPassword        = "MYSQL_PASSWORD"
	MysqlHost            = "MYSQL_HOST"
	MysqlPort            = "MYSQL_PORT"
	MysqlDatabase        = "MYSQL_DATABASE"
)

func GetEnvValue(key string) string {
	return os.Getenv(key)
}
