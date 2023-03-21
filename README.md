# kubeslice-monitoring

kubeslice-monitoring : Repository for kubeslice-monitoring package

## Using event schema

Event schema files are located in `config/events/controller.yaml` and `config/events/worker.yaml` respectively. It contains the following fields for each event

* name: Name of the event
* reason: Reason for raising the event
* action: Action which caused the event or the action which need to be taken
* type: Warning or Normal
* reportingController: Name of the component which reported the event
* message: Human readable message explaining the event

make sure to run `make generate` after modifying the above files

## Disabling events
TODO: fill this section

## Generate parsed event schema code

Run this command after making any changes to event schema files

```
make generate
```

## Using events framework in your component


1. Import events package and schema

```
import(
	"github.com/kubeslice/kubeslice-monitoring/pkg/events"
	"github.com/kubeslice/kubeslice-monitoring/pkg/schema"
)
```

2. Initialize event recorder


```
recorder := events.NewEventRecorder(k8sClient, schema, events.EventRecorderOptions{
  Version:   "1",
  Cluster:   "cluster-1",
  Component: "controller",
  Namespace: "test-ns",
  Slice: "test-clice",
  Project: "test-project",
})
```

Slice and Namespace names are optional if the recorder is not part of a slice specific or namespace specific component.

3. Raise events

```
err := recorder.RecordEvent(ctx, &events.Event{
  Object:            obj,
  RelatedObject:     robj,
  ReportingInstance: "controller",
  Name:              schema.EventSliceConfigDeletionFailed,
})
```

4. Raise events with slice/namespace/project name added at the time of raising events

In some cases, the recorder will be part of a controller which manages multiple namespaces. In that case,
events can be raises by providing namespace like below instead of Initializing the recorder with specific namespace name.

```
recorder.WithNamespace(ns).RecordEvent(...)

recorder.WithSlice(sliceName).RecordEvent(...)

recorder.WithProject(projectName).RecordEvent(...)
```


## Using metrics framework in your component


1. Import metrics package

```
import(
    "github.com/kubeslice/kubeslice-monitoring/pkg/metrics"
)
```

2. Initialize metric recorder


```
recorder := metrics.NewMetricRecorder(metrics.MetricRecorderOptions{
    Project:   "test-project",
    Cluster:   "cluster-1",
    Slice:     "test-slice",
    Namespace: "test-namepsace",
    Subsystem: "test-subsystem",
    Component: "controller",
})
```

3. Record metric

```
err := recorder.RecordMetric(context.Background(), &metrics.Metric{
    Type:  metrics.MetricTypeGauge,
    Name:  metrics.MetricNetPolViolation,
    Help:  "test metric help",
    Value: 1,
    Labels: map[string]string{
        "test_key": "test_value",
    },
})
```

4. Record metric with slice/namespace/project name added at the time of recording metrics

In some cases, the recorder will be part of a controller which manages multiple namespaces. In that case,
metrics can be recorded by providing namespace like below instead of Initializing the recorder with specific namespace name.

```
recorder.WithNamespace(ns).RecordMetric(...)

recorder.WithSlice(sliceName).RecordMetric(...)

recorder.WithProject(projectName).RecordMetric(...)
```
