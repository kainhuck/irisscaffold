package errno

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	_ Error = (*err)(nil)
	_ error = (*err)(nil)
)

type Error interface {
	WithErr(err error) Error
	WithHttpCode(int) Error
	WithDetail(formatMsg string, args ...interface{}) Error
	GetBusinessCode() int
	GetHttpCode() int
	GetMsg() string
	GetErr() error
	String() string
	Error() string
}

type err struct {
	HttpCode     int    // HTTP Code
	BusinessCode int    // Business Code
	Message      string // 描述信息
	Detail       string // 错误细节信息
	Err          error  // 错误信息
}

func (e *err) clone() *err {
	return &err{
		HttpCode:     e.HttpCode,
		BusinessCode: e.BusinessCode,
		Message:      e.Message,
		Detail:       e.Detail,
		Err:          e.Err,
	}
}

func NewError(businessCode int, msg string) Error {
	return NewErrorWithCode(http.StatusOK, businessCode, msg)
}

func NewErrorWithCode(httpCode int, businessCode int, msg string) Error {
	return &err{
		HttpCode:     httpCode,
		BusinessCode: businessCode,
		Message:      msg,
	}
}

func (e *err) Error() string {
	return e.String()
}

func (e *err) WithErr(err error) Error {
	e1 := e.clone()
	e1.Err = errors.WithStack(err)
	return e1
}

func (e *err) WithHttpCode(i int) Error {
	e1 := e.clone()
	e1.HttpCode = i
	return e1
}

func (e *err) WithDetail(formatMsg string, args ...interface{}) Error {
	e1 := e.clone()
	e1.Detail = fmt.Sprintf(formatMsg, args...)
	return e1
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	res := e.Message

	if e.Detail != "" {
		return res + "(" + e.Detail + ")"
	}
	return res
}

func (e *err) GetErr() error {
	return e.Err
}

func (e *err) String() string {
	return fmt.Sprintf("[%d]%s - Error: %v", e.BusinessCode, e.Message, e.Err)
}
