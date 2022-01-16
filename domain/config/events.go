package config

import (
	"github.com/aboglioli/configd/pkg/events"
)

var (
	ConfigCreatedTopic       = events.NewTopic("config", "created")
	ConfigNameChangedTopic   = events.NewTopic("config", "name_changed")
	ConfigConfigChangedTopic = events.NewTopic("config", "config_changed")
)

type ConfigCreated struct {
	Id        string                 `json:"id"`
	SchemaId  string                 `json:"schema_id"`
	Name      string                 `json:"name"`
	Config    map[string]interface{} `json:"config"`
	ConfigSum string                 `json:"config_sum"`
}

type ConfigNameChanged struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ConfigConfigChanged struct {
	Id        string                 `json:"id"`
	Config    map[string]interface{} `json:"config"`
	ConfigSum string                 `json:"config_sum"`
}
