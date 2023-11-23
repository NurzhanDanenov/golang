package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createAva(ctx *gin.Context) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return
	}

	var input meet.Image
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
}
