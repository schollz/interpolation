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
			out, err := Interpolate(tt.input, len(tt.input), None)
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
			out, err := Interpolate(tt.input, len(tt.input), DropSample)
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
			out, err := Interpolate(tt.input, len(tt.input), Linear)
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
			out, err := Interpolate(tt.input, len(tt.input), BSpline3)
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
			out, err := Interpolate(tt.input, len(tt.input), BSpline5)
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

func TestLagrange4Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0},
		{0.5, 1.0 - 0.5*0.5 - 0.25 + 0.5*0.125},
		{1.0, 1.0 - 11.0/6.0 + 1.0 - 1.0/6.0},
		{2.0, 0.0},
		{2.5, 0.0},
		{-0.5, 1.0 - 0.5*0.5 - 0.25 + 0.5*0.125},
		{-2.0, 0.0},
	}

	for _, tt := range tests {
		result := lagrange4Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("lagrange4Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestLagrange6Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0},
		{3.0, 0.0},
		{3.5, 0.0},
		{-3.0, 0.0},
	}

	for _, tt := range tests {
		result := lagrange6Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("lagrange6Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestWatteImpulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0},
		{0.5, 1.0 - 0.5*0.5 - 0.5*0.25},
		{1.0, 1.0 - 1.5*1.0 + 0.5*1.0},
		{1.5, 1.0 - 1.5*1.5 + 0.5*2.25},
		{2.0, 0.0},
		{2.5, 0.0},
		{-0.5, 1.0 - 0.5*0.5 - 0.5*0.25},
		{-2.0, 0.0},
	}

	for _, tt := range tests {
		result := watteImpulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("watteImpulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestParabolic2xImpulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 0.5},
		{0.5, 0.5 - 0.25*0.25},
		{1.0, 1.0 - 1.0 + 0.25*1.0},
		{1.5, 1.0 - 1.5 + 0.25*2.25},
		{2.0, 0.0},
		{2.5, 0.0},
		{-0.5, 0.5 - 0.25*0.25},
		{-2.0, 0.0},
	}

	for _, tt := range tests {
		result := parabolic2xImpulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("parabolic2xImpulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestOsculating4Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0}, // At x=0: 1 - 0 - 0 + 0 - 0 = 1
		{1.0, -4.0 + 18.0 - 29.0 + 21.5 - 7.5 + 1.0}, // At x=1, should be continuous
		{2.0, 0.0}, // At x=2 and beyond, should be 0
		{2.5, 0.0},
		{-1.0, -4.0 + 18.0 - 29.0 + 21.5 - 7.5 + 1.0}, // Symmetric
		{-2.0, 0.0},
	}

	for _, tt := range tests {
		result := osculating4Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("osculating4Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

func TestOsculating6Impulse(t *testing.T) {
	tests := []struct {
		x        float64
		expected float64
	}{
		{0.0, 1.0}, // At x=0: 1 - 0 - 0 + 0 - 0 = 1
		{3.0, 0.0}, // At x=3 and beyond, should be 0
		{3.5, 0.0},
		{-3.0, 0.0}, // Symmetric
	}

	for _, tt := range tests {
		result := osculating6Impulse(tt.x)
		if math.Abs(result-tt.expected) > 1e-10 {
			t.Errorf("osculating6Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
		}
	}
}

// Test resampling with different output sample counts
func TestInterpolateResampling(t *testing.T) {
	tests := []struct {
		name             string
		input            []float64
		outSamples       int
		interpolatorType InterpolatorType
	}{
		{
			name:             "upsample with linear",
			input:            []float64{0.0, 1.0, 2.0, 3.0},
			outSamples:       7,
			interpolatorType: Linear,
		},
		{
			name:             "downsample with linear",
			input:            []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0},
			outSamples:       3,
			interpolatorType: Linear,
		},
		{
			name:             "upsample with bspline3",
			input:            []float64{1.0, 2.0, 3.0},
			outSamples:       5,
			interpolatorType: BSpline3,
		},
		{
			name:             "downsample with bspline5",
			input:            []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			outSamples:       3,
			interpolatorType: BSpline5,
		},
		{
			name:             "same size with drop sample",
			input:            []float64{1.0, 2.0, 3.0},
			outSamples:       3,
			interpolatorType: DropSample,
		},
		{
			name:             "upsample with lagrange4",
			input:            []float64{1.0, 2.0, 3.0, 4.0},
			outSamples:       7,
			interpolatorType: Lagrange4,
		},
		{
			name:             "upsample with lagrange6",
			input:            []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			outSamples:       10,
			interpolatorType: Lagrange6,
		},
		{
			name:             "upsample with watte",
			input:            []float64{1.0, 2.0, 3.0, 4.0},
			outSamples:       7,
			interpolatorType: Watte,
		},
		{
			name:             "upsample with parabolic2x",
			input:            []float64{1.0, 2.0, 3.0, 4.0},
			outSamples:       7,
			interpolatorType: Parabolic2x,
		},
		{
			name:             "upsample with osculating4",
			input:            []float64{1.0, 2.0, 3.0, 4.0},
			outSamples:       7,
			interpolatorType: Osculating4,
		},
		{
			name:             "upsample with osculating6",
			input:            []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			outSamples:       10,
			interpolatorType: Osculating6,
		},
		{
			name:             "upsample with hermite4",
			input:            []float64{1.0, 2.0, 3.0, 4.0},
			outSamples:       7,
			interpolatorType: Hermite4,
		},
		{
			name:             "upsample with hermite6_3",
			input:            []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			outSamples:       10,
			interpolatorType: Hermite6_3,
		},
		{
			name:             "upsample with hermite6_5",
			input:            []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0},
			outSamples:       10,
			interpolatorType: Hermite6_5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, tt.outSamples, tt.interpolatorType)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != tt.outSamples {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), tt.outSamples)
			}

			// For debugging, print the output
			t.Logf("Input: %v, Output: %v", tt.input, out)
		})
	}
}

// Test that upsampling produces expected values for linear interpolation
func TestInterpolateLinearUpsampling(t *testing.T) {
	input := []float64{0.0, 2.0, 4.0}
	outSamples := 5
	out, err := Interpolate(input, outSamples, Linear)
	if err != nil {
		t.Fatalf("Interpolate() returned unexpected error: %v", err)
	}

	if len(out) != outSamples {
		t.Fatalf("Interpolate() output length = %d, want %d", len(out), outSamples)
	}

	// First and last values should match input endpoints
	if math.Abs(out[0]-input[0]) > 1e-10 {
		t.Errorf("First output value = %v, want %v", out[0], input[0])
	}
	if math.Abs(out[outSamples-1]-input[len(input)-1]) > 1e-10 {
		t.Errorf("Last output value = %v, want %v", out[outSamples-1], input[len(input)-1])
	}

	// Output should be monotonically increasing for monotonic input
	for i := 1; i < len(out); i++ {
		if out[i] < out[i-1] {
			t.Errorf("Output not monotonically increasing at index %d: %v < %v", i, out[i], out[i-1])
		}
	}

	t.Logf("Input: %v, Output: %v", input, out)
}

// Test edge cases
func TestInterpolateEdgeCases(t *testing.T) {
	tests := []struct {
		name             string
		input            []float64
		outSamples       int
		interpolatorType InterpolatorType
	}{
		{
			name:             "single output sample",
			input:            []float64{1.0, 2.0, 3.0},
			outSamples:       1,
			interpolatorType: Linear,
		},
		{
			name:             "empty input",
			input:            []float64{},
			outSamples:       5,
			interpolatorType: Linear,
		},
		{
			name:             "single input sample to multiple outputs",
			input:            []float64{5.0},
			outSamples:       3,
			interpolatorType: Linear,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, tt.outSamples, tt.interpolatorType)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			expectedLen := tt.outSamples
			if len(tt.input) == 0 {
				expectedLen = 0
			}

			if len(out) != expectedLen {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), expectedLen)
			}

			t.Logf("Input: %v, Output: %v", tt.input, out)
		})
	}
}

func TestInterpolateLagrange4(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Lagrange4)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateLagrange6(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Lagrange6)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateWatte(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Watte)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateParabolic2x(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Parabolic2x)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateOsculating4(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Osculating4)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateOsculating6(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Osculating6)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestHermite4Impulse(t *testing.T) {
tests := []struct {
x        float64
expected float64
}{
{0.0, 1.0},                           // At x=0: 1 - 0 + 0 = 1
{0.5, 1.0 - 2.5*0.25 + 1.5*0.125},    // At x=0.5
{1.0, 2.0 - 4.0 + 2.5 - 0.5},         // At x=1, should be continuous
{1.5, 2.0 - 6.0 + 2.5*2.25 - 0.5*3.375}, // At x=1.5
{2.0, 0.0},                           // At x=2 and beyond, should be 0
{2.5, 0.0},
{-0.5, 1.0 - 2.5*0.25 + 1.5*0.125},   // Symmetric
{-2.0, 0.0},
}

for _, tt := range tests {
result := hermite4Impulse(tt.x)
if math.Abs(result-tt.expected) > 1e-10 {
t.Errorf("hermite4Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
}
}
}

func TestHermite6_3Impulse(t *testing.T) {
tests := []struct {
x        float64
expected float64
}{
{0.0, 1.0},  // At x=0: 1 - 0 + 0 = 1
{3.0, 0.0},  // At x=3 and beyond, should be 0
{3.5, 0.0},
{-3.0, 0.0}, // Symmetric
}

for _, tt := range tests {
result := hermite6_3Impulse(tt.x)
if math.Abs(result-tt.expected) > 1e-10 {
t.Errorf("hermite6_3Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
}
}
}

func TestHermite6_5Impulse(t *testing.T) {
tests := []struct {
x        float64
expected float64
}{
{0.0, 1.0},  // At x=0: 1 - 0 + 0 + 0 - 0 = 1
{3.0, 0.0},  // At x=3 and beyond, should be 0
{3.5, 0.0},
{-3.0, 0.0}, // Symmetric
}

for _, tt := range tests {
result := hermite6_5Impulse(tt.x)
if math.Abs(result-tt.expected) > 1e-10 {
t.Errorf("hermite6_5Impulse(%v) = %v, want %v", tt.x, result, tt.expected)
}
}
}

func TestInterpolateHermite4(t *testing.T) {
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
out, err := Interpolate(tt.input, len(tt.input), Hermite4)
if err != nil {
t.Errorf("Interpolate() returned unexpected error: %v", err)
}

if len(out) != len(tt.input) {
t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
}
})
}
}

func TestInterpolateHermite6_3(t *testing.T) {
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
out, err := Interpolate(tt.input, len(tt.input), Hermite6_3)
if err != nil {
t.Errorf("Interpolate() returned unexpected error: %v", err)
}

if len(out) != len(tt.input) {
t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
}
})
}
}

func TestInterpolateHermite6_5(t *testing.T) {
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
out, err := Interpolate(tt.input, len(tt.input), Hermite6_5)
if err != nil {
t.Errorf("Interpolate() returned unexpected error: %v", err)
}

if len(out) != len(tt.input) {
t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
}
})
}
}

func TestInterpolateCubicSpline(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), CubicSpline)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateMonotonicCubic(t *testing.T) {
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
			name:  "monotonic increasing",
			input: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
		},
		{
			name:  "monotonic decreasing",
			input: []float64{5.0, 4.0, 3.0, 2.0, 1.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := Interpolate(tt.input, len(tt.input)*2, MonotonicCubic)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input)*2 {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input)*2)
			}
		})
	}
}

func TestInterpolateLanczos2(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Lanczos2)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateLanczos3(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Lanczos3)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateBezier(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Bezier)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

func TestInterpolateAkima(t *testing.T) {
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
			out, err := Interpolate(tt.input, len(tt.input), Akima)
			if err != nil {
				t.Errorf("Interpolate() returned unexpected error: %v", err)
			}

			if len(out) != len(tt.input) {
				t.Errorf("Interpolate() output length = %d, want %d", len(out), len(tt.input))
			}
		})
	}
}

// BenchmarkInterpolators benchmarks all interpolator types with 1000 input points to 500 output points
func BenchmarkInterpolators(b *testing.B) {
	// Generate 1000 random input points
	input := make([]float64, 1000)
	for i := range input {
		input[i] = math.Sin(float64(i) * 0.1)
	}
	
	outSamples := 500
	
	benchmarks := []struct {
		name             string
		interpolatorType InterpolatorType
	}{
		{"DropSample", DropSample},
		{"Linear", Linear},
		{"BSpline3", BSpline3},
		{"BSpline5", BSpline5},
		{"Lagrange4", Lagrange4},
		{"Lagrange6", Lagrange6},
		{"Watte", Watte},
		{"Parabolic2x", Parabolic2x},
		{"Osculating4", Osculating4},
		{"Osculating6", Osculating6},
		{"Hermite4", Hermite4},
		{"Hermite6_3", Hermite6_3},
		{"Hermite6_5", Hermite6_5},
		{"CubicSpline", CubicSpline},
		{"MonotonicCubic", MonotonicCubic},
		{"Lanczos2", Lanczos2},
		{"Lanczos3", Lanczos3},
		{"Bezier", Bezier},
		{"Akima", Akima},
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := Interpolate(input, outSamples, bm.interpolatorType)
				if err != nil {
					b.Fatalf("Interpolate() returned unexpected error: %v", err)
				}
			}
		})
	}
}

func TestInterpolateInt(t *testing.T) {
	tests := []struct {
		name             string
		input            []int
		outSamples       int
		interpolatorType InterpolatorType
	}{
		{
			name:             "empty input",
			input:            []int{},
			outSamples:       5,
			interpolatorType: Linear,
		},
		{
			name:             "single element",
			input:            []int{42},
			outSamples:       3,
			interpolatorType: Linear,
		},
		{
			name:             "two elements linear",
			input:            []int{0, 10},
			outSamples:       5,
			interpolatorType: Linear,
		},
		{
			name:             "multiple elements linear",
			input:            []int{0, 10, 20, 30},
			outSamples:       7,
			interpolatorType: Linear,
		},
		{
			name:             "negative values",
			input:            []int{-10, -5, 0, 5, 10},
			outSamples:       9,
			interpolatorType: Linear,
		},
		{
			name:             "upsample with bspline3",
			input:            []int{1, 2, 3, 4, 5},
			outSamples:       9,
			interpolatorType: BSpline3,
		},
		{
			name:             "downsample with linear",
			input:            []int{0, 10, 20, 30, 40, 50},
			outSamples:       3,
			interpolatorType: Linear,
		},
		{
			name:             "none type",
			input:            []int{1, 2, 3, 4, 5},
			outSamples:       5,
			interpolatorType: None,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := InterpolateInt(tt.input, tt.outSamples, tt.interpolatorType)
			if err != nil {
				t.Errorf("InterpolateInt() returned unexpected error: %v", err)
			}

			expectedLen := tt.outSamples
			if len(tt.input) == 0 {
				expectedLen = 0
			}

			if len(out) != expectedLen {
				t.Errorf("InterpolateInt() output length = %d, want %d", len(out), expectedLen)
			}

			// Verify output is []int
			for i, v := range out {
				if v != int(v) {
					t.Errorf("InterpolateInt() output[%d] = %v is not an integer", i, v)
				}
			}

			t.Logf("Input: %v, Output: %v", tt.input, out)
		})
	}
}

func TestInterpolateIntLinearUpsampling(t *testing.T) {
	input := []int{0, 10, 20}
	outSamples := 5
	out, err := InterpolateInt(input, outSamples, Linear)
	if err != nil {
		t.Fatalf("InterpolateInt() returned unexpected error: %v", err)
	}

	if len(out) != outSamples {
		t.Fatalf("InterpolateInt() output length = %d, want %d", len(out), outSamples)
	}

	// First and last values should match input endpoints
	if out[0] != input[0] {
		t.Errorf("First output value = %v, want %v", out[0], input[0])
	}
	if out[outSamples-1] != input[len(input)-1] {
		t.Errorf("Last output value = %v, want %v", out[outSamples-1], input[len(input)-1])
	}

	// Output should be monotonically increasing for monotonic input
	for i := 1; i < len(out); i++ {
		if out[i] < out[i-1] {
			t.Errorf("Output not monotonically increasing at index %d: %v < %v", i, out[i], out[i-1])
		}
	}

	t.Logf("Input: %v, Output: %v", input, out)
}

func TestInterpolateIntRounding(t *testing.T) {
	// Test that rounding is done correctly
	tests := []struct {
		name             string
		input            []int
		outSamples       int
		interpolatorType InterpolatorType
	}{
		{
			name:             "positive values with rounding",
			input:            []int{0, 3, 6},
			outSamples:       5,
			interpolatorType: Linear,
		},
		{
			name:             "negative values with rounding",
			input:            []int{-6, -3, 0},
			outSamples:       5,
			interpolatorType: Linear,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := InterpolateInt(tt.input, tt.outSamples, tt.interpolatorType)
			if err != nil {
				t.Fatalf("InterpolateInt() returned unexpected error: %v", err)
			}

			// Verify all outputs are integers
			for i, v := range out {
				if v != int(v) {
					t.Errorf("InterpolateInt() output[%d] = %v is not an integer", i, v)
				}
			}

			t.Logf("Input: %v, Output: %v", tt.input, out)
		})
	}
}

func TestInterpolateIntAllTypes(t *testing.T) {
	input := []int{1, 5, 10, 15, 20, 25}
	outSamples := 10

	interpolationTypes := []struct {
		name string
		typ  InterpolatorType
	}{
		{"None", None},
		{"DropSample", DropSample},
		{"Linear", Linear},
		{"BSpline3", BSpline3},
		{"BSpline5", BSpline5},
		{"Lagrange4", Lagrange4},
		{"Lagrange6", Lagrange6},
		{"Watte", Watte},
		{"Parabolic2x", Parabolic2x},
		{"Osculating4", Osculating4},
		{"Osculating6", Osculating6},
		{"Hermite4", Hermite4},
		{"Hermite6_3", Hermite6_3},
		{"Hermite6_5", Hermite6_5},
		{"CubicSpline", CubicSpline},
		{"MonotonicCubic", MonotonicCubic},
		{"Lanczos2", Lanczos2},
		{"Lanczos3", Lanczos3},
		{"Bezier", Bezier},
		{"Akima", Akima},
	}

	for _, interp := range interpolationTypes {
		t.Run(interp.name, func(t *testing.T) {
			expectedLen := outSamples
			if interp.typ == None {
				expectedLen = len(input)
			}

			out, err := InterpolateInt(input, outSamples, interp.typ)
			if err != nil {
				t.Errorf("InterpolateInt() returned unexpected error: %v", err)
			}

			if len(out) != expectedLen {
				t.Errorf("InterpolateInt() output length = %d, want %d", len(out), expectedLen)
			}

			// Verify all outputs are integers
			for i, v := range out {
				if v != int(v) {
					t.Errorf("InterpolateInt() output[%d] = %v is not an integer", i, v)
				}
			}

			t.Logf("%s - Input: %v, Output: %v", interp.name, input, out)
		})
	}
}

// BenchmarkInterpolateInt benchmarks InterpolateInt vs manual conversion
func BenchmarkInterpolateInt(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i * 10
	}
	outSamples := 500

	b.Run("InterpolateInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := InterpolateInt(input, outSamples, Linear)
			if err != nil {
				b.Fatalf("InterpolateInt() returned unexpected error: %v", err)
			}
		}
	})

	b.Run("ManualConversion", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Convert to float64
			inFloat := make([]float64, len(input))
			for j, v := range input {
				inFloat[j] = float64(v)
			}

			// Interpolate
			outFloat, err := Interpolate(inFloat, outSamples, Linear)
			if err != nil {
				b.Fatalf("Interpolate() returned unexpected error: %v", err)
			}

			// Convert back to int
			out := make([]int, len(outFloat))
			for j, v := range outFloat {
				if v >= 0 {
					out[j] = int(v + 0.5)
				} else {
					out[j] = int(v - 0.5)
				}
			}
		}
	})
}
