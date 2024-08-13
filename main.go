package main

import (
	"goland-crud-gin/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	const port = "8000"

	log.Info().Msg(`Server runnig on port: ` + port)

	routes := gin.Default()

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
