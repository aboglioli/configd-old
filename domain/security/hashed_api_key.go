package security

import (
	"fmt"
)

type HashedApiKey struct {
	hashedApiKey string
}

func NewHashedApiKey(hashedApiKey string) (HashedApiKey, error) {
	if len(hashedApiKey) != 64 {
		return HashedApiKey{}, fmt.Errorf("hashed api key %s is not SHA256", hashedApiKey)
	}

	return HashedApiKey{
		hashedApiKey: hashedApiKey,
	}, nil
}

func (a HashedApiKey) Value() string {
	return a.hashedApiKey
}

func (a HashedApiKey) Equals(o HashedApiKey) bool {
	return a.hashedApiKey == o.hashedApiKey
}

func (a HashedApiKey) Validate(apiKey ApiKey) bool {
	hash, err := apiKey.Hash()
	if err != nil {
		return false
	}

	return a.Equals(hash)
}
