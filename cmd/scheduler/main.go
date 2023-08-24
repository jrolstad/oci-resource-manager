package main

import (
	"github.com/jrolstad/oci-resource-manager/internal/clients"
	"github.com/jrolstad/oci-resource-manager/internal/core"
	"github.com/jrolstad/oci-resource-manager/internal/models"
	"github.com/jrolstad/oci-resource-manager/internal/orchestration"
	"github.com/oracle/oci-go-sdk/common"
	"os"
	"strconv"
	"time"
)

func main() {

	appConfig := getAppConfig()

	configProvider := common.DefaultConfigProvider()
	resourceScheduleClient, err := clients.NewResourceScheduleClient(configProvider, appConfig)
	core.ThrowIfError(err)
	resourceQueryClient, err := clients.NewResourceQueryClient(configProvider)
	core.ThrowIfError(err)
	regionClient, err := clients.NewRegionClient(configProvider)
	core.ThrowIfError(err)

	instance, err := orchestration.NewResourceActionSchedule(configProvider)
	core.ThrowIfError(err)

	// Run Indefinitely, refreshing when specified
	for {
		err = instance.Configure(resourceScheduleClient, resourceQueryClient, regionClient)
		core.ThrowIfError(err)

		instance.Start()
		time.Sleep(time.Duration(appConfig.RefreshIntervalMinutes) * time.Minute)
		instance.Stop()
	}
}

func getAppConfig() *models.AppConfig {
	return &models.AppConfig{
		ResourceScheduleCompartmentId: os.Getenv("resource_manager_compartmentid"),
		ResourceScheduleTableName:     os.Getenv("resource_manager_schedule_table"),
		RefreshIntervalMinutes:        getRefreshInterval(),
	}
}

func getRefreshInterval() int {
	interval, err := strconv.Atoi(os.Getenv("resource_manager_refresh_interval"))
	if err != nil {
		interval = 25
	}

	return interval
}
