package router

import (
	"goland-crud-gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagController *controllers.TagController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello server")
	})

	baseRoute := router.Group("/api")

	tagRouter := baseRoute.Group("/tag")
	tagRouter.POST("", tagController.Create)
	tagRouter.GET("", tagController.FindMany)
	tagRouter.GET("/:id", tagController.FindById)

	return router
}
