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
	// CubicSpline is the natural cubic spline interpolator with C² continuity
	CubicSpline
	// MonotonicCubic is the Fritsch-Carlson monotonic cubic interpolator (preserves monotonicity)
	MonotonicCubic
	// Lanczos2 is the windowed sinc interpolator with a=2 (4-point)
	Lanczos2
	// Lanczos3 is the windowed sinc interpolator with a=3 (6-point)
	Lanczos3
	// Bezier is the cubic Bezier curve interpolator
	Bezier
	// Akima is the Akima spline interpolator (robust to outliers)
	Akima
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

// cubicSplineCoefficients computes the coefficients for natural cubic spline
func cubicSplineCoefficients(x, y []float64) (a, b, c, d []float64) {
	n := len(x) - 1
	h := make([]float64, n)
	for i := 0; i < n; i++ {
		h[i] = x[i+1] - x[i]
	}

	// Solve tridiagonal system for second derivatives
	alpha := make([]float64, n)
	for i := 1; i < n; i++ {
		alpha[i] = (3/h[i])*(y[i+1]-y[i]) - (3/h[i-1])*(y[i]-y[i-1])
	}

	l := make([]float64, n+1)
	mu := make([]float64, n+1)
	z := make([]float64, n+1)
	l[0] = 1.0

	for i := 1; i < n; i++ {
		l[i] = 2*(x[i+1]-x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}

	l[n] = 1.0
	z[n] = 0.0
	c = make([]float64, n+1)
	b = make([]float64, n)
	d = make([]float64, n)
	a = make([]float64, n)

	c[n] = 0.0
	for j := n - 1; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (y[j+1]-y[j])/h[j] - h[j]*(c[j+1]+2*c[j])/3
		d[j] = (c[j+1] - c[j]) / (3 * h[j])
		a[j] = y[j]
	}

	return a, b, c, d
}

// monotonicCubicSlopes computes slopes for Fritsch-Carlson monotonic cubic interpolation
func monotonicCubicSlopes(x, y []float64) []float64 {
	n := len(x)
	delta := make([]float64, n-1)
	m := make([]float64, n)

	// Calculate secant slopes
	for i := 0; i < n-1; i++ {
		delta[i] = (y[i+1] - y[i]) / (x[i+1] - x[i])
	}

	// Initialize tangents
	m[0] = delta[0]
	for i := 1; i < n-1; i++ {
		if delta[i-1]*delta[i] <= 0 {
			m[i] = 0
		} else {
			m[i] = (delta[i-1] + delta[i]) / 2
		}
	}
	m[n-1] = delta[n-2]

	// Adjust tangents to preserve monotonicity
	for i := 0; i < n-1; i++ {
		if math.Abs(delta[i]) < 1e-10 {
			m[i] = 0
			m[i+1] = 0
		} else {
			alpha := m[i] / delta[i]
			beta := m[i+1] / delta[i]
			tau := 3.0
			if alpha*alpha+beta*beta > tau*tau {
				t := tau / math.Sqrt(alpha*alpha+beta*beta)
				m[i] = t * alpha * delta[i]
				m[i+1] = t * beta * delta[i]
			}
		}
	}

	return m
}

// lanczos2Impulse implements the Lanczos-2 windowed sinc impulse response
func lanczos2Impulse(x float64) float64 {
	absX := math.Abs(x)
	if absX < 1e-10 {
		return 1.0
	}
	if absX >= 2.0 {
		return 0.0
	}
	// sinc(x) * sinc(x/a) where a=2
	piX := math.Pi * absX
	return (math.Sin(piX) / piX) * (math.Sin(piX/2.0) / (piX / 2.0))
}

// lanczos3Impulse implements the Lanczos-3 windowed sinc impulse response
func lanczos3Impulse(x float64) float64 {
	absX := math.Abs(x)
	if absX < 1e-10 {
		return 1.0
	}
	if absX >= 3.0 {
		return 0.0
	}
	// sinc(x) * sinc(x/a) where a=3
	piX := math.Pi * absX
	return (math.Sin(piX) / piX) * (math.Sin(piX/3.0) / (piX / 3.0))
}

// bezierImpulse implements cubic Bezier curve interpolation
// Uses uniform parameterization with automatic control point generation
func bezierImpulse(x float64) float64 {
	absX := math.Abs(x)

	if absX >= 0 && absX < 1 {
		// Cubic Bezier basis function B1(t) for t in [0,1]
		t := absX
		t2 := t * t
		t3 := t2 * t
		// Smooth interpolation similar to smoothstep
		return 1.0 - 3.0*t2 + 2.0*t3
	} else if absX >= 1 && absX < 2 {
		// Smooth falloff in outer region
		t := 2.0 - absX
		return t * t * (3.0 - 2.0*t) / 8.0
	}
	return 0.0
}

// akimaSlopes computes slopes for Akima spline interpolation
func akimaSlopes(x, y []float64) []float64 {
	n := len(x)
	m := make([]float64, n)

	if n < 3 {
		// Not enough points for Akima, fall back to linear
		for i := 0; i < n-1; i++ {
			m[i] = (y[i+1] - y[i]) / (x[i+1] - x[i])
		}
		if n > 0 {
			m[n-1] = m[n-2]
		}
		return m
	}

	// Calculate segment slopes
	s := make([]float64, n+3)
	for i := 0; i < n-1; i++ {
		s[i+2] = (y[i+1] - y[i]) / (x[i+1] - x[i])
	}

	// Extrapolate slopes at boundaries
	s[1] = 2*s[2] - s[3]
	s[0] = 2*s[1] - s[2]
	s[n+1] = 2*s[n] - s[n-1]
	s[n+2] = 2*s[n+1] - s[n]

	// Calculate Akima slopes
	for i := 0; i < n; i++ {
		w1 := math.Abs(s[i+3] - s[i+2])
		w2 := math.Abs(s[i+1] - s[i])

		if w1+w2 < 1e-10 {
			m[i] = (s[i+1] + s[i+2]) / 2.0
		} else {
			m[i] = (w1*s[i+1] + w2*s[i+2]) / (w1 + w2)
		}
	}

	return m
}

// linearInterpolate performs optimized linear interpolation
// This specialized version only checks adjacent samples instead of all samples
func linearInterpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the two adjacent samples
		idx0 := int(pos)
		idx1 := idx0 + 1

		// Handle boundary cases
		if idx0 >= len(in)-1 {
			out[i] = in[len(in)-1]
			continue
		}

		// Linear interpolation between the two samples
		// distance from idx0 is the fractional part
		frac := pos - float64(idx0)

		// Linear interpolation: (1-frac)*val0 + frac*val1
		out[i] = in[idx0]*(1.0-frac) + in[idx1]*frac
	}

	return out
}

// dropSampleInterpolate performs optimized drop-sample (nearest neighbor) interpolation
// This specialized version picks the nearest input sample for each output sample
func dropSampleInterpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Round to nearest integer to get the closest sample
		idx := int(pos + 0.5)

		// Handle boundary cases
		if idx >= len(in) {
			idx = len(in) - 1
		}

		out[i] = in[idx]
	}

	return out
}

// bspline3Interpolate performs optimized B-spline 3 (cubic B-spline) interpolation
// This specialized version only checks 4 nearby samples (support ±2)
func bspline3Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 4 nearby samples (support is ±2)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 4 samples: centerIdx-1, centerIdx, centerIdx+1, centerIdx+2
		// This covers the range where |distance| < 2
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline bspline3 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x3 := x2 * absX
				impulse = 2.0/3.0 - x2 + 0.5*x3
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				impulse = 4.0/3.0 - 2.0*absX + x2 - x3/6.0
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// bspline5Interpolate performs optimized B-spline 5 (quintic B-spline) interpolation
// This specialized version only checks 6 nearby samples (support ±3)
func bspline5Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 6 nearby samples (support is ±3)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 6 samples: centerIdx-2 to centerIdx+3
		// This covers the range where |distance| < 3
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline bspline5 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 11.0/20.0 - 0.5*x2 + 0.25*x4 - x5/12.0
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 17.0/40.0 + 5.0*absX/8.0 - 7.0*x2/4.0 + 5.0*x3/4.0 - 3.0*x4/8.0 + x5/24.0
			} else if absX < 3 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 81.0/40.0 - 27.0*absX/8.0 + 9.0*x2/4.0 - 3.0*x3/4.0 + x4/8.0 - x5/120.0
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// lagrange4Interpolate performs optimized Lagrange 4-point interpolation
// This specialized version only checks 4 nearby samples (support ±2)
func lagrange4Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 4 nearby samples (support is ±2)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 4 samples: centerIdx-1, centerIdx, centerIdx+1, centerIdx+2
		// This covers the range where |distance| < 2
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline lagrange4 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x3 := x2 * absX
				impulse = 1.0 - 0.5*absX - x2 + 0.5*x3
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				impulse = 1.0 - 11.0*absX/6.0 + x2 - x3/6.0
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// lagrange6Interpolate performs optimized Lagrange 6-point interpolation
// This specialized version only checks 6 nearby samples (support ±3)
func lagrange6Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 6 nearby samples (support is ±3)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 6 samples: centerIdx-2 to centerIdx+3
		// This covers the range where |distance| < 3
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline lagrange6 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 1.0 - absX/3.0 - 5.0*x2/4.0 + 5.0*x3/12.0 + x4/4.0 - x5/12.0
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 1.0 - 13.0*absX/12.0 - 5.0*x2/8.0 + 25.0*x3/24.0 - 3.0*x4/8.0 + x5/24.0
			} else if absX < 3 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 1.0 - 137.0*absX/60.0 + 15.0*x2/8.0 - 17.0*x3/24.0 + x4/8.0 - x5/120.0
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// watteInterpolate performs optimized Watte tri-linear interpolation
// This specialized version only checks 4 nearby samples (support ±2)
func watteInterpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 4 nearby samples (support is ±2)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 4 samples: centerIdx-1, centerIdx, centerIdx+1, centerIdx+2
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline watte impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				impulse = 1.0 - 0.5*absX - 0.5*x2
			} else if absX < 2 {
				x2 := absX * absX
				impulse = 1.0 - 1.5*absX + 0.5*x2
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// parabolic2xInterpolate performs optimized parabolic 2x interpolation
// This specialized version only checks 4 nearby samples (support ±2)
func parabolic2xInterpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 4 nearby samples (support is ±2)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 4 samples: centerIdx-1, centerIdx, centerIdx+1, centerIdx+2
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline parabolic2x impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				impulse = 0.5 - 0.25*x2
			} else if absX < 2 {
				x2 := absX * absX
				impulse = 1.0 - absX + 0.25*x2
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// osculating4Interpolate performs optimized Osculating 4-point interpolation
// This specialized version only checks 4 nearby samples (support ±2)
func osculating4Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 4 nearby samples (support is ±2)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 4 samples: centerIdx-1, centerIdx, centerIdx+1, centerIdx+2
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline osculating4 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 1.0 - x2 - 4.5*x3 + 7.5*x4 - 3.0*x5
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = -4.0 + 18.0*absX - 29.0*x2 + 21.5*x3 - 7.5*x4 + x5
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// osculating6Interpolate performs optimized Osculating 6-point interpolation
// This specialized version only checks 6 nearby samples (support ±3)
func osculating6Interpolate(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}

	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
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

		// Get the 6 nearby samples (support is ±3)
		centerIdx := int(pos + 0.5) // Round to nearest
		sum := 0.0

		// Check 6 samples: centerIdx-2 to centerIdx+3
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			if j < 0 || j >= len(in) {
				continue
			}
			distance := pos - float64(j)
			absX := distance
			if absX < 0 {
				absX = -absX
			}

			// Inline osculating6 impulse calculation
			var impulse float64
			if absX < 1 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 1.0 - 1.25*x2 - (35.0/12.0)*x3 + 5.25*x4 - (25.0/12.0)*x5
			} else if absX < 2 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = -4.0 + 18.75*absX - 30.625*x2 + (545.0/24.0)*x3 - 7.875*x4 + (25.0/24.0)*x5
			} else if absX < 3 {
				x2 := absX * absX
				x3 := x2 * absX
				x4 := x2 * x2
				x5 := x4 * absX
				impulse = 18.0 - 38.25*absX + 31.875*x2 - (313.0/24.0)*x3 + 2.625*x4 - (5.0/24.0)*x5
			} else {
				impulse = 0.0
			}

			sum += in[j] * impulse
		}
		out[i] = sum
	}

	return out
}

// hermite4Interpolate implements optimized 4-point Hermite (Catmull-Rom) interpolation
// Support: ±2 (checks 4 samples per output)
func hermite4Interpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 4 samples: centerIdx-1 to centerIdx+2 (support ±2)
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline hermite4 impulse calculation
			var impulse float64
			if distance >= 0 && distance < 1 {
				x2 := distance * distance
				x3 := x2 * distance
				impulse = 1.0 - 2.5*x2 + 1.5*x3
			} else if distance >= 1 && distance < 2 {
				x2 := distance * distance
				x3 := x2 * distance
				impulse = 2.0 - 4.0*distance + 2.5*x2 - 0.5*x3
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
}

// hermite6_3Interpolate implements optimized 6-point, 3rd-order Hermite interpolation
// Support: ±3 (checks 6 samples per output)
func hermite6_3Interpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 6 samples: centerIdx-2 to centerIdx+3 (support ±3)
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline hermite6_3 impulse calculation
			var impulse float64
			if distance >= 0 && distance < 1 {
				x2 := distance * distance
				x3 := x2 * distance
				impulse = 1.0 - (7.0/3.0)*x2 + (4.0/3.0)*x3
			} else if distance >= 1 && distance < 2 {
				x2 := distance * distance
				x3 := x2 * distance
				impulse = 2.5 - (59.0/12.0)*distance + 3.0*x2 - (7.0/12.0)*x3
			} else if distance >= 2 && distance < 3 {
				x2 := distance * distance
				x3 := x2 * distance
				impulse = -1.5 + 1.75*distance - (2.0/3.0)*x2 + (1.0/12.0)*x3
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
}

// hermite6_5Interpolate implements optimized 6-point, 5th-order Hermite interpolation
// Support: ±3 (checks 6 samples per output)
func hermite6_5Interpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 6 samples: centerIdx-2 to centerIdx+3 (support ±3)
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline hermite6_5 impulse calculation
			var impulse float64
			if distance >= 0 && distance < 1 {
				x2 := distance * distance
				x3 := x2 * distance
				x4 := x2 * x2
				x5 := x4 * distance
				impulse = 1.0 - (25.0/12.0)*x2 + (5.0/12.0)*x3 + (13.0/12.0)*x4 - (5.0/12.0)*x5
			} else if distance >= 1 && distance < 2 {
				x2 := distance * distance
				x3 := x2 * distance
				x4 := x2 * x2
				x5 := x4 * distance
				impulse = 1.0 + (5.0/12.0)*distance - (35.0/8.0)*x2 + (35.0/8.0)*x3 - (13.0/8.0)*x4 + (5.0/24.0)*x5
			} else if distance >= 2 && distance < 3 {
				x2 := distance * distance
				x3 := x2 * distance
				x4 := x2 * x2
				x5 := x4 * distance
				impulse = 3.0 - (29.0/4.0)*distance + (155.0/24.0)*x2 - (65.0/24.0)*x3 + (13.0/24.0)*x4 - (1.0/24.0)*x5
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
}

// lanczos2Interpolate implements optimized Lanczos-2 windowed sinc interpolation
// Support: ±2 (checks 4 samples per output)
func lanczos2Interpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 4 samples: centerIdx-1 to centerIdx+2 (support ±2)
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline lanczos2 impulse calculation
			var impulse float64
			if distance < 1e-10 {
				impulse = 1.0
			} else if distance < 2.0 {
				// sinc(x) * sinc(x/a) where a=2
				piX := math.Pi * distance
				impulse = (math.Sin(piX) / piX) * (math.Sin(piX/2.0) / (piX / 2.0))
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
}

// lanczos3Interpolate implements optimized Lanczos-3 windowed sinc interpolation
// Support: ±3 (checks 6 samples per output)
func lanczos3Interpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 6 samples: centerIdx-2 to centerIdx+3 (support ±3)
		for j := centerIdx - 2; j <= centerIdx+3; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline lanczos3 impulse calculation
			var impulse float64
			if distance < 1e-10 {
				impulse = 1.0
			} else if distance < 3.0 {
				// sinc(x) * sinc(x/a) where a=3
				piX := math.Pi * distance
				impulse = (math.Sin(piX) / piX) * (math.Sin(piX/3.0) / (piX / 3.0))
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
}

// bezierInterpolate implements optimized cubic Bezier curve interpolation
// Support: ±2 (checks 4 samples per output)
func bezierInterpolate(in []float64, outSamples int) []float64 {
	out := make([]float64, outSamples)
	if len(in) == 0 {
		return out
	}
	if len(in) == 1 {
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	ratio := float64(len(in)-1) / float64(outSamples-1)
	lastIdx := len(in) - 1

	for i := range out {
		pos := float64(i) * ratio
		centerIdx := int(math.Round(pos))

		var sum float64
		// Check 4 samples: centerIdx-1 to centerIdx+2 (support ±2)
		for j := centerIdx - 1; j <= centerIdx+2; j++ {
			// Clamp to valid range
			idx := j
			if idx < 0 {
				idx = 0
			} else if idx > lastIdx {
				idx = lastIdx
			}

			distance := math.Abs(pos - float64(j))

			// Inline bezier impulse calculation
			var impulse float64
			if distance >= 0 && distance < 1 {
				// Cubic Bezier basis function B1(t) for t in [0,1]
				t := distance
				t2 := t * t
				t3 := t2 * t
				impulse = 1.0 - 3.0*t2 + 2.0*t3
			} else if distance >= 1 && distance < 2 {
				// Smooth falloff in outer region
				t := 2.0 - distance
				impulse = t * t * (3.0 - 2.0*t) / 8.0
			} else {
				impulse = 0.0
			}

			sum += in[idx] * impulse
		}
		out[i] = sum
	}

	return out
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
		return dropSampleInterpolate(in, outSamples), nil
	case Linear:
		return linearInterpolate(in, outSamples), nil
	case BSpline3:
		return bspline3Interpolate(in, outSamples), nil
	case BSpline5:
		return bspline5Interpolate(in, outSamples), nil
	case Lagrange4:
		return lagrange4Interpolate(in, outSamples), nil
	case Lagrange6:
		return lagrange6Interpolate(in, outSamples), nil
	case Watte:
		return watteInterpolate(in, outSamples), nil
	case Parabolic2x:
		return parabolic2xInterpolate(in, outSamples), nil
	case Osculating4:
		return osculating4Interpolate(in, outSamples), nil
	case Osculating6:
		return osculating6Interpolate(in, outSamples), nil
	case Hermite4:
		return hermite4Interpolate(in, outSamples), nil
	case Hermite6_3:
		return hermite6_3Interpolate(in, outSamples), nil
	case Hermite6_5:
		return hermite6_5Interpolate(in, outSamples), nil
	case CubicSpline:
		return applyCubicSpline(in, outSamples), nil
	case MonotonicCubic:
		return applyMonotonicCubic(in, outSamples), nil
	case Lanczos2:
		return lanczos2Interpolate(in, outSamples), nil
	case Lanczos3:
		return lanczos3Interpolate(in, outSamples), nil
	case Bezier:
		return bezierInterpolate(in, outSamples), nil
	case Akima:
		return applyAkimaSpline(in, outSamples), nil
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

// applyCubicSpline applies natural cubic spline interpolation
func applyCubicSpline(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}
	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	// Create x values for input points
	x := make([]float64, len(in))
	for i := range x {
		x[i] = float64(i)
	}

	// Compute spline coefficients
	a, b, c, d := cubicSplineCoefficients(x, in)

	out := make([]float64, outSamples)
	var ratio float64
	if outSamples > 1 {
		ratio = float64(len(in)-1) / float64(outSamples-1)
	} else {
		ratio = 0
	}

	for i := range out {
		pos := float64(i) * ratio
		j := int(pos)
		if j >= len(in)-1 {
			j = len(in) - 2
		}
		if j < 0 {
			j = 0
		}

		dx := pos - float64(j)
		dx2 := dx * dx
		dx3 := dx2 * dx

		out[i] = a[j] + b[j]*dx + c[j]*dx2 + d[j]*dx3
	}

	return out
}

// applyMonotonicCubic applies Fritsch-Carlson monotonic cubic interpolation
func applyMonotonicCubic(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}
	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	// Create x values for input points
	x := make([]float64, len(in))
	for i := range x {
		x[i] = float64(i)
	}

	// Compute monotonic slopes
	m := monotonicCubicSlopes(x, in)

	out := make([]float64, outSamples)
	var ratio float64
	if outSamples > 1 {
		ratio = float64(len(in)-1) / float64(outSamples-1)
	} else {
		ratio = 0
	}

	for i := range out {
		pos := float64(i) * ratio
		j := int(pos)
		if j >= len(in)-1 {
			j = len(in) - 2
		}
		if j < 0 {
			j = 0
		}

		h := x[j+1] - x[j]
		t := (pos - x[j]) / h
		t2 := t * t
		t3 := t2 * t

		// Hermite basis functions
		h00 := 2*t3 - 3*t2 + 1
		h10 := t3 - 2*t2 + t
		h01 := -2*t3 + 3*t2
		h11 := t3 - t2

		out[i] = h00*in[j] + h10*h*m[j] + h01*in[j+1] + h11*h*m[j+1]
	}

	return out
}

// applyAkimaSpline applies Akima spline interpolation
func applyAkimaSpline(in []float64, outSamples int) []float64 {
	if len(in) == 0 {
		return []float64{}
	}
	if len(in) == 1 {
		out := make([]float64, outSamples)
		for i := range out {
			out[i] = in[0]
		}
		return out
	}

	// Create x values for input points
	x := make([]float64, len(in))
	for i := range x {
		x[i] = float64(i)
	}

	// Compute Akima slopes
	m := akimaSlopes(x, in)

	out := make([]float64, outSamples)
	var ratio float64
	if outSamples > 1 {
		ratio = float64(len(in)-1) / float64(outSamples-1)
	} else {
		ratio = 0
	}

	for i := range out {
		pos := float64(i) * ratio
		j := int(pos)
		if j >= len(in)-1 {
			j = len(in) - 2
		}
		if j < 0 {
			j = 0
		}

		h := x[j+1] - x[j]
		t := (pos - x[j]) / h
		t2 := t * t
		t3 := t2 * t

		// Hermite basis functions
		h00 := 2*t3 - 3*t2 + 1
		h10 := t3 - 2*t2 + t
		h01 := -2*t3 + 3*t2
		h11 := t3 - t2

		out[i] = h00*in[j] + h10*h*m[j] + h01*in[j+1] + h11*h*m[j+1]
	}

	return out
}

// InterpolateInt performs interpolation on integer input data and returns integer output
// This function minimizes conversions by converting to float64 only once at the start
// and back to int only once at the end (with rounding)
func InterpolateInt(in []int, outSamples int, interpolatorType InterpolatorType) (out []int, err error) {
	if len(in) == 0 {
		return []int{}, nil
	}

	// Convert []int to []float64 once
	inFloat := make([]float64, len(in))
	for i, v := range in {
		inFloat[i] = float64(v)
	}

	// Perform interpolation using the existing float64 function
	outFloat, err := Interpolate(inFloat, outSamples, interpolatorType)
	if err != nil {
		return nil, err
	}

	// Convert []float64 back to []int once with rounding
	out = make([]int, len(outFloat))
	for i, v := range outFloat {
		// Round to nearest integer
		if v >= 0 {
			out[i] = int(v + 0.5)
		} else {
			out[i] = int(v - 0.5)
		}
	}

	return out, nil
}
