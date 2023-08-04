package main

import (
	"Users/witch/IdeaProjects/frameworktest/config"
	"Users/witch/IdeaProjects/frameworktest/controller"
	"Users/witch/IdeaProjects/frameworktest/helper"
	"Users/witch/IdeaProjects/frameworktest/model"
	"Users/witch/IdeaProjects/frameworktest/repository"
	"Users/witch/IdeaProjects/frameworktest/router"
	"Users/witch/IdeaProjects/frameworktest/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {

	log.Info().Msg("Starting server...")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagRepository := repository.NewTagsRepositoryImpl(db)

	tagService := service.NewTagsServiceImpl(tagRepository, validate)

	tagsController := controller.NewTagsController(tagService)

	routes := router.NewRouter(tagsController)

	routes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
