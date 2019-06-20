package setting

import (
	"schedule-management/config"
	"schedule-management/database"
	"schedule-management/database/mysql"
	"schedule-management/database/redis"
	"schedule-management/logger"
)

var cfg = config.GetConfig()

func InitMysql() (err error) {
	database.MysqlConn, err = mysql.NewConn(cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.DatabaseName, cfg.Mysql.MaxIdleConn, cfg.Mysql.MaxOpenConn)
	if err != nil {
		logger.Error("Connect database fail")
	} else {
		logger.Info("Connect database success", logger.LogFields{})
	}
	return
}

func InitRedis() (err error) {
	database.RedisConn, err = redis.NewConn(cfg.Redis.IP, cfg.Redis.Port, cfg.Redis.Password)
	if err != nil {
		logger.Error("Connect redis fail")
	} else {
		logger.Info("Connect redis success", logger.LogFields{})
	}
	return
}

func InitLogger() (err error) {
	logger.NewLogger()
	return
}