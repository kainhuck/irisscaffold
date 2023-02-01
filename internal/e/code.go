package e

const (
	Success         = 200
	ErrParameter    = 400
	ErrUnauthorized = 401
	ErrNotFound     = 404
	ErrInternal     = 500
)

const (
	ErrLoginFailed = 10000 + iota
	ErrLogoutFailed
)
