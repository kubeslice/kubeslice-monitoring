package metrics_test

import (
	"testing"

	"github.com/kubeslice/kubeslice-monitoring/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

// Make sure that missing constant labels are handled when initializing metrics
func TestMissingLabels(t *testing.T) {
	r := prometheus.NewRegistry()
	mf, err := metrics.NewMetricsFactory(r, metrics.MetricsFactoryOptions{
		Cluster: "test-cluster",
	})

	if err != nil {
		t.Error(err, "failed to initialize metrics factory")
	}

	c := mf.NewCounter("test_counter", "help", []string{"a"})

	m := &dto.Metric{}
	err = c.WithLabelValues("b").Write(m)
	if err != nil {
		t.Error(err, "failed to write counter")
	}

	metrics, _ := r.Gather()
	if expected, got := "kubeslice_test_counter", *metrics[0].Name; expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
	if expected, got := "help", *metrics[0].Help; expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}

	if expected, got := `label:<name:"a" value:"b" > label:<name:"slice_cluster" value:"test-cluster" > counter:<value:0 > `, m.String(); expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

// Make sure constant labels are present in the metric
func TestCurryLabels(t *testing.T) {
	r := prometheus.NewRegistry()
	mf, err := metrics.NewMetricsFactory(r, metrics.MetricsFactoryOptions{
		Cluster:             "cl",
		Project:             "pr",
		Namespace:           "ns",
		ReportingController: "rc",
		Slice:               "sl",
	})

	if err != nil {
		t.Error(err, "failed to initialize metrics factory")
	}

	c := mf.NewCounter("test_counter", "help", []string{})
	g := mf.NewGauge("test_gauge", "help", []string{})

	m := &dto.Metric{}
	err = c.WithLabelValues().Write(m)
	if err != nil {
		t.Error(err, "failed to write counter")
	}

	if expected, got := `label:<name:"slice_cluster" value:"cl" > label:<name:"slice_name" value:"sl" > label:<name:"slice_namespace" value:"ns" > label:<name:"slice_project" value:"pr" > label:<name:"slice_reporting_controller" value:"rc" > counter:<value:0 > `, m.String(); expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}

	m.Reset()
	err = g.WithLabelValues().Write(m)
	if err != nil {
		t.Error(err, "failed to write gauge")
	}

	if expected, got := `label:<name:"slice_cluster" value:"cl" > label:<name:"slice_name" value:"sl" > label:<name:"slice_namespace" value:"ns" > label:<name:"slice_project" value:"pr" > label:<name:"slice_reporting_controller" value:"rc" > gauge:<value:0 > `, m.String(); expected != got {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
