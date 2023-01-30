package bootstrap

import (
	"context"
	"fmt"
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

func NewApp(level string, init ...InitAppFunc) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(level)
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

	for _, f := range init {
		f(app)
	}

	return app
}

func Run(host string, port int, app *iris.Application) {
	pportEnv := os.Getenv("PPROF_PORT")
	pport, err := strconv.Atoi(pportEnv)
	if err == nil && port > 0 {
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

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
	}

	logrus.Fatalf("Run Error: %v", app.Run(iris.Server(srv), iris.WithoutInterruptHandler))
}
