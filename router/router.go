package router

import (
	"Users/witch/IdeaProjects/frameworktest/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api/v1")
	tagsRouter := baseRouter.Group("/tags")

	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PUT("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	return router
}
