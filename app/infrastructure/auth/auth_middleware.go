package auth

import (
	"encoding/json"
	"go_clean_arch_test/app/domain"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("accessToken")
		session := sessions.Default(ctx)
		// Json文字列がinterdace型で格納されている。dproxyのライブラリを使用して値を取り出す
		loginUserJson, err := dproxy.New(session.Get(accessToken)).String()

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		} else {
			var loginInfo domain.User
			// Json文字列のアンマーシャル
			err := json.Unmarshal([]byte(loginUserJson), &loginInfo)
			if err != nil {
				ctx.Status(http.StatusUnauthorized)
				ctx.Abort()
			} else {
				ctx.Next()
			}
		}
	}
}
