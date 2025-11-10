package interpolators

import "math"

// InterpolatorType defines the type of interpolation to use
type InterpolatorType int

const (
	// None returns the input data as-is without any interpolation
	None InterpolatorType = iota
	// DropSample is the 0th-order B-spline (1-point)
	DropSample
	// Linear is the 1st-order B-spline (2-point)
	Linear
	// BSpline3 is the 3rd-order B-spline (4-point)
	BSpline3
	// BSpline5 is the 5th-order B-spline (6-point)
	BSpline5
)

// dropSampleImpulse implements the drop-sample (0th-order B-spline) impulse response
// f(x) = 1 for 0 <= x < 1, 0 otherwise
func dropSampleImpulse(x float64) float64 {
	absX := math.Abs(x)
	if absX >= 0 && absX < 1 {
		return 1.0
	}
	return 0.0
}

// linearImpulse implements the linear (1st-order B-spline) impulse response
// f(x) = 1 - |x| for 0 <= |x| < 1, 0 for |x| >= 1
func linearImpulse(x float64) float64 {
	absX := math.Abs(x)
	if absX >= 0 && absX < 1 {
		return 1.0 - absX
	}
	return 0.0
}

// bspline3Impulse implements the 3rd-order B-spline (4-point) impulse response
func bspline3Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		return 2.0/3.0 - x2 + 0.5*x3
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		return 4.0/3.0 - 2.0*absX + x2 - x3/6.0
	}
	return 0.0
}

// bspline5Impulse implements the 5th-order B-spline (6-point) impulse response
func bspline5Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 11.0/20.0 - 0.5*x2 + 0.25*x4 - x5/12.0
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 17.0/40.0 + 5.0*absX/8.0 - 7.0*x2/4.0 + 5.0*x3/4.0 - 3.0*x4/8.0 + x5/24.0
	} else if absX >= 2 && absX < 3 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 81.0/40.0 - 27.0*absX/8.0 + 9.0*x2/4.0 - 3.0*x3/4.0 + x4/8.0 - x5/120.0
	}
	return 0.0
}

// Interpolate performs interpolation on the input data based on the specified type
func Interpolate(in []float64, interpolatorType InterpolatorType) (out []float64, err error) {
	switch interpolatorType {
	case None:
		// None type returns input exactly as it was
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	case DropSample:
		return applyInterpolation(in, dropSampleImpulse), nil
	case Linear:
		return applyInterpolation(in, linearImpulse), nil
	case BSpline3:
		return applyInterpolation(in, bspline3Impulse), nil
	case BSpline5:
		return applyInterpolation(in, bspline5Impulse), nil
	default:
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	}
}

// applyInterpolation applies the given impulse response function to interpolate the input data
func applyInterpolation(in []float64, impulse func(float64) float64) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	out := make([]float64, len(in))

	for i := range out {
		sum := 0.0
		// Apply the impulse response convolution
		for j := range in {
			distance := float64(i - j)
			sum += in[j] * impulse(distance)
		}
		out[i] = sum
	}

	return out
}
