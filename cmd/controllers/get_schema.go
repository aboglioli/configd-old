package controllers

import (
	"context"
	"net/http"

	"github.com/aboglioli/configd/application"
	"github.com/aboglioli/configd/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func GetSchema(c *gin.Context) {
	deps := dependencies.Get()

	serv := application.NewGetSchema(deps.SchemaRepository)

	cmd := application.GetSchemaCommand{
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
