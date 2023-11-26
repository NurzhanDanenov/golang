package handler

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"restapi/internal/entity"
)

func (h *Handler) createAva(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var image entity.Image

	filename, ok := ctx.Get("filePath")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filename not found"})
	}

	file, ok := ctx.Get("file")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "couldnt find file in request"})
		return
	}

	// upload file
	imageId, err := h.services.UploadImage.Upload(userId, image, file.(multipart.File), filename.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": imageId,
	})
	//userId, _ := strconv.Atoi(id)
	//update := map[string]string{
	//	"image_url": imageUrl,
	//}
	//updatedUser := models.UpdateUser(userId, update)
	//ctx.JSON(http.StatusOK, gin.H{"data": updatedUser})
	//return

	//var input entity.Image
	//if err := ctx.BindJSON(&input); err != nil {
	//	newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	//	return
	//}

}
