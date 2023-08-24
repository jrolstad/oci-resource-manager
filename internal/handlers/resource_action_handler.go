package handlers

import "github.com/jrolstad/oci-resource-manager/internal/models"

type ResourceActionHandler interface {
	ResourceType() string
	Action() string
	Execute(resource *models.Resource) error
}
