package logger

import (
	"encoding/json"
	"github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
	"schedule-management/config"
)

var cfg = config.GetConfig()

func NewLogger() {

	if cfg.Debug == false {
		hook := graylog.NewAsyncGraylogHook("127.0.0.1:12201", nil)
		defer hook.Flush()
		log.AddHook(hook)
	}
	// Log to file
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }
}

type LogFields struct {
	Host string `json:"host"`
	Source string `json:"source"`
	FullMessage string `json:"full_message"`
}

func structToMap(logFields LogFields) (result map[string]interface{}) {
	tmp, _ := json.Marshal(logFields)
	_ = json.Unmarshal(tmp, &result)
	return
}

func Info(msg string, fields LogFields) {
	mapFields := structToMap(fields)
	log.WithFields(mapFields).Info(msg)
}

func Error(msg string) {
	log.Error(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Debug(msg string) {
	log.Debug(msg)
}

func CustomInfo(msg string) {
	log.WithFields(log.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	}).Info(msg)
}

func CustomError(msg string) {
	log.WithFields(log.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	}).Error(msg)
}