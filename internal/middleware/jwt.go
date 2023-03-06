package middleware

import (
	"github.com/kainhuck/irisscaffold/internal/errno"
	"github.com/kainhuck/irisscaffold/internal/webmodel/response"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"net/http"
)

// Claims 自己修改字段，保证唯一即可
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func JwtVerify(sigKey string) context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	// Enable server-side token block feature (even before its expiration time):
	verifier.WithDefaultBlocklist()

	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		if err == nil {
			return
		}

		ctx.StopExecution()
		ctx.StatusCode(http.StatusOK)
		_ = ctx.JSON(response.ApiResponse{
			Code:      errno.ErrParameter.GetBusinessCode(),
			Message:   errno.ErrParameter.GetMsg(),
			Error:     "",
			Data:      nil,
			RequestID: ctx.GetID().(string),
		})
	}
	// Enable payload decryption with:
	// verifier.WithDecryption(encKey, nil)
	return verifier.Verify(func() interface{} {
		return new(Claims)
	})
}
