package main

import (
	"github.com/aboglioli/configd/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	// Schema
	s.GET("/schema/:schema_id", controllers.GetSchema)
	s.POST("/schema", controllers.CreateSchema)
	s.PUT("/schema/:schema_id", controllers.UpdateSchema)
	s.DELETE("/schema/:schema_id", controllers.DeleteSchema)

	// Config
	s.GET("/config/:config_id", controllers.GetConfig)
	s.POST("/config", controllers.CreateConfig)
	s.PUT("/config/:config_id", controllers.UpdateConfig)
	s.DELETE("/config/:config_id", controllers.DeleteConfig)

	// User
	s.POST("/login", controllers.LoginUser)
	s.POST("/user", controllers.RegisterUser)

	s.Run(":8080")
}
