package database

import (
	"schedule-management/database/mysql"
	"schedule-management/database/redis"
)

var (
	MysqlConn *mysql.MysqlConn
	RedisConn *redis.RedisConn
)