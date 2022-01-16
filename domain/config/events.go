package config

import (
	"github.com/aboglioli/configd/pkg/events"
)

var (
	CreatedTopic       = events.NewTopic("config", "created")
	NameChangedTopic   = events.NewTopic("config", "name_changed")
	ConfigChangedTopic = events.NewTopic("config", "config_changed")
)

type Created struct {
	Id        string                 `json:"id"`
	SchemaId  string                 `json:"schema_id"`
	Name      string                 `json:"name"`
	Config    map[string]interface{} `json:"config"`
	ConfigSum string                 `json:"config_sum"`
}

type NameChanged struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ConfigChanged struct {
	Id        string                 `json:"id"`
	Config    map[string]interface{} `json:"config"`
	ConfigSum string                 `json:"config_sum"`
}
