package e

var MsgFlag = map[int]string{
	Success:      "ok",
	ErrParameter: "请求参数错误",
	ErrNotFound:  "资源未找到",
	ErrInternal:  "服务器内部错误",
	ErrOther:     "其他错误类型",
}

func GetMsg(code int) string {
	msg, ok := MsgFlag[code]
	if ok {
		return msg
	}

	return MsgFlag[ErrInternal]
}
