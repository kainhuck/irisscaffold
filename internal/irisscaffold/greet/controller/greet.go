package controller

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/errno"
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
func (ctr *Controller) GreetHandler(ctx *context.Context) (interface{}, error) {
	var req request.GreetReq
	if err := ctx.ReadQuery(&req); err != nil {
		return nil, errno.ErrParameter.WithErr(err)
	}

	return ctr.app.Greet(req)
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
func (ctr *Controller) LoginHandler(ctx *context.Context) (interface{}, error) {
	var req request.LoginReq
	if err := ctx.ReadJSON(&req); err != nil {
		return nil, errno.ErrParameter.WithErr(err)
	}

	return ctr.app.Login(req)
}

// WebsocketHandler
// @Summary websocket
// @Description this is websocket
// @Tags iris
// @Param Authorization header string true "Bearer token"
// @Router /ws [get]
func (ctr *Controller) WebsocketHandler(ctx *context.Context) (interface{}, error) {
	conn, err := ctr.upGrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		return nil, errno.ErrUpGrade.WithErr(err)
	}

	ctr.app.Websocket(conn)

	ctx.StopExecution()

	return nil, nil
}

// JwtDemoHandler
// @Summary JwtDemo
// @Description this is JwtDemo
// @Tags iris
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.JwtDemoResp
// @Router /jwt/demo [get]
func (ctr *Controller) JwtDemoHandler(ctx *context.Context) (interface{}, error) {
	return ctr.app.JwtDemo(ctx)
}

// LogoutHandler
// @Summary Logout
// @Description this is Logout
// @Tags iris
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Router /logout [post]
func (ctr *Controller) LogoutHandler(ctx *context.Context) (interface{}, error) {
	return nil, ctr.app.Logout(ctx)
}
