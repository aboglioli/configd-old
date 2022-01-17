package security

import (
	"github.com/aboglioli/configd/pkg/models"
)

type Access string

const (
	READ_ONLY_ACCESS Access = "read_only"
	FULL_ACCESS      Access = "full_access"
)

type Authorization struct {
	hashedApiKey HashedApiKey
	resourceId   models.Id
	access       Access
}

func BuildAuthorization(
	hashedApiKey HashedApiKey,
	resourceId models.Id,
	access Access,
) (*Authorization, error) {
	return &Authorization{
		hashedApiKey: hashedApiKey,
		resourceId:   resourceId,
		access:       access,
	}, nil
}

func NewAuthorization(
	apiKey ApiKey,
	resourceId models.Id,
	access Access,
) (*Authorization, error) {
	hash, err := apiKey.Hash()
	if err != nil {
		return nil, err
	}

	return BuildAuthorization(hash, resourceId, access)
}

func (a *Authorization) HashedApiKey() HashedApiKey {
	return a.hashedApiKey
}

func (a *Authorization) ResourceId() models.Id {
	return a.resourceId
}

func (a *Authorization) Access() Access {
	return a.access
}
