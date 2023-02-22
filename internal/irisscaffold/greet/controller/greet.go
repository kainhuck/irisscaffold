package controller

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
)

func (ctr *Controller) GreetHandler(ctx *context.Context) {
	var req request.GreetReq
	if err := ctx.ReadQuery(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(ctr.app.Greet(req))
}

func (ctr *Controller) LoginHandler(ctx *context.Context) {
	var req request.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(ctr.app.Login(req))
}

func (ctr *Controller) JwtDemoHandler(ctx *context.Context) {
	ctx.SendResponse(ctr.app.JwtDemo(ctx))
}

func (ctr *Controller) LogoutHandler(ctx *context.Context) {
	ctx.SendNoBodyResponse(ctr.app.Logout(ctx))
}
