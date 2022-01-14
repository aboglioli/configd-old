package controllers

import (
	"context"
	"net/http"

	"github.com/aboglioli/configd/application"
	"github.com/aboglioli/configd/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func DeleteSchema(c *gin.Context) {
	deps := dependencies.Get()

	serv := application.NewDeleteSchema(deps.SchemaRepository)

	cmd := application.DeleteSchemaCommand{
		Id: c.Param("schema_id"),
	}

	res, err := serv.Exec(context.Background(), &cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}
