# Prometheus Mux Monitor

A lightweight Prometheus monitoring middleware for Go HTTP applications using Gorilla Mux. Provides comprehensive metrics collection for HTTP requests, system resources (CPU/Memory), and custom metrics.

## Features

- 📊 **HTTP Request Metrics**: Request count, duration, body size, slow requests
- 💻 **System Metrics**: CPU usage (user, system, idle) and memory utilization
- 👥 **Unique Visitor Tracking**: Built-in Bloom filter for UV counting
- 🎯 **Custom Metrics**: Support for Counter, Gauge, Histogram, and Summary types
- ⚡ **Lightweight**: Minimal dependencies and overhead
- 🔧 **Configurable**: Customizable metric paths, slow request thresholds, and buckets

## Installation

```bash
go get github.com/bolanosdev/prometheus-mux-monitor
```

## Quick Start

```go
package main

import (
    "net/http"
    "github.com/gorilla/mux"
    prometheus "github.com/bolanosdev/prometheus-mux-monitor"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    r := mux.NewRouter()
    
    // Get the monitor instance
    monitor := prometheus.GetMonitor()
    
    // Apply the monitoring interceptor
    r.Use(monitor.Interceptor)
    
    // Expose metrics endpoint
    r.Handle("/metrics", promhttp.Handler())
    
    // Your routes
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("x-status-code", "200")
        w.Write([]byte("Hello World"))
    })
    
    http.ListenAndServe(":8080", r)
}
```

## Configuration

### Set Metric Path

```go
monitor := prometheus.GetMonitor()
monitor.SetMetricPath("/custom-metrics")
```

### Configure Slow Request Threshold

```go
monitor := prometheus.GetMonitor()
monitor.SetSlowTime(10) // Requests slower than 10 seconds
```

### Set Duration Buckets

```go
monitor := prometheus.GetMonitor()
monitor.SetDuration([]float64{0.1, 0.5, 1.0, 2.0, 5.0, 10.0})
```

### Exclude Paths from Monitoring

```go
monitor := prometheus.GetMonitor()
monitor.SetExcludePaths([]string{"/health", "/ping"})
```

### Add Metric Prefix/Suffix

```go
monitor := prometheus.GetMonitor()
monitor.SetMetricPrefix("myapp_")
monitor.SetMetricSuffix("_total")
```

## Built-in Metrics

### HTTP Metrics

- `request_total`: Total number of HTTP requests
- `request_uv`: Unique visitors (client IPs)
- `request_uv_total`: Total unique visitors count
- `uri_request_total`: Requests per URI, method, and status code
- `request_body_total`: Total request body bytes received
- `response_body_total`: Total response body bytes sent
- `request_duration`: Request duration histogram
- `slow_request_total`: Count of slow requests (above threshold)

### System Metrics

- `cpu_user_total`: CPU time consumed by user processes
- `cpu_system_total`: CPU time consumed by system processes
- `cpu_idle_total`: CPU idle time
- `mem_used_total`: Memory usage percentage
- `mem_cached_total`: Cached memory percentage

## Custom Metrics

### Counter

```go
monitor := prometheus.GetMonitor()
err := monitor.AddMetric(&prometheus.Metric{
    Type:        prometheus.Counter,
    Name:        "custom_counter",
    Description: "Description of counter",
    Labels:      []string{"label1", "label2"},
})

// Increment counter
metric := monitor.GetMetric("custom_counter")
metric.Inc([]string{"value1", "value2"})

// Add specific value
metric.Add([]string{"value1", "value2"}, 5.0)
```

### Gauge

```go
monitor.AddMetric(&prometheus.Metric{
    Type:        prometheus.Gauge,
    Name:        "custom_gauge",
    Description: "Description of gauge",
    Labels:      []string{"label1"},
})

metric := monitor.GetMetric("custom_gauge")
metric.SetGaugeValue([]string{"value1"}, 42.0)
```

### Histogram

```go
monitor.AddMetric(&prometheus.Metric{
    Type:        prometheus.Histogram,
    Name:        "custom_histogram",
    Description: "Description of histogram",
    Labels:      []string{"label1"},
    Buckets:     []float64{0.1, 0.5, 1.0, 5.0, 10.0},
})

metric := monitor.GetMetric("custom_histogram")
metric.Observe([]string{"value1"}, 2.5)
```

### Summary

```go
monitor.AddMetric(&prometheus.Metric{
    Type:        prometheus.Summary,
    Name:        "custom_summary",
    Description: "Description of summary",
    Labels:      []string{"label1"},
    Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
})

metric := monitor.GetMetric("custom_summary")
metric.Observe([]string{"value1"}, 3.7)
```

## Response Status Code

To properly track status codes, set the `x-status-code` header in your response:

```go
r.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("x-status-code", "200")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
})
```

## Development

### Prerequisites

- Go 1.23 or higher

### Build

```bash
make build
```

### Run Tests

```bash
make test
```

### Generate Coverage Report

```bash
make coverage
```

### Format Code

```bash
make fmt
```

### Run Linter

```bash
make lint
```

### Run All Checks

```bash
make all
```

## Makefile Targets

- `make test` - Run all tests
- `make build` - Build the project
- `make coverage` - Generate coverage report (HTML)
- `make fmt` - Format code with gofmt
- `make vet` - Run go vet
- `make lint` - Run golangci-lint
- `make clean` - Remove build artifacts
- `make install` - Install dependencies
- `make all` - Run fmt, vet, test, and build
- `make help` - Show available targets

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Changelog

See [CHANGELOG.md](CHANGELOG.md) for a list of changes.

## Credits

Built with:
- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router
- [Prometheus Go Client](https://github.com/prometheus/client_golang) - Prometheus instrumentation
- [go-osstat](https://github.com/mackerelio/go-osstat) - System statistics
