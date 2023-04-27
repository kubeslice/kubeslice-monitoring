package events_test

import (
	"testing"

	controllerv1alpha1 "github.com/kubeslice/apis/pkg/controller/v1alpha1"
	"github.com/kubeslice/kubeslice-monitoring/pkg/events"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type k8sClientMock struct {
	createdObject client.Object
}

func (o *k8sClientMock) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	o.createdObject = obj
	return nil
}

func (o *k8sClientMock) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}

func (o *k8sClientMock) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return nil
}

func (o *k8sClientMock) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (o *k8sClientMock) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}

func TestRecordEvent(t *testing.T) {
	clientMock := &k8sClientMock{}

	recorder := events.NewEventRecorder(clientMock, newTestScheme(), events.EventsMap, events.EventRecorderOptions{
		Version:   "1",
		Cluster:   "cluster-1",
		Component: "controller",
		Namespace: "test-ns",
	})

	pod := &corev1.Pod{}

	err := recorder.RecordEvent(context.Background(), &events.Event{
		Object:            pod,
		RelatedObject:     nil,
		ReportingInstance: "controller",
		Name:              events.EventExampleEvent,
	})
	if err != nil {
		t.Error("event not recorded")
	}

	e := clientMock.createdObject.(*corev1.Event)
	if e.Namespace != "test-ns" {
		t.Error("invalid ns")
	}

	// TODO: test remaining fields
}

func newTestScheme() *runtime.Scheme {
	testScheme := runtime.NewScheme()
	_ = corev1.AddToScheme(testScheme)
	_ = controllerv1alpha1.AddToScheme(testScheme)
	return testScheme
}
