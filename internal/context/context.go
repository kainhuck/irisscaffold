package context

import (
	"github.com/kainhuck/irisscaffold/internal/errno"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
	"github.com/kataras/iris/v12"
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

type HandlerFunc func(ctx *Context) (interface{}, error)

func Handler(f HandlerFunc) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		if !ctx.IsStopped() { // 请求被终止
			handler(ctx, f)
		}
		release(ctx)
	}
}

func handler(ctx *Context, f HandlerFunc) {
	data, err := f(ctx)

	if ctx.IsStopped() {
		return
	}

	if err != nil {
		makeResp(ctx, err)
	} else {
		makeResp(ctx, data)
	}
}

func makeResp(ctx *Context, data interface{}) {
	rsp := &response.ApiResponse{
		RequestID: ctx.GetID().(string),
	}

	switch d := data.(type) {
	case errno.Error:
		rsp.Code = d.GetBusinessCode()
		rsp.Message = d.GetMsg()
		if err := d.GetErr(); err != nil {
			rsp.Error = err.Error()
		}
		ctx.StatusCode(d.GetHttpCode())
	case error:
		en := errno.ErrInternal.WithErr(d)
		rsp.Code = en.GetBusinessCode()
		rsp.Message = en.GetMsg()
		rsp.Error = d.Error()
		ctx.StatusCode(en.GetHttpCode())
	default:
		rsp.Code = errno.OK.GetBusinessCode()
		rsp.Message = errno.OK.GetMsg()
		rsp.Data = data
	}

	_ = ctx.JSON(rsp)
}
