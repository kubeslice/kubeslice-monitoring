# kubeslice-monitoring

kubeslice-monitoring : Repository for kubeslice-monitoring package

## Event Schema

It contains the following fields for each event

* name: Name of the event
* reason: Reason for raising the event
* action: Action which caused the event or the action which need to be taken
* type: Warning or Normal
* reportingController: Name of the component which reported the event
* message: Human readable message explaining the event


## Using events framework in your component

1. Copy the `hack/` directory to your component.

It contains `generate/generate.go` and `templates/schema.tmpl` files. You will require both to create a helper to generate Events Map. 
Modify template and output file paths in `generate.go` if needed.

2. Create a events schema yaml for your component in the following format:

```yaml
events:
  - name: ExampleEvent
    reason: ExampleEvent
    action: ExampleEvent
    type: Warning
    reportingController: controller
    message: ExampleEvent message.
  
```

3. Add the following lines to your Makefile. (Make path modifications if required.)

```Makefile
.PHONY: generate-events
generate-events:
	go run hack/generate/generate.go <PATH-TO-EVENT-SCHEMA-YAML>
	go fmt ./...
```

4. Generate Events Map

Run this command after making any changes to event schema files

```shell
make generate-events
```
This will generate `events_generated.go` file in the events directory. (Check file output path in `hack/generate/generate.go`)


5. Import events package from monitoring framework and the generated Events Map from your component.

```go
import(
	"github.com/kubeslice/kubeslice-monitoring/pkg/events"
	componentEvents "github.com/kubeslice/<YOUR-COMPONENT>/events"
)
```

6. Initialize event recorder


```go
recorder := events.NewEventRecorder(k8sClient, schema, componentEvents.EventsMap, events.EventRecorderOptions{
  Version:   "1",
  Cluster:   "cluster-1",
  Component: "controller",
  Namespace: "test-ns",
  Slice: "test-clice",
  Project: "test-project",
})
```

Slice and Namespace names are optional if the recorder is not part of a slice specific or namespace specific component.

7. Raise events

```go
err := recorder.RecordEvent(ctx, &events.Event{
  Object:            obj,
  RelatedObject:     robj,
  ReportingInstance: "controller",
  Name:              schema.EventSliceConfigDeletionFailed,
})
```

8. Raise events with slice/namespace name added at the time of raising events

In some cases, the recorder will be part of a controller which manages multiple namespaces. In that case,
events can be raises by providing namespace like below instead of Initializing the recorder with specific namespace name.

```go
recorder.WithNamespace(ns).RecordEvent(...)

recorder.WithSlice(sliceName).RecordEvent(...)
```


## Disabling events
TODO: fill this section