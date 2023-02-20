package metrics

import (
	"context"
	"github.com/kubeslice/kubeslice-monitoring/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"time"
)

// MetricRecorder provides a container with config parameters for the Prometheus Exporter
type MetricRecorder struct {
	Logger    *zap.SugaredLogger
	Scheme    *runtime.Scheme
	Version   string
	Project   string
	Cluster   string
	Slice     string
	Namespace string
	Subsystem string
}

type Metric struct {
	Type             MetricType
	Name             string
	Help             string
	Value            float64
	Labels           map[string]string
	histogramBuckets []float64
}

type MetricType string

var fixedLabels = []string{"slice_project", "slice_cluster", "slice_name"}

const (
	MetricTypeGauge     = "Gauge"
	MetricTypeCounter   = "Counter"
	MetricTypeHistogram = "Histogram"
	MetricTypeSummary   = "Summary"
)

// NewPrometheusProvider returns a Provider that produces Prometheus metrics.
// Namespace and subsystem are applied to all produced metrics.
func (mr *MetricRecorder) NewPrometheusProvider() *MetricRecorder {
	return &MetricRecorder{
		Logger:    mr.Logger,
		Scheme:    mr.Scheme,
		Version:   mr.Version,
		Project:   mr.Project,
		Cluster:   mr.Cluster,
		Slice:     mr.Slice,
		Namespace: mr.Namespace,
		Subsystem: mr.Subsystem,
	}
}

func (mr *MetricRecorder) RecordMetric(ctx context.Context, m *Metric) error {
	switch m.Type {
	case MetricTypeGauge:
		mr.recordGauge(ctx, m)
		break
	case MetricTypeCounter:
		mr.recordCounter(ctx, m)
		break
	case MetricTypeHistogram:
		mr.recordHistogram(ctx, m)
		break
	case MetricTypeSummary:
		mr.recordSummary(ctx, m)
		break
	default:
		mr.Logger.Errorf("Unknown metric type: %v", m.Type)
	}
	return nil
}

func (mr *MetricRecorder) recordGauge(ctx context.Context, m *Metric) {
	mr.Logger.Infof("Recording gauge metric: %v", m)
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: util.FlattenKey(mr.Namespace),
			Subsystem: util.FlattenKey(mr.Subsystem),
			Name:      util.FlattenKey(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenKey(mr.Project),
				"slice_cluster": util.FlattenKey(mr.Cluster),
				"slice_name":    util.FlattenKey(mr.Slice),
			}, util.FlattenMap(m.Labels)),
	).Set(m.Value)
}

func (mr *MetricRecorder) recordCounter(ctx context.Context, m *Metric) {
	mr.Logger.Infof("Recording counter metric: %v", m)
	metric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: util.FlattenKey(mr.Namespace),
			Subsystem: util.FlattenKey(mr.Subsystem),
			Name:      util.FlattenKey(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenKey(mr.Project),
				"slice_cluster": util.FlattenKey(mr.Cluster),
				"slice_name":    util.FlattenKey(mr.Slice),
			}, util.FlattenMap(m.Labels)),
	).Add(m.Value)
}

func (mr *MetricRecorder) recordHistogram(ctx context.Context, m *Metric) {
	mr.Logger.Infof("Recording histogram metric: %v", m)
	start := time.Now()
	metric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: util.FlattenKey(mr.Namespace),
			Subsystem: util.FlattenKey(mr.Subsystem),
			Name:      util.FlattenKey(m.Name),
			Help:      m.Help,
			Buckets:   m.histogramBuckets,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenKey(mr.Project),
				"slice_cluster": util.FlattenKey(mr.Cluster),
				"slice_name":    util.FlattenKey(mr.Slice),
			}, util.FlattenMap(m.Labels)),
	).Observe(time.Since(start).Seconds())
}

func (mr *MetricRecorder) recordSummary(ctx context.Context, m *Metric) {
	mr.Logger.Infof("Recording summary metric: %v", m)
	start := time.Now()
	metric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: util.FlattenKey(mr.Namespace),
			Subsystem: util.FlattenKey(mr.Subsystem),
			Name:      util.FlattenKey(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenKey(mr.Project),
				"slice_cluster": util.FlattenKey(mr.Cluster),
				"slice_name":    util.FlattenKey(mr.Slice),
			}, util.FlattenMap(m.Labels)),
	).Observe(time.Since(start).Seconds())
}
