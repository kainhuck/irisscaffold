package greet

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/controller"
	"github.com/kainhuck/irisscaffold/internal/middleware"
	"github.com/kataras/iris/v12"
	"net/http"
)

func InitRoutes(app *iris.Application, cfg *config.Config) {
	v1Route := app.Party("/api/v1")

	ctr := controller.NewController(cfg)

	v1Route.Handle(http.MethodGet, "/hello", context.Handler(ctr.GreetHandler))
	v1Route.Handle(http.MethodPost, "/login", context.Handler(ctr.LoginHandler))

	v1Route.Use(middleware.JwtVerify(cfg.Jwt.SigKey))
	{
		v1Route.Handle(http.MethodGet, "/auth", context.Handler(ctr.AuthHandler))
		v1Route.Handle(http.MethodPost, "/logout", context.Handler(ctr.LogoutHandler))
	}
}
