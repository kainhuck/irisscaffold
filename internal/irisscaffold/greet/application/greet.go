package application

import (
	"github.com/gorilla/websocket"
	"github.com/kainhuck/irisscaffold/internal/context"
	"github.com/kainhuck/irisscaffold/internal/errno"
	"github.com/kainhuck/irisscaffold/internal/middleware"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
)

func (app *Application) Greet(req request.GreetReq) (interface{}, error) {
	app.log.Infof("Greet: %v", "test")
	return &response.GreetResp{Name: req.Name}, nil
}

func (app *Application) Login(req request.LoginReq) (interface{}, error) {
	user, err := app.dbClient.GetUserByName(req.Username)
	if err != nil {
		app.log.Errorf("login failed: %v", err)
		return nil, errno.ErrLoginFailed.WithErr(err)
	}

	if user.Password != req.Password {
		return nil, errno.ErrLoginFailed.WithDetail("密码不正确")
	}

	signer := jwt.NewSigner(jwt.HS256, app.cfg.Jwt.SigKey, time.Duration(app.cfg.Jwt.ExpireTime)*time.Second)
	// 自行修改
	token, err := signer.Sign(middleware.Claims{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		app.log.Errorf("login failed: %v", err)
		return nil, errno.ErrLoginFailed.WithErr(err)
	}

	return &response.LoginResp{Token: string(token)}, nil
}

func (app *Application) Websocket(conn *websocket.Conn) {
	app.log.Debugf("[%v] join", conn.RemoteAddr().String())
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				app.log.Debugf("[%v] leave", conn.RemoteAddr().String())
				return
			default:
				app.log.Errorf("[%v] disconnect, error: %v", conn.RemoteAddr().String(), err)
				return
			}
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte("hello: "+string(message))); err != nil {
			app.log.Errorf("write failed: %v", err)
			return
		}
	}
}

func (app *Application) JwtDemo(ctx *context.Context) (interface{}, error) {
	claims := jwt.Get(ctx.Context).(*middleware.Claims)

	return &response.JwtDemoResp{
		Username: claims.Username,
		Password: claims.Password,
	}, nil
}

func (app *Application) Logout(ctx *context.Context) error {
	if err := ctx.Logout(); err != nil {
		app.log.Errorf("logout failed: %v", err)
		return errno.ErrLogoutFailed.WithErr(err)
	}

	return nil
}
