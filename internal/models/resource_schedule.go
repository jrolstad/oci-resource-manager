package models

type ResourceSchedule struct {
	Id        string
	Name      string
	Schedule  string
	Action    string
	Region    []string
	Resources string
}
