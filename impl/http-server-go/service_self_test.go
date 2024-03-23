package main

import "testing"

func TestBuildCPUMetric(t *testing.T) {
	// Given
	used := uint64(23)
	total := uint64(1201)

	// When
	metric := buildCPUMetric(used, total)
	value := *metric.Value.Get()

	// Then
	expectedValue := 1.92
	if value != expectedValue {
		t.Fatalf("Expected value to be %f, actual %f", expectedValue, value)
	}
}

func TestBuildMemoryMetric(t *testing.T) {
	// Given
	used := uint64(14333321216)
	total := uint64(17176596480)

	// When
	metric := buildMemoryMetric(used, total)
	ratio := *metric.Ratio.Get()
	thresholds := metric.Thresholds
	value := *metric.Value.Get()

	// Then
	expectedRatio := 0.83
	if ratio != expectedRatio {
		t.Fatalf("Expected ratio to be %f, actual %f", expectedRatio, ratio)
	}
	expectedThresholds := []float64{1288.24, 1460.01}
	if thresholds[0] != expectedThresholds[0] {
		t.Fatalf("Expected thresholds[0] to be %f, actual %f", expectedThresholds[0], thresholds[0])
	}
	if thresholds[1] != expectedThresholds[1] {
		t.Fatalf("Expected thresholds[1] to be %f, actual %f", expectedThresholds[1], thresholds[1])
	}
	expectedValue := 1433.33
	if value != expectedValue {
		t.Fatalf("Expected value to be %f, actual %f", expectedValue, value)
	}
}
