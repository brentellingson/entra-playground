package main

import (
	"github.com/brentellingson/entra-playground/internal/api"
	"github.com/brentellingson/entra-playground/internal/config"
	_ "github.com/brentellingson/entra-playground/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg := config.NewConfig()

	r := gin.Default()
	server := api.NewServer(cfg)
	server.Register(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) { c.Redirect(302, "/swagger/index.html") })
	_ = r.Run(":8080")
}
