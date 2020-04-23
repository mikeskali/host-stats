package log

import (
"go.uber.org/zap"
"strings"
)

func newLogger(conf string) *zap.SugaredLogger {
	var logger *zap.SugaredLogger

	switch strings.ToLower(env) {
	case "dev", "test", "debug":
		l, _ := zap.NewDevelopment()
		logger = l.Sugar()
	default:
		l, _ := zap.NewProduction()
		logger = l.Sugar()
	}
	logger.Named(serviceName)

	return logger
}

