package controllers

import (
	"context"
	"net/http"

	"github.com/aboglioli/configd/application"
	"github.com/aboglioli/configd/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func UpdateSchema(c *gin.Context) {
	deps := dependencies.Get()

	serv := application.NewUpdateSchema(deps.SchemaRepository)

	var cmd application.UpdateSchemaCommand
	if err := c.BindJSON(&cmd); err != nil {
		return
	}

	cmd.Id = c.Param("schema_id")

	res, err := serv.Exec(context.Background(), &cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}
