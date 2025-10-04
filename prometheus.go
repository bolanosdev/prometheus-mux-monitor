package prometheus

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricType int

const (
	None MetricType = iota
	Counter
	Gauge
	Histogram
	Summary

	defaultMetricPath = "/metrics"
	defaultSlowTime   = int32(5)
)

var (
	defaultExcludePaths = []string{}
	defaultDuration     = []float64{0.1, 0.3, 1.2, 5, 10}
	defaultCPUUsage     = []float64{.5, 1, 3, 5, 10, 20, 25, 30, 40, 50, 60, 70, 80, 90, 100}
	defaultMemUsage     = []float64{.5, 1, 3, 5, 10, 20, 25, 30, 40, 50, 60, 70, 80, 90, 100}
	monitor             *Monitor

	promTypeHandler = map[MetricType]func(metric *Metric) error{
		Counter:   counterHandler,
		Gauge:     gaugeHandler,
		Histogram: histogramHandler,
		Summary:   summaryHandler,
	}
)

type Monitor struct {
	slowTime     int32
	metricPath   string
	excludePaths []string
	reqDuration  []float64
	cpuUsage     []float64
	memUsage     []float64
	metrics      map[string]*Metric
	metadata     map[string]string
}

func GetMonitor() *Monitor {
	if monitor == nil {
		monitor = &Monitor{
			metricPath:   defaultMetricPath,
			slowTime:     defaultSlowTime,
			excludePaths: defaultExcludePaths,
			reqDuration:  defaultDuration,
			cpuUsage:     defaultCPUUsage,
			memUsage:     defaultMemUsage,
			metrics:      make(map[string]*Metric),
			metadata:     make(map[string]string),
		}
	}
	return monitor
}

func (m *Monitor) GetMetric(name string) *Metric {
	if metric, ok := m.metrics[name]; ok {
		return metric
	}
	return &Metric{}
}

func (m *Monitor) SetMetricPath(path string) {
	m.metricPath = path
}

func (m *Monitor) SetExcludePaths(paths []string) {
	m.excludePaths = paths
}

func (m *Monitor) SetSlowTime(slowTime int32) {
	m.slowTime = slowTime
}

func (m *Monitor) SetDuration(duration []float64) {
	m.reqDuration = duration
}

func (m *Monitor) SetMetricPrefix(prefix string) {
	metricRequestTotal = prefix + metricRequestTotal
	metricRequestUVTotal = prefix + metricRequestUVTotal
	metricURIRequestTotal = prefix + metricURIRequestTotal
	metricRequestBody = prefix + metricRequestBody
	metricResponseBody = prefix + metricResponseBody
	metricRequestDuration = prefix + metricRequestDuration
	metricSlowRequest = prefix + metricSlowRequest
}

func (m *Monitor) SetMetricSuffix(suffix string) {
	metricRequestTotal += suffix
	metricRequestUVTotal += suffix
	metricURIRequestTotal += suffix
	metricRequestBody += suffix
	metricResponseBody += suffix
	metricRequestDuration += suffix
	metricSlowRequest += suffix
}

func (m *Monitor) AddMetric(metric *Metric) error {
	if _, ok := m.metrics[metric.Name]; ok {
		return errors.Errorf("metric '%s' is existed", metric.Name)
	}

	if metric.Name == "" {
		return errors.Errorf("metric name cannot be empty.")
	}
	if f, ok := promTypeHandler[metric.Type]; ok {
		if err := f(metric); err == nil {
			prometheus.MustRegister(metric.vec)
			m.metrics[metric.Name] = metric
			return nil
		}
	}
	return errors.Errorf("metric type '%d' not existed.", metric.Type)
}

func counterHandler(metric *Metric) error {
	metric.vec = prometheus.NewCounterVec(
		prometheus.CounterOpts{Name: metric.Name, Help: metric.Description},
		metric.Labels,
	)
	return nil
}

func gaugeHandler(metric *Metric) error {
	metric.vec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{Name: metric.Name, Help: metric.Description},
		metric.Labels,
	)
	return nil
}

func histogramHandler(metric *Metric) error {
	if len(metric.Buckets) == 0 {
		return errors.Errorf("metric '%s' is histogram type, cannot lose bucket param.", metric.Name)
	}
	metric.vec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: metric.Name, Help: metric.Description, Buckets: metric.Buckets},
		metric.Labels,
	)
	return nil
}

func summaryHandler(metric *Metric) error {
	if len(metric.Objectives) == 0 {
		return errors.Errorf("metric '%s' is summary type, cannot lose objectives param.", metric.Name)
	}
	metric.vec = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{Name: metric.Name, Help: metric.Description, Objectives: metric.Objectives},
		metric.Labels,
	)
	return nil
}
