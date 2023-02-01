package controller

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/irisscaffold/greet/application"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
)

func GreetHandler(ctx *context.Context) {
	var req request.GreetReq
	if err := ctx.ReadQuery(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(application.Greet(req))
}

func LoginHandler(ctx *context.Context) {
	var req request.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(application.Login(req))
}

func AuthHandler(ctx *context.Context) {
	ctx.SendResponse(application.Auth(ctx))
}

func LogoutHandler(ctx *context.Context) {
	ctx.SendNoBodyResponse(application.Logout(ctx))
}
