package clients

import (
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/resourcesearch"
)

type ResourceQueryClient interface {
	Query(region string, query string) ([]*models.Resource, error)
}

func NewResourceQueryClient(config common.ConfigurationProvider) (ResourceQueryClient, error) {
	client, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(config)
	return &OciResourceQueryClient{
		config: config,
		client: client,
	}, err
}
