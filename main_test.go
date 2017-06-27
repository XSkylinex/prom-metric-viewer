package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHelp(t *testing.T) {
	goroutinesHelp := "# HELP go_goroutines Number of goroutines that currently exist."

	name, help := parseHelp(goroutinesHelp)
	assert.Equal(t, "go_goroutines", name)
	assert.Equal(t, "Number of goroutines that currently exist.", help)
}

func TestParseType(t *testing.T) {
	goroutinesType := "# TYPE go_goroutines gauge"

	name, types := parseType(goroutinesType)
	assert.Equal(t, "go_goroutines", name)
	assert.Equal(t, "gauge", types)
}

func TestParseMetric(t *testing.T) {
	lines := []string{
		`go_gc_duration_seconds 0.00019936000000000002`,
		`go_gc_duration_seconds{quantile="0"} 0.00019936000000000002`,
		`go_gc_duration_seconds{quantile="0.25"} 0.000281616`,
		`go_gc_duration_seconds{quantile="0.5"} 0.000290125`,
		`go_gc_duration_seconds{quantile="0.75"} 0.000317352`,
		`go_gc_duration_seconds{quantile="1"} 0.004849631`,
	}

	for _, line := range lines {
		assert.Equal(t, "go_gc_duration_seconds", parseName(line))
	}
}

func TestParseRawMetric(t *testing.T) {
	metric, raw := parseRawMetric(`process_cpu_seconds_total 0.06`)
	assert.Equal(t, "process_cpu_seconds_total", metric)
	assert.Equal(t, `process_cpu_seconds_total`, raw.Element)
	assert.Equal(t, 0.06, raw.Value)

	metric, raw = parseRawMetric(`node_filesystem_size{device="/dev/mapper/root",fstype="ext4",mountpoint="/rootfs"} 5.270837248e+10`)
	assert.Equal(t, "node_filesystem_size", metric)
	assert.Equal(t, `node_filesystem_size{device="/dev/mapper/root",fstype="ext4",mountpoint="/rootfs"}`, raw.Element)
	assert.Equal(t, 5.270837248e+10, raw.Value)

	metric, raw = parseRawMetric(`elasticsearch_filesystem_data_available_bytes{cluster="prod",host="10.10.1.1",mount="/srv/elasticsearch/data1 (/dev/sdb)",name="prod-1",path="/srv/elasticsearch/data1/nodes/0"} 2.83319881216e+11`)
	assert.Equal(t, "elasticsearch_filesystem_data_available_bytes", metric)
	assert.Equal(t, `elasticsearch_filesystem_data_available_bytes{cluster="prod",host="10.10.1.1",mount="/srv/elasticsearch/data1 (/dev/sdb)",name="prod-1",path="/srv/elasticsearch/data1/nodes/0"}`, raw.Element)
	assert.Equal(t, 2.83319881216e+11, raw.Value)
}
