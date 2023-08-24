package clients

import (
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/nosql"
)

type ResourceScheduleClient interface {
	GetAll() ([]*models.ResourceSchedule, error)
}

func NewResourceScheduleClient(config common.ConfigurationProvider, appConfig *models.AppConfig) (ResourceScheduleClient, error) {
	client, err := nosql.NewNosqlClientWithConfigurationProvider(config)
	return &OracleResourceScheduleClient{
		appConfig: appConfig,
		config:    config,
		client:    client,
	}, err
}
