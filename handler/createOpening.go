package handler

import (
	"net/http"

	"github.com/PauloHPMKT/goopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	// initialize request struct and bind it to the request body
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request) // populate everyting in the request struct

	if err := request.Validate(); err != nil {
		logger.Errorf("Error validating request: %v", err.Error())
		sendError(ctx, 400, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:    request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote:  *request.Remote,
		Link:    request.Link,
		Salary:  request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening")
		return
	} 
	sendResponse(ctx, "create-opening", opening)
}