package metrics_test

import (
	"context"
	"github.com/kubeslice/kubeslice-monitoring/pkg/metrics"
	"github.com/kubeslice/kubeslice-monitoring/pkg/schema"
	"testing"
	"time"
)

var (
	project     = "dummy"
	sliceName   = "red"
	clusterName = "cluster-1"
	namespace   = "kubeslice-dummy"
)

func TestRecordGaugeMetric(t *testing.T) {
	recorder := metrics.NewMetricRecorder(metrics.MetricRecorderOptions{
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "test-subsystem",
		Component: "controller",
	})

	err := recorder.RecordMetric(context.Background(), &metrics.Metric{
		Type:  metrics.MetricTypeGauge,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	if err != nil {
		t.Error("metric recording failed")
	}
}

func TestRecordCounterMetric(t *testing.T) {
	recorder := metrics.NewMetricRecorder(metrics.MetricRecorderOptions{
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "test-subsystem",
		Component: "controller",
	})

	err := recorder.RecordMetric(context.Background(), &metrics.Metric{
		Type:  metrics.MetricTypeCounter,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	if err != nil {
		t.Error("metric recording failed")
	}
}

func TestRecordHistogramMetric(t *testing.T) {
	recorder := metrics.NewMetricRecorder(metrics.MetricRecorderOptions{
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "test-subsystem",
		Component: "controller",
	})

	err := recorder.RecordMetric(context.Background(), &metrics.Metric{
		Type:  metrics.MetricTypeHistogram,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
		HistogramBuckets: []float64{1, 2, 3},
		Time:             time.Now(),
	})
	if err != nil {
		t.Error("metric recording failed")
	}
}

func TestRecordSummaryMetric(t *testing.T) {
	recorder := metrics.NewMetricRecorder(metrics.MetricRecorderOptions{
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "test-subsystem",
		Component: "controller",
	})

	err := recorder.RecordMetric(context.Background(), &metrics.Metric{
		Type:  metrics.MetricTypeSummary,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
		Time: time.Now(),
	})
	if err != nil {
		t.Error("metric recording failed")
	}
}
