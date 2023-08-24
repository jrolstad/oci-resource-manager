package handlers

import (
	"context"
	"github.com/jrolstad/oci-resource-manager/internal/logging"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/nosql"
)

type NoSqlTableNoOpHandler struct {
	Config common.ConfigurationProvider
}

func (t *NoSqlTableNoOpHandler) ResourceType() string {
	return "NoSqlTable"
}

func (t *NoSqlTableNoOpHandler) Action() string {
	return "terminate"
}

func (t *NoSqlTableNoOpHandler) Execute(resource *models.Resource) error {
	client, err := nosql.NewNosqlClientWithConfigurationProvider(t.Config)
	if err != nil {
		return err
	}

	request := nosql.GetTableRequest{
		TableNameOrId: common.String(resource.Id),
	}
	response, err := client.GetTable(context.Background(), request)

	logging.LogInfo("Found NoSql Table", "id", response.Id, "name", response.Name)
	return err
}
