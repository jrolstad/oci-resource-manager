package clients

import (
	"context"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/identity"
)

type OciRegionClient struct {
	config common.ConfigurationProvider
	client identity.IdentityClient
}

func (t *OciRegionClient) GetAvailableRegions() ([]*models.Region, error) {
	regionResponse, err := t.client.ListRegions(context.Background())
	if err != nil {
		return make([]*models.Region, 0), err
	}

	regions := make([]*models.Region, len(regionResponse.Items))
	for index, region := range regionResponse.Items {
		regions[index] = &models.Region{Name: *region.Key}
	}

	return regions, nil
}

func (t *OciRegionClient) GetDefaultRegion() (*models.Region, error) {

	tenancyId, err := t.config.TenancyOCID()
	if err != nil {
		return nil, err
	}

	tenancyRequest := identity.GetTenancyRequest{TenancyId: common.String(tenancyId)}
	tenancyResponse, err := t.client.GetTenancy(context.Background(), tenancyRequest)
	if err != nil {
		return nil, err
	}

	return &models.Region{Name: *tenancyResponse.HomeRegionKey}, nil
}
