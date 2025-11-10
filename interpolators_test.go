package interpolators

import (
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
