package main

import (
	"goland-crud-gin/configs"
	"goland-crud-gin/controllers"
	"goland-crud-gin/helper"
	"goland-crud-gin/model"
	"goland-crud-gin/repository"
	"goland-crud-gin/router"
	"goland-crud-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	const port = "8000"

	log.Info().Msg(`Server runnig on port: ` + port)

	db := configs.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.TagModel{})

	tagRepository := repository.NewTagRepositoryImpl(db)

	tagService := services.NewTagServiceImpl(tagRepository, validate)

	tagController := controllers.NewTagController(tagService)

	routes := router.NewRouter(tagController)

	routes.GET("api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, server!")
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)
}
