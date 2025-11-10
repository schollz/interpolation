# interpolation

[![CI](https://github.com/schollz/interpolation/actions/workflows/CI.yml/badge.svg)](https://github.com/schollz/interpolation/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/schollz/interpolation/branch/main/graph/badge.svg)](https://codecov.io/gh/schollz/interpolation)
[![Go Reference](https://pkg.go.dev/badge/github.com/schollz/interpolation.svg)](https://pkg.go.dev/github.com/schollz/interpolation)
[![Release](https://img.shields.io/github/v/release/schollz/interpolation.svg)](https://github.com/schollz/interpolation/releases/latest)


A Go package with an array of polynomial interpolators for resampling audio.

<img width="1584" height="813" alt="newplot(1)" src="https://github.com/user-attachments/assets/64e2a7ae-3606-4456-9acf-337fc91e72ea" />

## Available Interpolators

This package includes 20 different interpolation methods:

### Basic Interpolators
- **None** - Returns input data as-is
- **DropSample** - 0th-order B-spline (nearest neighbor/sample-and-hold)
- **Linear** - 1st-order B-spline (linear interpolation)

### B-Spline Interpolators
- **BSpline3** - 3rd-order B-spline (4-point)
- **BSpline5** - 5th-order B-spline (6-point)

### Lagrange Interpolators
- **Lagrange4** - 4-point, 3rd-order Lagrange interpolator
- **Lagrange6** - 6-point, 5th-order Lagrange interpolator

### Hermite Interpolators
- **Hermite4** - 4-point, 3rd-order Hermite (Catmull-Rom spline)
- **Hermite6_3** - 6-point, 3rd-order Hermite interpolator
- **Hermite6_5** - 6-point, 5th-order Hermite interpolator

### Osculating Interpolators
- **Osculating4** - 4-point, 5th-order 2nd-order-osculating interpolator
- **Osculating6** - 6-point, 5th-order 2nd-order-osculating interpolator

### Specialized Interpolators
- **Watte** - 4-point, 2nd-order Watte tri-linear interpolator
- **Parabolic2x** - 4-point, 2nd-order parabolic 2x interpolator

### Spline Interpolators
- **CubicSpline** - Natural cubic spline with C² continuity
- **MonotonicCubic** - Fritsch-Carlson monotonic cubic (preserves monotonicity)
- **Akima** - Akima spline (robust to outliers)

### Windowed Sinc Interpolators
- **Lanczos2** - Windowed sinc with a=2 (4-point, high quality)
- **Lanczos3** - Windowed sinc with a=3 (6-point, highest quality)

### Other
- **Bezier** - Cubic Bezier curve interpolation

## Benchmarks

```bash
goos: linux
goarch: amd64
pkg: github.com/schollz/interpolation
cpu: 13th Gen Intel(R) Core(TM) i9-13900K
BenchmarkInterpolators/DropSample-32         	    1513	    692931 ns/op
BenchmarkInterpolators/Linear-32             	    1660	    691202 ns/op
BenchmarkInterpolators/BSpline3-32           	    1502	    787052 ns/op
BenchmarkInterpolators/BSpline5-32           	    1592	    753884 ns/op
BenchmarkInterpolators/Lagrange4-32          	    1710	    695600 ns/op
BenchmarkInterpolators/Lagrange6-32          	    1594	    739541 ns/op
BenchmarkInterpolators/Watte-32              	    1714	    695756 ns/op
BenchmarkInterpolators/Parabolic2x-32        	    1519	    768830 ns/op
BenchmarkInterpolators/Osculating4-32        	    1681	    710967 ns/op
BenchmarkInterpolators/Osculating6-32        	    1622	    739333 ns/op
BenchmarkInterpolators/Hermite4-32           	    1690	    695956 ns/op
BenchmarkInterpolators/Hermite6_3-32         	    1561	    735592 ns/op
BenchmarkInterpolators/Hermite6_5-32         	    1633	    732925 ns/op
```
