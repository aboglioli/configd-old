package security

import (
	"github.com/aboglioli/configd/pkg/models"
)

type Access string

const (
	READ_ONLY   Access = "read_only"
	FULL_ACCESS Access = "full_access"
)

type Authorization struct {
	hashedApiKey HashedApiKey
	configId     models.Id
	access       Access
}

func BuildAuthorization(
	hashedApiKey HashedApiKey,
	configId models.Id,
	access Access,
) (*Authorization, error) {
	return &Authorization{
		hashedApiKey: hashedApiKey,
		configId:     configId,
		access:       access,
	}, nil
}

func NewAuthorization(
	apiKey ApiKey,
	configId models.Id,
	access Access,
) (*Authorization, error) {
	hash, err := apiKey.Hash()
	if err != nil {
		return nil, err
	}

	return BuildAuthorization(hash, configId, access)
}

func (a *Authorization) HashedApiKey() HashedApiKey {
	return a.hashedApiKey
}

func (a *Authorization) ConfigId() models.Id {
	return a.configId
}

func (a *Authorization) Access() Access {
	return a.access
}
