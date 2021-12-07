package handlerUsers

import (
	util "test-agit/helpers/utils"
	users "test-agit/internal/controllers/users"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service users.Service
}

func NewHandler(service users.Service) *handler {
	return &handler{service: service}
}

func (h *handler) Store(ctx *gin.Context) {

	var payload users.CreateRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	validate := payload.Validator()
	if validate != "" {
		util.FAILED(ctx, validate)
		return
	}

	data, err := h.service.Create(&payload)
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	util.JSON(ctx, "Berhasil create", data)
	return
}
func (h *handler) Update(ctx *gin.Context) {

	var payload users.UpdateRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	validate := payload.Validator()
	if validate != "" {
		util.FAILED(ctx, validate)
		return
	}

	data, err := h.service.Update(&payload)
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	util.JSON(ctx, "Berhasil update", data)
	return
}
func (h *handler) Delete(ctx *gin.Context) {

	dataFilter := &users.Filter{
		Id: uint(util.StringToInt(ctx.Request.URL.Query().Get("id"))),
	}

	errData := h.service.Destroy(dataFilter)
	if errData != nil {
		util.ERROR(ctx, errData)
		return
	}

	util.JSON(ctx, "Berhasil delete", dataFilter)
	return
}
func (h *handler) Detail(ctx *gin.Context) {

	dataFilter := &users.Filter{
		Id: uint(util.StringToInt(ctx.Request.URL.Query().Get("id"))),
	}

	data, err := h.service.Detail(dataFilter)
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	util.JSON(ctx, "Berhasil", data)
	return
}
func (h *handler) List(ctx *gin.Context) {

	data, err := h.service.List()
	if err != nil {
		util.ERROR(ctx, err)
		return
	}

	util.JSON(ctx, "Berhasil", data)
	return
}
