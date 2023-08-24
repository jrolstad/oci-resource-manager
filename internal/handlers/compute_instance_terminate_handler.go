package handlers

import (
	"context"
	"github.com/jrolstad/oci-resource-manager/internal/logging"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

type ComputeInstanceTerminateHandler struct {
	Config common.ConfigurationProvider
}

func (t *ComputeInstanceTerminateHandler) ResourceType() string {
	return "instance"
}

func (t *ComputeInstanceTerminateHandler) Action() string {
	return "terminate"
}

func (t *ComputeInstanceTerminateHandler) Execute(resource *models.Resource) error {
	client, err := core.NewComputeClientWithConfigurationProvider(t.Config)
	if err != nil {
		return err
	}

	request := core.TerminateInstanceRequest{
		InstanceId: common.String(resource.Id),
	}
	_, err = client.TerminateInstance(context.Background(), request)
	logging.LogInfo("Terminated Instance Pool", "id", resource.Id)
	return err
}
