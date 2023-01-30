package context

import (
	"github.com/kainhuck/irisscaffold/internal/e"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
	"github.com/kataras/iris/v12"
	"net/http"
	"sync"
)

type Context struct {
	iris.Context
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

func (c *Context) SendResponse(code int, data interface{}, err error) {
	c.StatusCode(http.StatusOK)

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	resp := &response.ApiResponse{
		Code:      code,
		Message:   e.GetMsg(code),
		Error:     errStr,
		Data:      data,
		RequestID: c.GetID().(string),
	}

	_ = c.JSON(resp)
}

func (c *Context) SendNoBodyResponse(code int, err error) {
	c.SendResponse(code, nil, err)
}

func Handler(h func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		if !ctx.IsStopped() { // 请求被终止
			h(ctx)
		}
		release(ctx)
	}
}
