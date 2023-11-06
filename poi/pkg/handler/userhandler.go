package handler

import (
	"net/http"
	"restapi/meet"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
	var input meet.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) Login(ctx *gin.Context) {

}
