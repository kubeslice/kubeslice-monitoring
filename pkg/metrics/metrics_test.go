package metrics

import (
	"context"
	"github.com/kubeslice/kubeslice-monitoring/pkg/logger"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"
)

func TestRecordGaugeMetric(t *testing.T) {
	project := "cisco"
	sliceName := "red"
	clusterName := "cluster-1"
	namespace := "kubeslice-cisco"
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Scheme:    newTestScheme(),
		Version:   "1",
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeGauge,
		Name:  SliceNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func TestRecordCounterMetric(t *testing.T) {
	project := "cisco"
	sliceName := "red"
	clusterName := "cluster-1"
	namespace := "kubeslice-cisco"
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Scheme:    newTestScheme(),
		Version:   "1",
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeCounter,
		Name:  SliceNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func TestRecordHistogramMetric(t *testing.T) {
	project := "cisco"
	sliceName := "red"
	clusterName := "cluster-1"
	namespace := "kubeslice-cisco"
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Scheme:    newTestScheme(),
		Version:   "1",
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeHistogram,
		Name:  SliceNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func TestRecordSummaryMetric(t *testing.T) {
	project := "cisco"
	sliceName := "red"
	clusterName := "cluster-1"
	namespace := "kubeslice-cisco"
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Scheme:    newTestScheme(),
		Version:   "1",
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeSummary,
		Name:  SliceNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func newTestScheme() *runtime.Scheme {
	testScheme := runtime.NewScheme()
	_ = corev1.AddToScheme(testScheme)
	return testScheme
}
