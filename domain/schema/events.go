package schema

import (
	"github.com/aboglioli/configd/pkg/events"
)

var (
	CreatedTopic      = events.NewTopic("schema", "created")
	NameChangedTopic  = events.NewTopic("schema", "name_changed")
	PropsChangedTopic = events.NewTopic("schema", "props_changed")
)

type Created struct {
	Id    string                 `json:"id"`
	Name  string                 `json:"name"`
	Props map[string]interface{} `json:"props"`
}

type NameChanged struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PropsChanged struct {
	Id    string                 `json:"id"`
	Props map[string]interface{} `json:"props"`
}
