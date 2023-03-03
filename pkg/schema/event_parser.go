package schema

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"path"
)

type EventConfig struct {
	DisabledEvents []string
}

type EventType string

var (
	EventTypeWarning EventType = "Warning"
	EventTypeNormal  EventType = "Normal"
)

type EventSchemaList struct {
	Events []EventSchema
}

type EventSchema struct {
	Name                string
	Reason              string
	Action              string
	Type                EventType
	ReportingController string
	Message             string
}

func GetEvent(name string) (*EventSchema, error) {
	file1 := "controller.yaml"
	file2 := "worker.yaml"
	controllerFilePath := path.Join("../../config/events", file1)
	workerFilePath := path.Join("../../config/events", file2)
	dir := os.Getenv("EVENT_SCHEMA_PATH")
	if dir != "" {
		controllerFilePath = path.Join(dir, file1)
		workerFilePath = path.Join(dir, file2)
	}
	controllerEvents, err := parseEvent(controllerFilePath)
	if err != nil {
		return nil, err
	}
	workerEvents, err := parseEvent(workerFilePath)
	if err != nil {
		return nil, err
	}
	events := append(controllerEvents, workerEvents...)
	for _, event := range events {
		if event.Name == name {
			return &event, nil
		}
	}
	return nil, nil
}

func IsEventDisabled(name string) bool {
	configs, err := parseConfig("/events/event-schema/controller.yaml")
	if err != nil {
		return false
	}
	for _, config := range configs {
		if config == name {
			return true
		}
	}
	return false
}

func parseEvent(filepath string) ([]EventSchema, error) {
	var eventSchema EventSchemaList
	event, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(event, &eventSchema)
	if err != nil {
		return nil, err
	}
	return eventSchema.Events, nil
}

func parseConfig(filepath string) ([]string, error) {
	var eventConfig EventConfig
	event, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(event, &eventConfig)
	if err != nil {
		return nil, err
	}
	return eventConfig.DisabledEvents, nil
}
