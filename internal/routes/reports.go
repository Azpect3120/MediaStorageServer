package routes

import (
	"fmt"
	"net/http"

	"github.com/Azpect3120/MediaStorageServer/internal/database"
	"github.com/Azpect3120/MediaStorageServer/internal/reports"
	"github.com/gin-gonic/gin"
)

// Sends a folder report to the user
func SendReport(db *database.Database, ctx *gin.Context) {
	id := ctx.Param("id")
	email := ctx.Param("email")

	if id == "" || email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Please provided required parameters."})
		return
	}

	if valid := ValidateID(id); !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Please enter a valid id." })
		return
	}

	if valid := ValidateEmail(email); !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{ "status": http.StatusBadRequest, "error": "Please enter a valid email." })
		return
	}


	report, err := reports.Generate(db, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}


	emailContent, err := reports.String(report)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	err = reports.SendEmail(email, "Generated Report", emailContent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": fmt.Sprintf("Report was generated and send to %s", email) })
}
