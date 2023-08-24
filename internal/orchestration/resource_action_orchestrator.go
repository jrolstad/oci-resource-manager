package orchestration

import (
	"fmt"
	"github.com/jrolstad/oci-resource-manager/internal/handlers"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"strings"
)

type ResourceActionOrchestrator struct {
	config   common.ConfigurationProvider
	handlers map[string]func(resource *models.Resource) error
}

func NewResourceActionHandler(config common.ConfigurationProvider) *ResourceActionOrchestrator {
	instance := &ResourceActionOrchestrator{
		config:   config,
		handlers: make(map[string]func(resource *models.Resource) error),
	}
	instance.registerHandlers()

	return instance
}

func (t *ResourceActionOrchestrator) IsSupported(resource *models.Resource, action string) bool {
	return t.handlers[t.getHandlerKey(resource.Type, action)] != nil
}

func (t *ResourceActionOrchestrator) Process(resource *models.Resource, action string) error {
	handler := t.handlers[t.getHandlerKey(resource.Type, action)]
	if handler == nil {
		return nil
	}

	return handler(resource)
}

func (t *ResourceActionOrchestrator) getHandlerKey(resourceType string, action string) string {
	return fmt.Sprintf("%s|%s", strings.ToLower(resourceType), strings.ToLower(action))
}

func (t *ResourceActionOrchestrator) registerHandlers() {
	t.registerResources(&handlers.ComputeInstanceStopHandler{Config: t.config})
	t.registerResources(&handlers.ComputeInstanceTerminateHandler{Config: t.config})
	t.registerResources(&handlers.NoSqlTableNoOpHandler{Config: t.config})
}

func (t *ResourceActionOrchestrator) registerResources(resource handlers.ResourceActionHandler) {
	t.handlers[t.getHandlerKey(resource.ResourceType(), resource.Action())] = resource.Execute
}
