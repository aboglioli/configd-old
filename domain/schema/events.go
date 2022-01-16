package schema

import (
	"github.com/aboglioli/configd/pkg/events"
)

var (
	SchemaCreatedTopic      = events.NewTopic("schema", "created")
	SchemaNameChangedTopic  = events.NewTopic("schema", "name_changed")
	SchemaPropsChangedTopic = events.NewTopic("schema", "props_changed")
)

type SchemaCreated struct {
	Id    string                 `json:"id"`
	Name  string                 `json:"name"`
	Props map[string]interface{} `json:"props"`
}

type SchemaNameChanged struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SchemaPropsChanged struct {
	Id    string                 `json:"id"`
	Props map[string]interface{} `json:"props"`
}
