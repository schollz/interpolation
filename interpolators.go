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
	// Lagrange4 is the 4-point, 3rd-order Lagrange interpolator
	Lagrange4
	// Lagrange6 is the 6-point, 5th-order Lagrange interpolator
	Lagrange6
	// Watte is the 4-point, 2nd-order Watte tri-linear interpolator
	Watte
	// Parabolic2x is the 4-point, 2nd-order parabolic 2x interpolator
	Parabolic2x
	// Osculating4 is the 4-point, 5th-order 2nd-order-osculating interpolator
	Osculating4
	// Osculating6 is the 6-point, 5th-order 2nd-order-osculating interpolator
	Osculating6
	// Hermite4 is the 4-point, 3rd-order Hermite interpolator (Catmull-Rom spline)
	Hermite4
	// Hermite6_3 is the 6-point, 3rd-order Hermite interpolator
	Hermite6_3
	// Hermite6_5 is the 6-point, 5th-order Hermite interpolator
	Hermite6_5
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

// lagrange4Impulse implements the 4-point, 3rd-order Lagrange impulse response
func lagrange4Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		return 1.0 - 0.5*absX - x2 + 0.5*x3
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		return 1.0 - 11.0*absX/6.0 + x2 - x3/6.0
	}
	return 0.0
}

// lagrange6Impulse implements the 6-point, 5th-order Lagrange impulse response
func lagrange6Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - absX/3.0 - 5.0*x2/4.0 + 5.0*x3/12.0 + x4/4.0 - x5/12.0
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - 13.0*absX/12.0 - 5.0*x2/8.0 + 25.0*x3/24.0 - 3.0*x4/8.0 + x5/24.0
	} else if absX >= 2 && absX < 3 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - 137.0*absX/60.0 + 15.0*x2/8.0 - 17.0*x3/24.0 + x4/8.0 - x5/120.0
	}
	return 0.0
}

// watteImpulse implements the 4-point, 2nd-order Watte tri-linear impulse response
// Formula: f(x) = 1 - 1/2*x - 1/2*x² for 0 ≤ x < 1
//
//	       1 - 3/2*x + 1/2*x² for 1 ≤ x < 2
//	       0 for x ≥ 2
//	f(-x) otherwise (symmetric)
func watteImpulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		return 1.0 - 0.5*absX - 0.5*x2
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		return 1.0 - 1.5*absX + 0.5*x2
	}
	return 0.0
}

// parabolic2xImpulse implements the 4-point, 2nd-order parabolic 2x impulse response
// Formula: f(x) = 1/2 - 1/4*x² for 0 ≤ x < 1
//
//	       1 - x + 1/4*x² for 1 ≤ x < 2
//	       0 for x ≥ 2
//	f(-x) otherwise (symmetric)
func parabolic2xImpulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		return 0.5 - 0.25*x2
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		return 1.0 - absX + 0.25*x2
	}
	return 0.0
}

// osculating4Impulse implements the 4-point, 5th-order 2nd-order-osculating impulse response
// Formula: f(x) = 1 - x² - 9/2*x³ + 15/2*x⁴ - 3*x⁵ for 0 ≤ x < 1
//
//	       -4 + 18x - 29x² + 43/2*x³ - 15/2*x⁴ + x⁵ for 1 ≤ x < 2
//	       0 for x ≥ 2
//	f(-x) otherwise (symmetric)
func osculating4Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - x2 - 4.5*x3 + 7.5*x4 - 3.0*x5
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return -4.0 + 18.0*absX - 29.0*x2 + 21.5*x3 - 7.5*x4 + x5
	}
	return 0.0
}

// osculating6Impulse implements the 6-point, 5th-order 2nd-order-osculating impulse response
// Formula: f(x) = 1 - 5/4*x² - 35/12*x³ + 21/4*x⁴ - 25/12*x⁵ for 0 ≤ x < 1
//
//	       -4 + 75/4*x - 245/8*x² + 545/24*x³ - 63/8*x⁴ + 25/24*x⁵ for 1 ≤ x < 2
//	       18 - 153/4*x + 255/8*x² - 313/24*x³ + 21/8*x⁴ - 5/24*x⁵ for 2 ≤ x < 3
//	       0 for x ≥ 3
//	f(-x) otherwise (symmetric)
func osculating6Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - 1.25*x2 - (35.0/12.0)*x3 + 5.25*x4 - (25.0/12.0)*x5
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return -4.0 + 18.75*absX - 30.625*x2 + (545.0/24.0)*x3 - 7.875*x4 + (25.0/24.0)*x5
	} else if absX >= 2 && absX < 3 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 18.0 - 38.25*absX + 31.875*x2 - (313.0/24.0)*x3 + 2.625*x4 - (5.0/24.0)*x5
	}
	return 0.0
}

// hermite4Impulse implements the 4-point, 3rd-order Hermite impulse response
// Also known as the Catmull-Rom spline, or α = -1/2 case of cardinal splines
// Formula: f(x) = 1 - 5/2*x² + 3/2*x³ for 0 ≤ x < 1
//          2 - 4*x + 5/2*x² - 1/2*x³ for 1 ≤ x < 2
//          0 for x ≥ 2
//          f(-x) otherwise (symmetric)
func hermite4Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		return 1.0 - 2.5*x2 + 1.5*x3
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		return 2.0 - 4.0*absX + 2.5*x2 - 0.5*x3
	}
	return 0.0
}

// hermite6_3Impulse implements the 6-point, 3rd-order Hermite impulse response
// First derivative matches with the first derivatives of the Lagrangians
// Formula: f(x) = 1 - 7/3*x² + 4/3*x³ for 0 ≤ x < 1
//          5/2 - 59/12*x + 3*x² - 7/12*x³ for 1 ≤ x < 2
//          -3/2 + 7/4*x - 2/3*x² + 1/12*x³ for 2 ≤ x < 3
//          0 for x ≥ 3
//          f(-x) otherwise (symmetric)
func hermite6_3Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		return 1.0 - (7.0/3.0)*x2 + (4.0/3.0)*x3
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		return 2.5 - (59.0/12.0)*absX + 3.0*x2 - (7.0/12.0)*x3
	} else if absX >= 2 && absX < 3 {
		x2 := absX * absX
		x3 := x2 * absX
		return -1.5 + 1.75*absX - (2.0/3.0)*x2 + (1.0/12.0)*x3
	}
	return 0.0
}

// hermite6_5Impulse implements the 6-point, 5th-order Hermite impulse response
// Linear ramp between two Lagrangians
// Formula: f(x) = 1 - 25/12*x² + 5/12*x³ + 13/12*x⁴ - 5/12*x⁵ for 0 ≤ x < 1
//          1 + 5/12*x - 35/8*x² + 35/8*x³ - 13/8*x⁴ + 5/24*x⁵ for 1 ≤ x < 2
//          3 - 29/4*x + 155/24*x² - 65/24*x³ + 13/24*x⁴ - 1/24*x⁵ for 2 ≤ x < 3
//          0 for x ≥ 3
//          f(-x) otherwise (symmetric)
func hermite6_5Impulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 - (25.0/12.0)*x2 + (5.0/12.0)*x3 + (13.0/12.0)*x4 - (5.0/12.0)*x5
	} else if absX >= 1 && absX < 2 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 1.0 + (5.0/12.0)*absX - (35.0/8.0)*x2 + (35.0/8.0)*x3 - (13.0/8.0)*x4 + (5.0/24.0)*x5
	} else if absX >= 2 && absX < 3 {
		x2 := absX * absX
		x3 := x2 * absX
		x4 := x2 * x2
		x5 := x4 * absX
		return 3.0 - (29.0/4.0)*absX + (155.0/24.0)*x2 - (65.0/24.0)*x3 + (13.0/24.0)*x4 - (1.0/24.0)*x5
	}
	return 0.0
}

// Interpolate performs interpolation on the input data based on the specified type
func Interpolate(in []float64, outSamples int, interpolatorType InterpolatorType) (out []float64, err error) {
	switch interpolatorType {
	case None:
		// None type returns input exactly as it was
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	case DropSample:
		return applyInterpolation(in, outSamples, dropSampleImpulse), nil
	case Linear:
		return applyInterpolation(in, outSamples, linearImpulse), nil
	case BSpline3:
		return applyInterpolation(in, outSamples, bspline3Impulse), nil
	case BSpline5:
		return applyInterpolation(in, outSamples, bspline5Impulse), nil
	case Lagrange4:
		return applyInterpolation(in, outSamples, lagrange4Impulse), nil
	case Lagrange6:
		return applyInterpolation(in, outSamples, lagrange6Impulse), nil
	case Watte:
		return applyInterpolation(in, outSamples, watteImpulse), nil
	case Parabolic2x:
		return applyInterpolation(in, outSamples, parabolic2xImpulse), nil
	case Osculating4:
		return applyInterpolation(in, outSamples, osculating4Impulse), nil
	case Osculating6:
		return applyInterpolation(in, outSamples, osculating6Impulse), nil
	case Hermite4:
		return applyInterpolation(in, outSamples, hermite4Impulse), nil
	case Hermite6_3:
		return applyInterpolation(in, outSamples, hermite6_3Impulse), nil
	case Hermite6_5:
		return applyInterpolation(in, outSamples, hermite6_5Impulse), nil
	default:
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	}
}

// applyInterpolation applies the given impulse response function to interpolate the input data
func applyInterpolation(in []float64, outSamples int, impulse func(float64) float64) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	out := make([]float64, outSamples)

	// Calculate the ratio to map output samples to input samples
	var ratio float64
	if outSamples > 1 {
		ratio = float64(len(in)-1) / float64(outSamples-1)
	} else {
		ratio = 0
	}

	for i := range out {
		// Calculate the position in the input array
		pos := float64(i) * ratio
		sum := 0.0

		// Apply the impulse response convolution
		for j := range in {
			distance := pos - float64(j)
			sum += in[j] * impulse(distance)
		}
		out[i] = sum
	}

	return out
}
