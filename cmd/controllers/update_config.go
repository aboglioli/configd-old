package controllers

import (
	"context"
	"net/http"

	"github.com/aboglioli/configd/application"
	"github.com/aboglioli/configd/cmd/dependencies"
	"github.com/gin-gonic/gin"
)

func UpdateConfig(c *gin.Context) {
	deps := dependencies.Get()

	serv := application.NewUpdateConfig(deps.SchemaRepository, deps.ConfigRepository)

	var cmd application.UpdateConfigCommand
	if err := c.BindJSON(&cmd); err != nil {
		return
	}

	cmd.Id = c.Param("config_id")

	res, err := serv.Exec(context.Background(), &cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &res)
}
