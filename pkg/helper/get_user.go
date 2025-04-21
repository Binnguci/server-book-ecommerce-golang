package helper

import "github.com/gin-gonic/gin"

func GetUserFromContext(ctx *gin.Context) string {
	user, ok := ctx.Get("user")
	if !ok {
		return ""
	}
	username, ok := user.(string)
	if !ok {
		return ""
	}
	return username
}
