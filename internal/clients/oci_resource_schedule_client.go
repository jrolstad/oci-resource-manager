package clients

import (
	"context"
	"fmt"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/nosql"
	"strings"
)

type OracleResourceScheduleClient struct {
	appConfig *models.AppConfig
	config    common.ConfigurationProvider
	client    nosql.NosqlClient
}

func (t *OracleResourceScheduleClient) GetAll() ([]*models.ResourceSchedule, error) {
	//TODO: Implement paging
	queryRequest := nosql.QueryRequest{
		QueryDetails: nosql.QueryDetails{
			CompartmentId: common.String(t.appConfig.ResourceScheduleCompartmentId),
			Statement:     common.String(t.buildGetAllStatement(t.appConfig.ResourceScheduleTableName)),
		},
	}
	data, err := t.client.Query(context.Background(), queryRequest)
	if err != nil {
		return make([]*models.ResourceSchedule, 0), err
	}

	result := make([]*models.ResourceSchedule, len(data.Items))
	for index, item := range data.Items {
		regionRaw := t.getRowValue(item, "region")
		regions := strings.Split(regionRaw, ",")

		result[index] = &models.ResourceSchedule{
			Id:        t.getRowValue(item, "id"),
			Name:      t.getRowValue(item, "name"),
			Schedule:  t.getRowValue(item, "schedule"),
			Action:    t.getRowValue(item, "action"),
			Region:    regions,
			Resources: t.getRowValue(item, "resources"),
		}
	}

	return result, nil
}

func (t *OracleResourceScheduleClient) buildGetAllStatement(tableName string) string {
	return fmt.Sprintf("SELECT * FROM %s", tableName)
}

func (t *OracleResourceScheduleClient) getRowValue(row map[string]interface{}, field string) string {
	value := row[field]
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}
