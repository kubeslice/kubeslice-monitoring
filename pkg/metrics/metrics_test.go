package metrics

import (
	"context"
	"github.com/kubeslice/kubeslice-monitoring/pkg/logger"
	"github.com/kubeslice/kubeslice-monitoring/pkg/schema"
	"github.com/stretchr/testify/require"
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
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeGauge,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func TestRecordCounterMetric(t *testing.T) {
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeCounter,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
	})
	require.Nil(t, err)
}

func TestRecordHistogramMetric(t *testing.T) {
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeHistogram,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
		HistogramBuckets: []float64{1, 2, 3},
		Time:             time.Now(),
	})
	require.Nil(t, err)
}

func TestRecordSummaryMetric(t *testing.T) {
	recorder := MetricRecorder{
		Logger:    logger.NewLogger(),
		Project:   project,
		Cluster:   clusterName,
		Slice:     sliceName,
		Namespace: namespace,
		Subsystem: "controller",
	}

	err := recorder.RecordMetric(context.Background(), &Metric{
		Type:  MetricTypeSummary,
		Name:  schema.MetricNetPolViolation,
		Help:  "test metric help",
		Value: 1,
		Labels: map[string]string{
			"test_key": "test_value",
		},
		Time: time.Now(),
	})
	require.Nil(t, err)
}
