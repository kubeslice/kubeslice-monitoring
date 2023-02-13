package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	metricRecorder := MetricRecorder{
		Logger:        nil,
		Scheme:        nil,
		Version:       "1",
		Project:       "avesha",
		Cluster:       "cluster",
		Slice:         "red",
		Namespace:     "ns",
		Subsystem:     "controller",
		Registry:      metrics.DefaultRegistry,
		PromRegistry:  prometheus.DefaultRegisterer,
		FlushInterval: 1 * time.Second,
	}
	prometheusClient := metricRecorder.NewPrometheusProvider()
	go prometheusClient.UpdatePrometheusMetrics()
}

type MetricRecorder struct {
	Logger        *zap.SugaredLogger
	Scheme        *runtime.Scheme
	Version       string
	Project       string
	Cluster       string
	Slice         string
	Namespace     string
	Subsystem     string
	Registry      metrics.Registry
	PromRegistry  prometheus.Registerer
	FlushInterval time.Duration
}

// PrometheusConfig provides a container with config parameters for the
// Prometheus Exporter
type PrometheusConfig struct {
	Logger           *zap.SugaredLogger
	Scheme           *runtime.Scheme
	Version          string
	Project          string
	Cluster          string
	Slice            string
	Namespace        string
	Subsystem        string
	Registry         metrics.Registry      // Registry to be exported
	PromRegistry     prometheus.Registerer //Prometheus registry
	FlushInterval    time.Duration         //interval to update prom metrics
	gauges           map[string]prometheus.Gauge
	customMetrics    map[string]*CustomCollector
	histogramBuckets []float64
	timerBuckets     []float64
	mutex            *sync.Mutex
}

// NewPrometheusProvider returns a Provider that produces Prometheus metrics.
// Namespace and subsystem are applied to all produced metrics.
func (mr *MetricRecorder) NewPrometheusProvider() *PrometheusConfig {
	return &PrometheusConfig{
		Logger:           mr.Logger,
		Scheme:           mr.Scheme,
		Version:          mr.Version,
		Project:          mr.Project,
		Cluster:          mr.Cluster,
		Slice:            mr.Slice,
		Namespace:        mr.Namespace,
		Subsystem:        mr.Subsystem,
		Registry:         mr.Registry,
		PromRegistry:     mr.PromRegistry,
		FlushInterval:    mr.FlushInterval,
		gauges:           make(map[string]prometheus.Gauge),
		customMetrics:    make(map[string]*CustomCollector),
		histogramBuckets: []float64{0.05, 0.1, 0.25, 0.50, 0.75, 0.9, 0.95, 0.99},
		timerBuckets:     []float64{0.50, 0.95, 0.99, 0.999},
		mutex:            new(sync.Mutex),
	}
}

func (pc *PrometheusConfig) UpdatePrometheusMetrics() {
	for range time.Tick(pc.FlushInterval) {
		pc.UpdatePrometheusMetricsOnce()
	}
}

func (pc *PrometheusConfig) gaugeFromNameAndValue(name string, val float64) {
	key := pc.createKey(name)
	g, ok := pc.gauges[key]
	if !ok {
		g = prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: pc.flattenKey(pc.Namespace),
			Subsystem: pc.flattenKey(pc.Subsystem),
			Name:      pc.flattenKey(name),
			Help:      name,
		})
		pc.PromRegistry.Register(g)
		pc.gauges[key] = g
	}
	g.Set(val)
}

func (pc *PrometheusConfig) flattenKey(key string) string {
	key = strings.Replace(key, " ", "_", -1)
	key = strings.Replace(key, ".", "_", -1)
	key = strings.Replace(key, "-", "_", -1)
	key = strings.Replace(key, "=", "_", -1)
	key = strings.Replace(key, "/", "_", -1)
	return key
}

func (pc *PrometheusConfig) createKey(name string) string {
	return fmt.Sprintf("%s_%s_%s", pc.Namespace, pc.Subsystem, name)
}

func (pc *PrometheusConfig) UpdatePrometheusMetricsOnce() error {
	pc.Registry.Each(func(name string, i interface{}) {
		switch metric := i.(type) {
		case metrics.Counter:
			pc.gaugeFromNameAndValue(name, float64(metric.Count()))
		case metrics.Gauge:
			pc.gaugeFromNameAndValue(name, float64(metric.Value()))
		case metrics.GaugeFloat64:
			pc.gaugeFromNameAndValue(name, metric.Value())
		case metrics.Histogram:
			samples := metric.Snapshot().Sample().Values()
			if len(samples) > 0 {
				lastSample := samples[len(samples)-1]
				pc.gaugeFromNameAndValue(name, float64(lastSample))
			}
			pc.histogramFromNameAndMetric(name, metric, pc.histogramBuckets)
		case metrics.Meter:
			snapshot := metric.Snapshot()
			pc.gaugeFromNameAndValue(name+"_rate1", snapshot.Rate1())
			pc.gaugeFromNameAndValue(name+"_rate5", snapshot.Rate5())
			pc.gaugeFromNameAndValue(name+"_rate15", snapshot.Rate15())
			pc.gaugeFromNameAndValue(name+"_rate_mean", snapshot.RateMean())
			pc.gaugeFromNameAndValue(name+"_count", float64(snapshot.Count()))
		case metrics.Timer:
			snapshot := metric.Snapshot()
			pc.gaugeFromNameAndValue(name+"_rate1", snapshot.Rate1())
			pc.gaugeFromNameAndValue(name+"_rate5", snapshot.Rate5())
			pc.gaugeFromNameAndValue(name+"_rate15", snapshot.Rate15())
			pc.gaugeFromNameAndValue(name+"_rate_mean", snapshot.RateMean())
			pc.gaugeFromNameAndValue(name+"_count", float64(snapshot.Count()))
			pc.gaugeFromNameAndValue(name+"_sum", float64(snapshot.Sum()))
			pc.gaugeFromNameAndValue(name+"_max", float64(snapshot.Max()))
			pc.gaugeFromNameAndValue(name+"_min", float64(snapshot.Min()))
			pc.gaugeFromNameAndValue(name+"_mean", snapshot.Mean())
			pc.gaugeFromNameAndValue(name+"_variance", snapshot.Variance())
			pc.gaugeFromNameAndValue(name+"_std_dev", snapshot.StdDev())
			pc.histogramFromNameAndMetric(name, metric, pc.timerBuckets)
		}
	})
	return nil
}

func (pc *PrometheusConfig) histogramFromNameAndMetric(name string, goMetric interface{}, buckets []float64) {
	key := pc.createKey(name)

	collector, ok := pc.customMetrics[key]
	if !ok {
		collector = NewCustomCollector(pc.mutex)
		pc.PromRegistry.MustRegister(collector)
		pc.customMetrics[key] = collector
	}

	var ps []float64
	var count uint64
	var sum float64
	var typeName string

	switch metric := goMetric.(type) {
	case metrics.Histogram:
		snapshot := metric.Snapshot()
		ps = snapshot.Percentiles(buckets)
		count = uint64(snapshot.Count())
		sum = float64(snapshot.Sum())
		typeName = "histogram"
	case metrics.Timer:
		snapshot := metric.Snapshot()
		ps = snapshot.Percentiles(buckets)
		count = uint64(snapshot.Count())
		sum = float64(snapshot.Sum())
		typeName = "timer"
	default:
		panic(fmt.Sprintf("unexpected metric type %T", goMetric))
	}

	bucketVals := make(map[float64]uint64)
	for ii, bucket := range buckets {
		bucketVals[bucket] = uint64(ps[ii])
	}

	desc := prometheus.NewDesc(
		prometheus.BuildFQName(
			pc.flattenKey(pc.Namespace),
			pc.flattenKey(pc.Subsystem),
			fmt.Sprintf("%s_%s", pc.flattenKey(name), typeName),
		),
		pc.flattenKey(name),
		[]string{},
		map[string]string{},
	)

	if constHistogram, err := prometheus.NewConstHistogram(
		desc,
		count,
		sum,
		bucketVals,
	); err == nil {
		pc.mutex.Lock()
		collector.metric = constHistogram
		pc.mutex.Unlock()
	}
}

// for collecting prometheus.constHistogram objects
type CustomCollector struct {
	prometheus.Collector

	metric prometheus.Metric
	mutex  *sync.Mutex
}

func NewCustomCollector(mutex *sync.Mutex) *CustomCollector {
	return &CustomCollector{
		mutex: mutex,
	}
}
