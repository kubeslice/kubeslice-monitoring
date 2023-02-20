package schema

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
)

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
	controllerEvents, err := parseEvent("../../config/events/controller.yaml")
	if err != nil {
		return nil, err
	}
	workerEvents, err := parseEvent("../../config/events/worker.yaml")
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
