package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	// 设置日志格式为 JSON，方便结构化日志处理
	log.SetFormatter(&log.JSONFormatter{})

	// 默认将日志输出到标准输出（控制台）
	log.SetOutput(os.Stdout)

	// 根据环境变量 ENV 设置日志级别
	env := os.Getenv("ENV")
	if env == "prod" {
		log.SetLevel(log.WarnLevel) // 生产环境仅记录警告及以上日志
	} else {
		log.SetLevel(log.DebugLevel) // 开发环境记录调试级别及以上日志
	}
}

var Logger = log.WithFields(log.Fields{
	"ENV": os.Getenv("ENV"),
})
