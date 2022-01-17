package security

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var apiKeyCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// API Key
type ApiKey struct {
	apiKey string
}

func NewApiKey(apiKey string) (ApiKey, error) {
	if len(apiKey) < 10 {
		return ApiKey{}, fmt.Errorf("api key %s too short", apiKey)
	}

	return ApiKey{
		apiKey: apiKey,
	}, nil
}

func GenerateApiKey() (ApiKey, error) {
	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	keys := make([]rune, 36)
	for i := range keys {
		keys[i] = apiKeyCharacters[rand.Intn(len(apiKeyCharacters))]
	}

	return NewApiKey(string(keys))
}

func (a ApiKey) Value() string {
	return a.apiKey
}

func (a ApiKey) Equals(o ApiKey) bool {
	return a.apiKey == o.apiKey
}

func (a ApiKey) Hash() (HashedApiKey, error) {
	hashedApiKey := sha256.Sum256([]byte(a.apiKey))

	return NewHashedApiKey(hex.EncodeToString(hashedApiKey[:]))
}
