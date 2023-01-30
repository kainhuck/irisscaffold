package greet

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/controller"
	"github.com/kataras/iris/v12"
	"net/http"
)

func InitRoutes(app *iris.Application) {
	v1Route := app.Party("/api/v1")

	v1Route.Handle(http.MethodGet, "/hello", context.Handler(controller.GreetHandler))
}
