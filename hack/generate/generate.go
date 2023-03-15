package main

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kubeslice/kubeslice-monitoring/pkg/schema"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("generating event schema code from schema file")

	file1 := "controller.yaml"
	file2 := "worker.yaml"
	pwd, err := os.Getwd()
	handleError(err)
	controllerFilePath := path.Join(pwd, "config/events", file1)
	workerFilePath := path.Join(pwd, "config/events", file2)
	controllerEvents, err := parseEvent(controllerFilePath)
	handleError(err)
	workerEvents, err := parseEvent(workerFilePath)
	handleError(err)

	events := append(controllerEvents, workerEvents...)

	t, err := template.ParseFiles("hack/templates/schema.tmpl")
	handleError(err)

	f, err := os.Create("pkg/schema/events_generated.go")
	handleError(err)
	t.Execute(f, events)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseEvent(filepath string) ([]schema.EventSchema, error) {
	var eventSchema struct {
		Events []schema.EventSchema
	}
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
