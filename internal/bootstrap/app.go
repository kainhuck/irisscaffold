package bootstrap

import (
	"context"
	"fmt"
	"github.com/kainhuck/irisscaffold/internal/configx"
	"github.com/kainhuck/irisscaffold/internal/logger"
	irisCtx "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/cors"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

type InitAppFunc = func(app *iris.Application)

func NewApp(log logger.Logger) *iris.Application {
	app := iris.New()
	app.Logger().SetLevel(log.GetLevel())
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

	app.Use(func(ctx *irisCtx.Context) {
		start := time.Now()
		body, _ := io.ReadAll(ctx.Request().Body)

		ctx.Next()

		log.Infof("[IRIS_LOG] request_id: (%v), remote_addr: (%v), [%v] url: %v, body: %s, time_cost: %dms",
			ctx.GetID(), ctx.Request().RemoteAddr, ctx.Request().Method, ctx.Request().URL.String(), body, time.Since(start).Nanoseconds()/1e6)
	})

	return app
}

func Run(log logger.Logger, cfg configx.ServiceInfo, app *iris.Application) {
	pportEnv := os.Getenv("PPROF_PORT")
	pport, err := strconv.Atoi(pportEnv)
	if err == nil && pport > 0 {
		go func() {
			log.Infof("Listen pprof at 0.0.0.0:%d\n", pport)
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
		log.Errorf("Run Error: %v", app.Run(iris.Addr(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)), iris.WithoutInterruptHandler))
	} else if cfg.Schema == "https" {
		log.Errorf("Run Error: %v", app.Run(iris.TLS(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), cfg.CertFile, cfg.KeyFile), iris.WithoutInterruptHandler))
	} else {
		log.Errorf("Unknown schema: %s", cfg.Schema)
	}
}
