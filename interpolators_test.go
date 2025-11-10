package interpolators

import (
	"math"
	"testing"
)

func TestInterpolateNone(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
	}{
		{
			name:  "empty input",
			input: []float64{},
		},
		{
			name:  "single element",
			input: []float64{1.5},
		},
		{
			name:  "multiple elements",
			input: []float64{1.0, 2.5, 3.7, 4.2, 5.9},
		},
		{
			name:  "negative values",
			input: []float64{-1.5, -2.3, -0.5},
		},
		{
			name:  "mixed values",
			input: []float64{-1.5, 0.0, 2.3, -3.7, 4.2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, None)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			// Check that output length matches input length
			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}

			// Check that each element matches
			for i := range tt.input {
				if out[i] != tt.input[i] {
					t.Errorf("Interpolate() output[%d] = %v, want %v", i, out[i], tt.input[i])
				}
			}

			// Verify that modifying output doesn't affect input (separate copy)
			if len(out) > 0 {
				original := tt.input[0]
				out[0] = 999.999
				if len(tt.input) > 0 && tt.input[0] != original {
					t.Errorf("Interpolate() modified input array, input should be unchanged")
				}
			}
		})
	}
}

func TestInterpolateDropSample(t *testing.T) {
	tests := []struct {
		name     string
		input    []float64
		expected []float64
	}{
		{
			name:     "empty input",
			input:    []float64{},
			expected: []float64{},
		},
		{
			name:     "single element",
			input:    []float64{1.0},
			expected: []float64{1.0},
		},
		{
			name:     "two elements",
			input:    []float64{1.0, 2.0},
			expected: []float64{1.0, 2.0},
		},
		{
			name:     "multiple elements",
			input:    []float64{1.0, 2.0, 3.0, 4.0},
			expected: []float64{1.0, 2.0, 3.0, 4.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, DropSample)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.expected) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.expected))
			}

			for i := range tt.expected {
				if math.Abs(out[i]-tt.expected[i]) > 1e-10 {
					t.Errorf("Interpolate() output[%d] = %v, want %v", i, out[i], tt.expected[i])
				}
			}
		})
	}
}

func TestInterpolateLinear(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
	}{
		{
			name:  "empty input",
			input: []float64{},
		},
		{
			name:  "single element",
			input: []float64{1.0},
		},
		{
			name:  "two elements",
			input: []float64{1.0, 2.0},
		},
		{
			name:  "multiple elements",
			input: []float64{1.0, 2.0, 3.0, 4.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, Linear)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateBSpline3(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
	}{
		{
			name:  "empty input",
			input: []float64{},
		},
		{
			name:  "single element",
			input: []float64{1.0},
		},
		{
			name:  "multiple elements",
			input: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, BSpline3)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateBSpline5(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
	}{
		{
			name:  "empty input",
			input: []float64{},
		},
		{
			name:  "single element",
			input: []float64{1.0},
		},
		{
			name:  "multiple elements",
			input: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, BSpline5)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

// Test impulse response functions directly
func TestDropSampleImpulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0},
		{0.5, 1.0},
		{0.99, 1.0},
		{1.0, 0.0},
		{1.5, 0.0},
		{-0.5, 1.0},
		{-1.0, 0.0},
	}

	for _, tt := range tests {
		result := dropSampleImpulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("dropSampleImpulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestLinearImpulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0},
		{0.5, 0.5},
		{1.0, 0.0},
		{1.5, 0.0},
		{-0.5, 0.5},
		{-1.0, 0.0},
	}

	for _, tt := range tests {
		result := linearImpulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("linearImpulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestBSpline3Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 2.0 / 3.0},
		{2.0, 0.0},
		{2.5, 0.0},
	}

	for _, tt := range tests {
		result := bspline3Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("bspline3Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestBSpline5Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 11.0 / 20.0},
		{3.0, 0.0},
		{3.5, 0.0},
	}

	for _, tt := range tests {
		result := bspline5Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("bspline5Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}
