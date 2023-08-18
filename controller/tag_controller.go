package controller

import (
	"example/web-service-gin/data/request"
	"example/web-service-gin/data/response"
	"example/web-service-gin/helper"
	"example/web-service-gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagsController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

//Create controller
func (controller *TagsController) Create(c *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}

	err := c.ShouldBindJSON(&createTagsRequest)

	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusCreated,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, webResponse)
}

//Update controller
func (controller *TagsController) Update(c *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := c.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

//Delete controller
func (controller *TagsController) Delete(c *gin.Context) {
	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

//FindById controller
func (controller *TagsController) FindById(c *gin.Context) {
	tagId := c.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

//FindAll controller
func (controller *TagsController) FindAll(c *gin.Context) {
	tagResponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tagResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
