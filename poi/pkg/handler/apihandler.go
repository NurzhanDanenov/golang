package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createAva(ctx *gin.Context) {
	id, _ := ctx.Get(userCtx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
