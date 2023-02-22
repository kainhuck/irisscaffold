package application

import (
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/middleware"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/sirupsen/logrus"
	"time"
)

func (app *Application) Greet(req request.GreetReq) (code int, data interface{}, err error) {
	return e.Success, response.GreetResp{Name: req.Name}, nil
}

func (app *Application) Login(req request.LoginReq) (code int, data interface{}, err error) {
	user, err := app.dbClient.GetUserByName(req.Username)
	if err != nil {
		logrus.Errorf("login failed: %v", err)
		return e.ErrLoginFailed, nil, err
	}

	if user.Password != req.Password {
		return e.ErrLoginFailed, nil, nil
	}

	signer := jwt.NewSigner(jwt.HS256, app.cfg.Jwt.SigKey, time.Duration(app.cfg.Jwt.ExpireTime)*time.Second)
	// 自行修改
	token, err := signer.Sign(middleware.Claims{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		logrus.Errorf("login failed: %v", err)
		return e.ErrLoginFailed, nil, err
	}

	return e.Success, response.LoginResp{Token: string(token)}, nil
}

func (app *Application) JwtDemo(ctx *context.Context) (code int, data interface{}, err error) {
	claims := jwt.Get(ctx.Context).(*middleware.Claims)

	return e.Success, response.JwtDemoResp{
		Username: claims.Username,
		Password: claims.Password,
	}, nil
}

func (app *Application) Logout(ctx *context.Context) (code int, err error) {
	if err = ctx.Logout(); err != nil {
		logrus.Errorf("logout failed: %v", err)
		return e.ErrLogoutFailed, err
	}

	return e.Success, nil
}
