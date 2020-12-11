package log

import (
	"fmt"
	"os"
	"strconv"

	"go.uber.org/zap"
)

var _logger *zap.SugaredLogger
var _debugMode bool

func init() {
	// 从环境变量读取Debug模式
	if v, ok := os.LookupEnv("APP_DEBUG"); ok {
		if v, err := strconv.ParseBool(v); err == nil {
			_debugMode = v
		}
	}

	// 初始化日志记录器
	_logger = newZapLogger()
}

func Terminate() {
	err := _logger.Sync()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Terminate log error: ", err.Error())
	}
}

func Logger() *zap.SugaredLogger {
	return _logger
}

func IsDebug() bool {
	return _debugMode
}
