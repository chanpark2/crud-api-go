package controller

import (
	"Users/witch/IdeaProjects/frameworktest/data/request"
	"Users/witch/IdeaProjects/frameworktest/data/response"
	"Users/witch/IdeaProjects/frameworktest/helper"
	"Users/witch/IdeaProjects/frameworktest/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{tagsService: service}
}

// Create CreateTags godoc
// @Summary Create new tags
// @Description Create new tags
// @Tags tags
// @Accept  application/json
// @Produce  application/json
// @Param tags body request.CreateTagsRequest true "Create Tags"
// @Success 200 {object} response.Response{}
// @Router /tags [post]
func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// Update UpdateTags godoc
// @Summary Update tags
// @Description Update tags
// @Param tagId path int true "Tag ID"
// @Param tags body request.UpdateTagsRequest true "Update Tags"
// @Tags tags
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} response.Response{}
// @Router /tags/{tagId} [put]
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete DeleteTags godoc
// @Summary Delete tags
// @Description Delete tags
// @Param tagId path int true "Tag ID"
// @Tags tags
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} response.Response{}
// @Router /tags/{tagId} [delete]
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// FindById FindByIdTags godoc
// @Summary FindById tags
// @Description FindById tags
// @Param tagId path int true "Tag ID"
// @Tags tags
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} response.Response{}
// @Router /tags/{tagId} [get]
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	result := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAll FindAllTags godoc
// @Summary FindAll tags
// @Description FindAll tags
// @Tags tags
// @Accept  application/json
// @Produce  application/json
// @Success 200 {object} response.Response{}
// @Router /tags [get]
func (controller *TagsController) FindAll(ctx *gin.Context) {
	result := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
