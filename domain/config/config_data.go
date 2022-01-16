package config

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type ConfigData map[string]interface{}

func (cd ConfigData) Hash() string {
	b, err := json.Marshal(cd)
	if err != nil {
		return ""
	}

	hash := sha256.Sum256(b)
	return hex.EncodeToString(hash[:])
}
