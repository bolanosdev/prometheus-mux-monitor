# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2025-10-03

### Added
- Comprehensive test suite with 14 tests covering core functionality
- Makefile with common development tasks (test, build, coverage, lint, fmt, vet)
- Tests moved to separate `tests/` directory for better organization
- CHANGELOG.md to track version history

### Changed
- **BREAKING**: `GetMonitor()` no longer accepts pgxpool.Pool parameter
- `initMetrics()` now called automatically in Interceptor middleware
- Cleaned up code comments and formatting

### Removed
- **BREAKING**: Removed `Use()` middleware function
- **BREAKING**: Removed database pool functionality from Monitor struct
- **BREAKING**: Removed pgxpool and all database dependencies
- **BREAKING**: Removed pgx-related imports and functionality
- Removed unused comments throughout codebase

## [0.1.0] - 2024-XX-XX

### Added
- Initial release
- HTTP request monitoring and metrics collection
- Prometheus metrics integration
- CPU and memory metrics tracking
- Bloom filter for unique visitor tracking
- Support for custom metrics (Counter, Gauge, Histogram, Summary)
- Configurable slow request tracking
- Request duration histograms
