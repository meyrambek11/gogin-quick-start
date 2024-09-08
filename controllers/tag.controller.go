package controllers

import (
	"goland-crud-gin/data/requests"
	"goland-crud-gin/data/responses"
	"goland-crud-gin/helper"
	"goland-crud-gin/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService services.TagService
}

func NewTagController(service services.TagService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := requests.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)
	controller.tagService.Create(createTagRequest)

	webResponse := responses.Reponse{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

func (controller *TagController) FindMany(ctx *gin.Context) {
	tagReponse := controller.tagService.FindMany()

	webResponse := responses.Reponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagReponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("id")

	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagReponse := controller.tagService.FindOneById(id)

	webResponse := responses.Reponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagReponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// func (controller *TagController) Delete(ctx *gin.Context) {

// }

// func (controller *TagController) Update(ctx *gin.Context) {

// }
