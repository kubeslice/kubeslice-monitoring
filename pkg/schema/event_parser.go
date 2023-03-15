package schema

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/util/yaml"
)

type EventConfig struct {
	DisabledEvents []string
}

type EventType string

var (
	EventTypeWarning EventType = "Warning"
	EventTypeNormal  EventType = "Normal"
)

type EventSchema struct {
	Name                string
	Reason              string
	Action              string
	Type                EventType
	ReportingController string `yaml:"reportingController"`
	Message             string
}

func GetEvent(name string) (*EventSchema, error) {
	if _, ok := eventsMap[name]; !ok {
		return nil, fmt.Errorf("Invalid event")
	}
	return eventsMap[name], nil
}

func IsEventDisabled(name string) bool {
	controllerFilePath := "/events/event-schema/controller.yaml"
	workerFilePath := "/events/event-schema/worker.yaml"
	controllerConfigs, err := parseConfig(controllerFilePath)
	if err != nil {
		return false
	}
	workerConfigs, err := parseConfig(workerFilePath)
	if err != nil {
		return false
	}
	configs := append(controllerConfigs, workerConfigs...)
	for _, config := range configs {
		if config == name {
			return true
		}
	}
	return false
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
