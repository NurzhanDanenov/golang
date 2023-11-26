package handler

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type errorst struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statuscode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statuscode, errorst{message})
}
