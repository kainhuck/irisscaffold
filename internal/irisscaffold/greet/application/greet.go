package application

import (
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/webmodel/request"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
)

func Greet(req request.GreetReq) (code int, data interface{}, err error) {
	return e.Success, response.GreetResp{Name: req.Name}, nil
}
