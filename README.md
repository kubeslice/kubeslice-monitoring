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

### Generate parsed event schema code

Run this command after making any changes to event schema files

```
make generate
```
