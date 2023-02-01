package greet

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/config"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/controller"
	"github.com/kainhuck/irisscaffold/internal/middleware"
	"github.com/kataras/iris/v12"
	"net/http"
)

func InitRoutes(app *iris.Application) {
	v1Route := app.Party("/api/v1")

	v1Route.Handle(http.MethodGet, "/hello", context.Handler(controller.GreetHandler))
	v1Route.Handle(http.MethodPost, "/login", context.Handler(controller.LoginHandler))

	v1Route.Use(middleware.JwtVerify(config.Cfg.Jwt.SigKey))
	v1Route.Handle(http.MethodGet, "/auth", context.Handler(controller.AuthHandler))
	v1Route.Handle(http.MethodPost, "/logout", context.Handler(controller.LogoutHandler))
}
