package tests

import (
	"testing"

	prometheus "github.com/bolanosdev/prometheus-mux-monitor"
	prom "github.com/prometheus/client_golang/prometheus"
)

func TestGetMonitor(t *testing.T) {
	m := prometheus.GetMonitor()
	if m == nil {
		t.Fatal("expected monitor to be initialized")
	}

	m2 := prometheus.GetMonitor()
	if m != m2 {
		t.Error("expected GetMonitor to return singleton instance")
	}
}

func TestMonitor_SetMetricPath(t *testing.T) {
	m := prometheus.GetMonitor()
	path := "/custom-metrics"
	m.SetMetricPath(path)
}

func TestMonitor_SetSlowTime(t *testing.T) {
	m := prometheus.GetMonitor()
	slowTime := int32(10)
	m.SetSlowTime(slowTime)
}

func TestMonitor_SetExcludePaths(t *testing.T) {
	m := prometheus.GetMonitor()
	paths := []string{"/ping", "/health"}
	m.SetExcludePaths(paths)
}

func TestMonitor_SetDuration(t *testing.T) {
	m := prometheus.GetMonitor()
	duration := []float64{1.0, 2.0, 3.0}
	m.SetDuration(duration)
}

func TestMonitor_AddMetric_Counter(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Counter,
		Name:        "test_counter",
		Description: "Test counter metric",
		Labels:      []string{"label1"},
	}

	err := m.AddMetric(metric)
	if err != nil {
		t.Fatalf("failed to add counter metric: %v", err)
	}
}

func TestMonitor_AddMetric_Gauge(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Gauge,
		Name:        "test_gauge",
		Description: "Test gauge metric",
		Labels:      []string{"label1"},
	}

	err := m.AddMetric(metric)
	if err != nil {
		t.Fatalf("failed to add gauge metric: %v", err)
	}
}

func TestMonitor_AddMetric_Histogram(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Histogram,
		Name:        "test_histogram",
		Description: "Test histogram metric",
		Labels:      []string{"label1"},
		Buckets:     []float64{1, 2, 3},
	}

	err := m.AddMetric(metric)
	if err != nil {
		t.Fatalf("failed to add histogram metric: %v", err)
	}
}

func TestMonitor_AddMetric_HistogramNoBuckets(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Histogram,
		Name:        "test_histogram_no_buckets",
		Description: "Test histogram metric without buckets",
		Labels:      []string{"label1"},
	}

	err := m.AddMetric(metric)
	if err == nil {
		t.Error("expected error when adding histogram without buckets")
	}
}

func TestMonitor_AddMetric_Summary(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Summary,
		Name:        "test_summary",
		Description: "Test summary metric",
		Labels:      []string{"label1"},
		Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01},
	}

	err := m.AddMetric(metric)
	if err != nil {
		t.Fatalf("failed to add summary metric: %v", err)
	}
}

func TestMonitor_AddMetric_Duplicate(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Counter,
		Name:        "test_duplicate",
		Description: "Test duplicate metric",
		Labels:      []string{"label1"},
	}

	err := m.AddMetric(metric)
	if err != nil {
		t.Fatalf("failed to add metric: %v", err)
	}

	err = m.AddMetric(metric)
	if err == nil {
		t.Error("expected error when adding duplicate metric")
	}
}

func TestMonitor_AddMetric_EmptyName(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Counter,
		Name:        "",
		Description: "Test empty name",
		Labels:      []string{"label1"},
	}

	err := m.AddMetric(metric)
	if err == nil {
		t.Error("expected error when adding metric with empty name")
	}
}

func TestMonitor_GetMetric(t *testing.T) {
	prom.DefaultRegisterer = prom.NewRegistry()
	m := prometheus.GetMonitor()

	metric := &prometheus.Metric{
		Type:        prometheus.Counter,
		Name:        "test_get_metric",
		Description: "Test get metric",
		Labels:      []string{"label1"},
	}

	_ = m.AddMetric(metric)

	retrieved := m.GetMetric("test_get_metric")
	if retrieved == nil {
		t.Fatal("failed to retrieve metric")
	}
}

func TestMonitor_GetMetric_NotFound(t *testing.T) {
	m := prometheus.GetMonitor()

	retrieved := m.GetMetric("non_existent")
	if retrieved == nil {
		t.Fatal("expected empty metric to be returned")
	}
}
