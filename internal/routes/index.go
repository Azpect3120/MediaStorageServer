package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirects the user to github page
func Index (ctx *gin.Context) {
	ctx.Redirect(http.StatusPermanentRedirect, "https://github.com/Azpect3120/MediaStorageServer")
}
