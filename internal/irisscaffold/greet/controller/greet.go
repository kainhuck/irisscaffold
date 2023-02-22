package controller

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
)

// GreetHandler
// @Summary Greet
// @Description this is Greet
// @Tags iris
// @Produce application/json
// @Param object query request.GreetReq true "args"
// @Success 200 {object} response.GreetResp
// @Router /hello [get]
func (ctr *Controller) GreetHandler(ctx *context.Context) {
	var req request.GreetReq
	if err := ctx.ReadQuery(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(ctr.app.Greet(req))
}

// LoginHandler
// @Summary Login
// @Description this is login
// @Tags iris
// @Accept application/json
// @Produce application/json
// @Param object body request.LoginReq true "args"
// @Success 200 {object} response.LoginResp
// @Router /login [post]
func (ctr *Controller) LoginHandler(ctx *context.Context) {
	var req request.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		ctx.SendNoBodyResponse(e.ErrParameter, err)
		return
	}

	ctx.SendResponse(ctr.app.Login(req))
}

// JwtDemoHandler
// @Summary JwtDemo
// @Description this is JwtDemo
// @Tags iris
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.JwtDemoResp
// @Router /jwt/demo [get]
func (ctr *Controller) JwtDemoHandler(ctx *context.Context) {
	ctx.SendResponse(ctr.app.JwtDemo(ctx))
}

// LogoutHandler
// @Summary Logout
// @Description this is Logout
// @Tags iris
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Router /logout [post]
func (ctr *Controller) LogoutHandler(ctx *context.Context) {
	ctx.SendNoBodyResponse(ctr.app.Logout(ctx))
}
