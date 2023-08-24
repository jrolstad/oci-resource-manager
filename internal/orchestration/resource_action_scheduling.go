package orchestration

import (
	"errors"
	"github.com/go-co-op/gocron"
	"github.com/jrolstad/oci-resource-manager/internal/clients"
	"github.com/jrolstad/oci-resource-manager/internal/logging"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/oracle/oci-go-sdk/common"
	"time"
)

type ResourceActionScheduler struct {
	Scheduler *gocron.Scheduler
	Handler   *ResourceActionOrchestrator
}

func NewResourceActionSchedule(config common.ConfigurationProvider) (ResourceActionScheduler, error) {
	return ResourceActionScheduler{
		Scheduler: gocron.NewScheduler(time.UTC),
		Handler:   NewResourceActionHandler(config),
	}, nil
}

func (t *ResourceActionScheduler) Configure(scheduleClient clients.ResourceScheduleClient,
	resourceQueryClient clients.ResourceQueryClient,
	regionClient clients.RegionClient) error {

	logging.LogInfo("Starting Resource Action Schedule Process")

	schedules, err := scheduleClient.GetAll()
	if err != nil {
		return err
	}
	logging.LogInfo("Obtained Resource Schedules", "count", len(schedules))

	schedulingErrors := make([]error, 0)
	for _, schedule := range schedules {
		scheduleErr := t.createSchedule(schedule, t.Scheduler, resourceQueryClient, regionClient)
		if scheduleErr != nil {
			schedulingErrors = append(schedulingErrors, scheduleErr)
		}
		logging.LogInfo("Created Schedule",
			"name", schedule.Name,
			"schedule", schedule.Schedule)
	}

	logging.LogInfo("Configured Resource Action Schedule Processes")
	return errors.Join(schedulingErrors...)
}

func (t *ResourceActionScheduler) Start() {
	t.Scheduler.StartAsync()
	logging.LogInfo("Started Resource Action Schedule Processes", "scheduled_processes", t.Scheduler.Len())
}

func (t *ResourceActionScheduler) Stop() {
	t.Scheduler.Stop()
	t.Scheduler.Clear()
	logging.LogInfo("Stopped Resource Action Schedule Processes", "scheduled_processes", t.Scheduler.Len())
}

func (t *ResourceActionScheduler) createSchedule(schedule *models.ResourceSchedule,
	scheduler *gocron.Scheduler,
	resourceQueryClient clients.ResourceQueryClient,
	regionClient clients.RegionClient) error {
	_, err := scheduler.Cron(schedule.Schedule).Do(t.runSchedule, schedule, resourceQueryClient, regionClient)
	return err
}

func (t *ResourceActionScheduler) runSchedule(schedule *models.ResourceSchedule,
	resourceQueryClient clients.ResourceQueryClient,
	regionClient clients.RegionClient) error {

	logging.LogInfo("Executing Resource Schedule", "name", schedule.Name)
	regions, err := t.resolveRegions(schedule.Region, regionClient)
	if err != nil {
		return err
	}

	processingErrors := make([]error, 0)
	for _, region := range regions {
		resources, queryErr := resourceQueryClient.Query(region.Name, schedule.Resources)
		if queryErr != nil {
			processingErrors = append(processingErrors, queryErr)
		}

		for _, resource := range resources {
			if t.Handler.IsSupported(resource, schedule.Action) {
				resourceErr := t.Handler.Process(resource, schedule.Action)
				if resourceErr != nil {
					processingErrors = append(processingErrors, resourceErr)
				}
			} else {
				logging.LogInfo("Unsupported resource action",
					"id", resource.Id,
					"type", resource.Type,
					"action", schedule.Action,
				)
			}
		}
	}

	return errors.Join(processingErrors...)
}

func (t *ResourceActionScheduler) resolveRegions(regions []string, regionClient clients.RegionClient) ([]*models.Region, error) {
	if len(regions) == 1 && regions[0] == "*" {
		regions, err := regionClient.GetAvailableRegions()
		return regions, err
	}
	if len(regions) == 1 && regions[0] == "" {
		defaultRegion, err := regionClient.GetDefaultRegion()
		return []*models.Region{defaultRegion}, err
	}

	resolvedRegions := make([]*models.Region, len(regions))
	for index, region := range regions {
		resolvedRegions[index] = &models.Region{Name: region}
	}
	return resolvedRegions, nil
}
