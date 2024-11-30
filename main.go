package main

import (
	"github.com/brentellingson/entra-playground/internal/api"
	_ "github.com/brentellingson/entra-playground/internal/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	server := api.NewServer()
	server.Register(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) { c.Redirect(302, "/swagger/index.html") })
	_ = r.Run(":8080")
}
