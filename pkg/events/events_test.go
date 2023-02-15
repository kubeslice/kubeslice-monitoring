package events

import (
	controllerv1alpha1 "github.com/kubeslice/apis/pkg/controller/v1alpha1"
	"github.com/kubeslice/kubeslice-monitoring/pkg/logger"
	"github.com/kubeslice/kubeslice-monitoring/pkg/schema"
	"github.com/kubeslice/kubeslice-monitoring/pkg/util/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"strings"
	"testing"
)

func TestRecordEvent(t *testing.T) {
	clientMock := &mocks.Client{}
	project := "cisco"
	sliceName := "red"
	clusterName := "cluster-1"
	namespace := "kubeslice-cisco"
	recorder := EventRecorder{
		Client:    clientMock,
		Logger:    logger.NewLogger(),
		Scheme:    newTestScheme(),
		Version:   "1",
		Cluster:   clusterName,
		Tenant:    project,
		Slice:     sliceName,
		Namespace: namespace,
		Component: "controller",
	}
	sliceConfig := &controllerv1alpha1.SliceConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      recorder.Slice,
			Namespace: recorder.Namespace,
		},
	}

	event, err := schema.GetEvent(schema.EventSliceDeletionFailed)
	require.Nil(t, err)
	clientMock.On("Create", context.Background(), mock.MatchedBy(func(evt *corev1.Event) bool {
		return !evt.FirstTimestamp.IsZero() && strings.HasPrefix(evt.Name, sliceName) && evt.Namespace == namespace &&
			evt.InvolvedObject.Kind == "SliceConfig" && evt.InvolvedObject.Name == sliceName && evt.InvolvedObject.Namespace == namespace &&
			evt.Reason == event.Reason && evt.Action == event.Action && evt.Type == string(event.Type) &&
			evt.ReportingController == event.ReportingController && evt.Message == event.Message
	})).Return(nil)

	err = recorder.RecordEvent(context.Background(), &Event{
		Object:            sliceConfig,
		RelatedObject:     nil,
		ReportingInstance: "controller",
		Name:              schema.EventSliceDeletionFailed,
	})
	require.Nil(t, err)
	clientMock.AssertExpectations(t)
}

func newTestScheme() *runtime.Scheme {
	testScheme := runtime.NewScheme()
	_ = corev1.AddToScheme(testScheme)
	_ = controllerv1alpha1.AddToScheme(testScheme)
	return testScheme
}
