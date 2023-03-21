package metrics

import (
	"context"
	"github.com/kubeslice/kubeslice-monitoring/pkg/logger"
	"github.com/kubeslice/kubeslice-monitoring/pkg/util"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"time"
)

// MetricRecorder is used to record metrics from a component
type MetricRecorder interface {
	// RecordMetric is used to record a new metric
	RecordMetric(context.Context, *Metric) error
	// WithSlice returns a new recorder with slice name added
	WithSlice(string) MetricRecorder
	// WithNamespace returns a new recorder with namespace name added
	WithNamespace(string) MetricRecorder
	// WithProject returns a new recorder with project name added
	WithProject(string) MetricRecorder
}

func NewMetricRecorder(o MetricRecorderOptions) MetricRecorder {
	log := logger.NewLogger().With("name", o.Component)
	return &metricRecorder{
		Logger:  log,
		Options: o,
	}
}

// MetricRecorderOptions provides a container with config parameters for the Prometheus Exporter
type MetricRecorderOptions struct {
	// Project is the name of the project
	Project string
	// Cluster  is the name of the cluster
	Cluster string
	// Slice is the name of the slice
	Slice string
	// Namespace is the namespace this metric recorder corresponds to
	Namespace string
	// Subsystem is the subsystem this metric recorder corresponds to (optional)
	Subsystem string
	// Component is the component which uses the metric recorder
	Component string
}

type Metric struct {
	// Type is the type of the metric
	Type MetricType
	// Name is the name of the metric
	Name string
	// Help is the help text of the metric
	Help string
	// Value is the value of the metric
	Value float64
	// Labels are the labels of the metric
	Labels map[string]string
	// HistogramBuckets are the buckets of the histogram metric
	HistogramBuckets []float64
	// Time is the time of the metric for histogram and summary metrics (optional)
	Time time.Time
}

type metricRecorder struct {
	Logger  *zap.SugaredLogger
	Options MetricRecorderOptions
}

// WithSlice returns a new recorder with added slice name for raising metrics
func (mr *metricRecorder) WithSlice(slice string) MetricRecorder {
	m := mr.Copy()
	m.Options.Slice = slice
	return m
}

// WithNamespace returns a new recorder with added namespace name for raising metrics
func (mr *metricRecorder) WithNamespace(ns string) MetricRecorder {
	m := mr.Copy()
	m.Options.Namespace = ns
	return m
}

// WithProject returns a new recorder with added project name for raising metrics
func (mr *metricRecorder) WithProject(project string) MetricRecorder {
	m := mr.Copy()
	m.Options.Project = project
	return m
}

type MetricType string

var fixedLabels = []string{"slice_project", "slice_cluster", "slice_name"}

const (
	MetricTypeGauge     = "Gauge"
	MetricTypeCounter   = "Counter"
	MetricTypeHistogram = "Histogram"
	MetricTypeSummary   = "Summary"
)

// Copy returns a Provider that produces Prometheus metrics.
// Namespace and subsystem are applied to all produced metrics.
func (mr *metricRecorder) Copy() *metricRecorder {
	return &metricRecorder{
		Logger: mr.Logger,
		Options: MetricRecorderOptions{
			Project:   mr.Options.Project,
			Cluster:   mr.Options.Cluster,
			Slice:     mr.Options.Slice,
			Namespace: mr.Options.Namespace,
			Subsystem: mr.Options.Subsystem,
			Component: mr.Options.Component,
		},
	}
}

// RecordMetric records a metric to the Prometheus Exporter
func (mr *metricRecorder) RecordMetric(ctx context.Context, m *Metric) error {
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

// recordGauge records a gauge metric to the Prometheus Exporter
func (mr *metricRecorder) recordGauge(ctx context.Context, m *Metric) {
	mr.Logger.Debugf("Recording gauge metric: %v", m)
	metric := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: util.FlattenString(mr.Options.Namespace),
			Subsystem: util.FlattenString(mr.Options.Subsystem),
			Name:      util.FlattenString(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenString(mr.Options.Project),
				"slice_cluster": util.FlattenString(mr.Options.Cluster),
				"slice_name":    util.FlattenString(mr.Options.Slice),
			}, util.FlattenMap(m.Labels)),
	).Set(m.Value)
}

// recordCounter records a counter metric to the Prometheus Exporter
func (mr *metricRecorder) recordCounter(ctx context.Context, m *Metric) {
	mr.Logger.Debugf("Recording counter metric: %v", m)
	metric := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: util.FlattenString(mr.Options.Namespace),
			Subsystem: util.FlattenString(mr.Options.Subsystem),
			Name:      util.FlattenString(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenString(mr.Options.Project),
				"slice_cluster": util.FlattenString(mr.Options.Cluster),
				"slice_name":    util.FlattenString(mr.Options.Slice),
			}, util.FlattenMap(m.Labels)),
	).Add(m.Value)
}

// recordHistogram records a histogram metric to the Prometheus Exporter
func (mr *metricRecorder) recordHistogram(ctx context.Context, m *Metric) {
	mr.Logger.Debugf("Recording histogram metric: %v", m)
	metric := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: util.FlattenString(mr.Options.Namespace),
			Subsystem: util.FlattenString(mr.Options.Subsystem),
			Name:      util.FlattenString(m.Name),
			Help:      m.Help,
			Buckets:   m.HistogramBuckets,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenString(mr.Options.Project),
				"slice_cluster": util.FlattenString(mr.Options.Cluster),
				"slice_name":    util.FlattenString(mr.Options.Slice),
			}, util.FlattenMap(m.Labels)),
	).Observe(time.Since(m.Time).Seconds())
}

// recordSummary records a summary metric to the Prometheus Exporter
func (mr *metricRecorder) recordSummary(ctx context.Context, m *Metric) {
	mr.Logger.Debugf("Recording summary metric: %v", m)
	metric := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: util.FlattenString(mr.Options.Namespace),
			Subsystem: util.FlattenString(mr.Options.Subsystem),
			Name:      util.FlattenString(m.Name),
			Help:      m.Help,
		},
		append(fixedLabels, util.KeysFromMap(m.Labels)...),
	)
	metric.With(
		util.MergeMaps(
			prometheus.Labels{
				"slice_project": util.FlattenString(mr.Options.Project),
				"slice_cluster": util.FlattenString(mr.Options.Cluster),
				"slice_name":    util.FlattenString(mr.Options.Slice),
			}, util.FlattenMap(m.Labels)),
	).Observe(time.Since(m.Time).Seconds())
}
