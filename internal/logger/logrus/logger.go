package logrus

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger(level string, file string) *Logger {
	lvl, _ := logrus.ParseLevel(level)

	log := logrus.New()

	log.SetLevel(lvl)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})

	if file != "" {
		writer, _ := rotatelogs.New(
			file+".%Y%m%d",
			rotatelogs.WithLinkName(file),
			rotatelogs.WithRotationCount(10),
			rotatelogs.WithRotationTime(time.Hour*24),
		)
		log.SetOutput(writer)
	}

	return &Logger{
		log,
	}
}

func (l *Logger) GetLevel() string {
	return l.Level.String()
}
