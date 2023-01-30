package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func Init(level string, filePath string) {
	if filePath != "" {
		writer, _ := rotatelogs.New(
			filePath+".%Y%m%d",
			rotatelogs.WithLinkName(filePath),
			rotatelogs.WithRotationCount(10),
			rotatelogs.WithRotationTime(time.Hour*24),
		)
		logrus.SetOutput(writer)
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.TraceLevel
	}

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
