package events_test

import (
	"testing"

	controllerv1alpha1 "github.com/kubeslice/apis/pkg/controller/v1alpha1"
	"github.com/kubeslice/kubeslice-monitoring/pkg/events"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	o.createdObject = obj
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

func makeObjectReference(kind, name, namespace string) corev1.ObjectReference {
	return corev1.ObjectReference{
		Kind:       kind,
		Name:       name,
		Namespace:  namespace,
		UID:        "C934D34AFB20242",
		APIVersion: "version",
		FieldPath:  "spec.containers{mycontainer}",
	}
}

func makeEvent(reason, message string, involvedObject corev1.ObjectReference) corev1.Event {
	eventTime := metav1.Now()
	event := corev1.Event{
		Reason:         reason,
		Message:        message,
		InvolvedObject: involvedObject,
		Source: corev1.EventSource{
			Component: "kubelet",
			Host:      "kublet.node1",
		},
		Count:          1,
		FirstTimestamp: eventTime,
		LastTimestamp:  eventTime,
		Type:           corev1.EventTypeNormal,
	}
	return event
}

func TestEventAggregatorByReasonFunc(t *testing.T) {
	event1 := makeEvent("end-of-world", "it was fun", makeObjectReference("Pod", "pod1", "other"))
	event2 := makeEvent("end-of-world", "it was fun", makeObjectReference("Pod", "pod1", "other"))
	event3 := makeEvent("nevermind", "it was a bug", makeObjectReference("Pod", "pod1", "other"))

	localKey1 := events.GetEventKey(&event1)
	localKey2 := events.GetEventKey(&event2)
	localKey3 := events.GetEventKey(&event3)

	if localKey1 != localKey2 {
		t.Errorf("Expected %v to equal %v", localKey1, localKey2)
	}
	if localKey1 == localKey3 {
		t.Errorf("Expected %v to not equal %v", localKey1, localKey3)
	}
}

func TestRecordEventWithCache(t *testing.T) {
	clientMock := &k8sClientMock{}

	recorder := events.NewEventRecorder(clientMock, newTestScheme(), events.EventsMap, events.EventRecorderOptions{
		Version:   "1",
		Cluster:   "cluster-1",
		Component: "controller",
		Namespace: "test-ns",
	})

	pod := &corev1.Pod{}

	// first-event
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
	if e.Count != 1 {
		t.Error("invalid Count")
	}

	// duplicate event
	err = recorder.RecordEvent(context.Background(), &events.Event{
		Object:            pod,
		RelatedObject:     nil,
		ReportingInstance: "controller",
		Name:              events.EventExampleEvent,
	})
	if err != nil {
		t.Error("event not recorded")
	}

	e = clientMock.createdObject.(*corev1.Event)
	if e.Count != 2 {
		t.Error("invalid Count")
	}
}

func TestEventWithEmptyNamespaceReference(t *testing.T) {
	clientMock := &k8sClientMock{}

	recorder := events.NewEventRecorder(clientMock, newTestScheme(), events.EventsMap, events.EventRecorderOptions{
		Version:   "1",
		Cluster:   "cluster-1",
		Component: "controller",
		Namespace: "test",
	})

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-pod",
		},
	}
	err := recorder.RecordEvent(context.Background(), &events.Event{
		Object:            pod,
		RelatedObject:     nil,
		ReportingInstance: "controller",
		Name:              events.EventExampleEvent,
	})
	if err != nil {
		t.Error("event not recorded ", err)
	}

	e := clientMock.createdObject.(*corev1.Event)
	if e.InvolvedObject.Name == "" {
		t.Error("InvolvedObject Name is empty")
	}
	if e.InvolvedObject.Namespace == "" {
		t.Error("InvolvedObject NameSpace is empty")
	}
	if e.InvolvedObject.Name != e.InvolvedObject.Namespace {
		t.Error("involved object name and namespace do not match")
	}
}
