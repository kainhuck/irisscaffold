package errno

var (
	CommonErrType = ErrorType(10000)
	GreetErrType  = ErrorType(20000)
)

var OK = NewError(200, "ok")

var (
	ErrParameter    = CommonErrType.NewError("请求参数错误")
	ErrUnauthorized = CommonErrType.NewError("访问未授权")
	ErrNotFound     = CommonErrType.NewError("资源未找到")
	ErrInternal     = CommonErrType.NewError("服务器内部错误")
)

var (
	ErrLoginFailed  = GreetErrType.NewError("登录失败")
	ErrLogoutFailed = GreetErrType.NewError("登出失败")
	ErrUpGrade      = GreetErrType.NewError("协议升级失败")
)
