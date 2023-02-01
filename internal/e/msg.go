package e

var MsgFlag = map[int]string{
	Success:         "ok",
	ErrParameter:    "请求参数错误",
	ErrUnauthorized: "访问未授权",
	ErrNotFound:     "资源未找到",
	ErrInternal:     "服务器内部错误",
	ErrLoginFailed:  "登录失败",
	ErrLogoutFailed: "登出失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}

	return MsgFlag[ErrInternal]
}
