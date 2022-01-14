package main

import (
	"github.com/aboglioli/configd/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	// Schema
	s.POST("/schema", controllers.CreateSchema)
	s.PUT("/schema/:schema_id", controllers.UpdateSchema)
	s.DELETE("/schema/:schema_id", controllers.DeleteSchema)

	// Config
	s.POST("/schema/:schema_id/config", controllers.CreateConfig)
	s.PUT("/config/:config_id", controllers.UpdateConfig)
	s.DELETE("/config/:config_id", controllers.DeleteConfig)

	s.Run(":8080")
}
