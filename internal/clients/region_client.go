package clients

import (
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
)

type RegionClient interface {
	GetAvailableRegions() ([]*models.Region, error)
	GetDefaultRegion() (*models.Region, error)
}

func NewRegionClient(config common.ConfigurationProvider) (RegionClient, error) {
	client, err := identity.NewIdentityClientWithConfigurationProvider(config)
	return &OciRegionClient{
		config: config,
		client: client,
	}, err
}
