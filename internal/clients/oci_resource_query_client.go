package clients

import (
	"context"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/resourcesearch"
)

type OciResourceQueryClient struct {
	config common.ConfigurationProvider
	client resourcesearch.ResourceSearchClient
}

func (t *OciResourceQueryClient) Query(region string, query string) ([]*models.Resource, error) {
	//TODO: Implement paging and multiple regions
	searchRequest := resourcesearch.SearchResourcesRequest{
		SearchDetails: &resourcesearch.StructuredSearchDetails{
			Query: common.String(query),
		},
	}

	searchResponse, err := t.client.SearchResources(context.Background(), searchRequest)
	if err != nil {
		return make([]*models.Resource, 0), err
	}

	results := make([]*models.Resource, len(searchResponse.Items))
	for index, item := range searchResponse.Items {
		results[index] = &models.Resource{
			Id:                 t.readSearchValue(item.Identifier),
			Region:             region,
			AvailabilityDomain: t.readSearchValue(item.AvailabilityDomain),
			Type:               t.readSearchValue(item.ResourceType),
			State:              t.readSearchValue(item.LifecycleState),
		}

	}

	return results, nil
}

func (t *OciResourceQueryClient) readSearchValue(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
