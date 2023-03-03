package schema

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/yaml"
	"os"
	"path"
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
	Enabled             bool
}

func parseEvent1(filepath string) ([]EventSchema, error) {
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

func main() {
	// Read the file
	file1 := "controller.yaml"
	controllerFilePath := path.Join("config/events", file1)
	controllerEvents, err := parseEvent1(controllerFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the file
	file, err := os.Create("pkg/schema/event_detail.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "package schema\n\n"+"var Events = map[string]EventSchema {\n")
	// Write the struct to the file
	for _, event := range controllerEvents {
		_, err = fmt.Fprintf(file, "  Event%s: {\n"+
			"    Reason: \"%s\",\n"+
			"    Action: \"%s\",\n"+
			"    Type: \"%s\",\n"+
			"    ReportingController: \"%s\",\n"+
			"    Message: \"%s\",\n"+
			"  },\n",
			event.Name, event.Reason, event.Action, event.Type, event.ReportingController, event.Message)
	}
	fmt.Fprintf(file, "}")
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
