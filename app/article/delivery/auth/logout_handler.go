package delivery

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("accessToken")
	session := sessions.Default(ctx)
	session.Delete(accessToken)
	session.Save()
}
