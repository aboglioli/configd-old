package controllers

import (
	"context"
	"net/http"

	"github.com/aboglioli/configd/application"
	"github.com/aboglioli/configd/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func CreateConfig(c *gin.Context) {
	deps := dependencies.Get()

	serv := application.NewCreateConfig(deps.SchemaRepository, deps.ConfigRepository)

	var cmd application.CreateConfigCommand
	if err := c.BindJSON(&cmd); err != nil {
		return
	}

	cmd.SchemaId = c.Param("schema_id")

	res, err := serv.Exec(context.Background(), &cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}
