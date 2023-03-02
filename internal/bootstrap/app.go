package bootstrap

import (
	"context"
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"github.com/kainhuck/irisscaffold/internal/logger"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type InitAppFunc = func(app *iris.Application)

func NewApp(logLevel string) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(logLevel)
	app.UseRouter(requestid.New())
	app.UseRouter(recover.New())
	app.UseRouter(cors.New().
		ExtractOriginFunc(cors.DefaultOriginExtractor).
		ReferrerPolicy(cors.NoReferrerWhenDowngrade).
		AllowOriginFunc(cors.AllowAnyOrigin).
		Handler())
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_ = app.Shutdown(ctx)
	})

	app.Use(logger.Handler)

	return app
}

func Run(cfg configx.ServiceInfo, app *iris.Application) {
	pportEnv := os.Getenv("PPROF_PORT")
	pport, err := strconv.Atoi(pportEnv)
	if err == nil && pport > 0 {
		go func() {
			logrus.Infof("Listen pprof at 0.0.0.0:%d\n", pport)
			_ = http.ListenAndServe(fmt.Sprintf(":%v", pport), nil)
		}()
	}

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		//关闭所有主机
		_ = app.Shutdown(ctx)
	})

	if cfg.Schema == "http" {
		logrus.Fatalf("Run Error: %v", app.Run(iris.Addr(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)), iris.WithoutInterruptHandler))
	} else if cfg.Schema == "https" {
		logrus.Fatalf("Run Error: %v", app.Run(iris.TLS(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), cfg.CertFile, cfg.KeyFile), iris.WithoutInterruptHandler))
	} else {
		logrus.Fatalf("Unkonwn schema: %s", cfg.Schema)
	}
}
