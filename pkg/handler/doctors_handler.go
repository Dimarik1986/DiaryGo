package handler

import (
	"DiaryGo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateDoctors(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	var input models.Doctor
	if err = ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	updateDoctors := h.services.UpdateDoctors(&input, id)

	ctx.JSON(http.StatusOK, updateDoctors)
}

func (h *Handler) DeleteDoctor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
	}
	delDoctors := h.services.DeleteDoctors(id)

	ctx.JSON(http.StatusOK, delDoctors)
}

func (h *Handler) GetAllDoctors(ctx *gin.Context) {
	getAllDoctors := h.services.GetAllDoctors()
	ctx.JSON(http.StatusOK, map[string]interface{}{"getAllDoctors": getAllDoctors})
}
