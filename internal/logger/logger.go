package logger

import (
	"github.com/kainhuck/irisscaffold/internal/configx"
	"github.com/kainhuck/irisscaffold/internal/logger/logrus"
)

type Logger interface {
	GetLevel() string

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
}

func NewLogger(cfg configx.LogConfig) Logger {
	return logrus.NewLogger(cfg.LogLevel, cfg.FilePath)
}
