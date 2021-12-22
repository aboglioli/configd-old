package config

import (
	"github.com/aboglioli/configd/domain/schema"
)

type Config struct {
	schemaName *schema.Name
	name       string
	config     map[string]interface{}
}
