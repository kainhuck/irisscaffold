package logger

import (
	"bytes"
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"github.com/kataras/iris/v12"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func Init(cfg configx.LogConfig) {
	if cfg.FilePath != "" {
		writer, _ := rotatelogs.New(
			cfg.FilePath+".%Y%m%d",
			rotatelogs.WithLinkName(cfg.FilePath),
			rotatelogs.WithRotationCount(10),
			rotatelogs.WithRotationTime(time.Hour*24),
		)
		logrus.SetOutput(writer)
	}

	lvl, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		lvl = logrus.TraceLevel
	}

	logrus.SetLevel(lvl)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Handler(ctx iris.Context) {
	method := ctx.Request().Method
	start := time.Now()
	fields := make(map[string]interface{})
	fields["type"] = "IrisLog"
	fields["ip"] = ctx.Request().RemoteAddr
	fields["method"] = method
	fields["url"] = ctx.Request().URL.String()
	fields["proto"] = ctx.Request().Proto
	fields["request_id"] = ctx.GetID()

	// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body, err := io.ReadAll(ctx.Request().Body)
		if err == nil {
			defer ctx.Request().Body.Close()
			buf := bytes.NewBuffer(body)
			ctx.Request().Body = io.NopCloser(buf)
			fields["content_length"] = ctx.GetContentLength()
			fields["body"] = string(body)
		}
	}
	ctx.Next()

	//下面是返回日志
	fields["res_status"] = ctx.ResponseWriter().StatusCode()
	timeConsuming := time.Since(start).Nanoseconds() / 1e6
	fields["time_cost"] = fmt.Sprintf("%dms", timeConsuming)
	logrus.WithFields(fields).Infof("%s %s %d", ctx.Request().Method, ctx.Request().URL.Path, ctx.ResponseWriter().StatusCode())
}
