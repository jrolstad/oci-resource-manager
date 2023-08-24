# Oracle Cloud Infrastructure (OCI) Resource Manager
One of the aspects of the cloud that often surprises customers are unexpected costs of resources, especially around development and testing resources that continue to run after they are no longer needed.  Unlike other cloud providers, OCI does not have fully automated Resource Auto-Stop/Start/Terminate functionality that acts on resources in a tenancy on a scheduled basis that exacerbates this issue.

The OCI Resource Manager solution allows customers of Oracle Cloud Infrastructure to schedule the start, stop, and termination of onboarded resource types so their lifecycle can be managed individually or across an entire tenancy.  When implemented, this automated management of resource state enables reduced consumption of resources when they are not needed, effectively manage state for ephemeral resources such as those used for testing / development, and reduce costs when using the cloud.

## Resource Action Schedule
To support the auto-stop/start/terminate functionality, the OCI Resource Manager consists of two components - a definition of schedules and associated actions, and a background worker process to implement them.

Each Resource Action Schedule instance contains a schedule expression / definition, action to take (start, stop, terminate), query or list of resources to apply the action to when triggered, and regions to apply to.  When a schedule is triggered, performs the action on the resources defined in the resource definition and capture any state changes in the execution history logs.

### Resource Action Schedule Attributes
|Name|Description|
|---|---|
|Schedule|	CRON expression that defines when the schedule is triggered|
|Action|	Action to take on resources when triggered.  One of the following options: Start, Stop, Terminate|
|Resource Definition|Either a list of specific resources or a query that identifies which resources to act on when triggered|
|Regions| The list of regions where resources to act on are located at.  Wildcards are supported where * will act on all regions for the tenancy and null will only act on the home region|

### Examples
|Scenario| Configuration                                                                                                                                                                                                                            |
|---|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|Terminate all resources over 30 days old in a compartment on a daily basis| <ul><li>Schedule: _0 0 0 ? * * *_</li><li>Action: _Terminate_</li><li>Resource Definition: _query all resourceswhere compartmentId = 'compartmentOcid' && timeCreated < 'now -30d'_</li><li>Regions: _*_</li></ul>                       |
|Stop a specific compute resource at 5pm PT every weekday| <ul><li>Schedule: _0 17 0 ? * MON,TUE,WED,THU,FRI *_</li><li>Action: _Stop_</li><li>Resource Definition: _ocid1.compartment.oc1..aaaaaaaalfpzro3dbkzeq7oli7vexwor6zpngfuumchacmnqodbj6fau2gbq_</li><li>Regions: _us-ashburn-1_</li></ul> |

## Implementation
Once the Resource Action Schedule is defined, it needs to be continually evaluated based on its Schedule CRON expression and when triggered, act on the resources defined in the Resource Definition that are not in the desired state.  

### Components

![Component Diagram](/docs/oci-resource-manager.png)

|Name|Purpose|Type|
|---|---|---|
|Resource Action Schedule|Persistent storage for schedule definitions|OCI NoSql Table|
|Scheduler|Background worker that continually evaluates the defined schedules and performs required actions|OCI Container Instance|

### Scheduler Flow
![Sequence Diagram](/docs/sequence-diagram.png)
