package handlers

import (
	"context"
	"github.com/jrolstad/oci-resource-manager/internal/logging"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/core"
)

type ComputeInstanceStopHandler struct {
	Config common.ConfigurationProvider
}

func (t *ComputeInstanceStopHandler) ResourceType() string {
	return "instance"
}

func (t *ComputeInstanceStopHandler) Action() string {
	return "stop"
}

func (t *ComputeInstanceStopHandler) Execute(resource *models.Resource) error {
	client, err := core.NewComputeClientWithConfigurationProvider(t.Config)
	if err != nil {
		return err
	}

	request := core.InstanceActionRequest{
		InstanceId: common.String(resource.Id),
		Action:     core.InstanceActionActionStop,
	}
	response, err := client.InstanceAction(context.Background(), request)
	logging.LogInfo("Stopped Instance Pool", "id", response.Id, "state", response.LifecycleState)
	return err
}
