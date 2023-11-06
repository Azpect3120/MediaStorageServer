package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sends the servers current status back to the user
func Status (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{ "Status": "Service is active!" })
}
