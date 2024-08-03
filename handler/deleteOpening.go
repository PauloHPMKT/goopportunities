package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DelelteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}