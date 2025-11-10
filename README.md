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
BenchmarkInterpolators/DropSample-32         	    1630	    727904 ns/op
BenchmarkInterpolators/Linear-32             	    1574	    731301 ns/op
BenchmarkInterpolators/BSpline3-32           	    1430	    830065 ns/op
BenchmarkInterpolators/BSpline5-32           	    1489	    800694 ns/op
BenchmarkInterpolators/Lagrange4-32          	    1626	    731049 ns/op
BenchmarkInterpolators/Lagrange6-32          	    1422	    776290 ns/op
BenchmarkInterpolators/Watte-32              	    1641	    729529 ns/op
BenchmarkInterpolators/Parabolic2x-32        	    1482	    812000 ns/op
BenchmarkInterpolators/Osculating4-32        	    1616	    740163 ns/op
BenchmarkInterpolators/Osculating6-32        	    1532	    782763 ns/op
BenchmarkInterpolators/Hermite4-32           	    1638	    733020 ns/op
BenchmarkInterpolators/Hermite6_3-32         	    1544	    770836 ns/op
BenchmarkInterpolators/Hermite6_5-32         	    1543	    770243 ns/op
BenchmarkInterpolators/CubicSpline-32        	   60253	     19571 ns/op
BenchmarkInterpolators/MonotonicCubic-32     	  141849	      9216 ns/op
BenchmarkInterpolators/Lanczos2-32           	    1585	    757603 ns/op
BenchmarkInterpolators/Lanczos3-32           	    1473	    765207 ns/op
BenchmarkInterpolators/Bezier-32             	    1560	    758721 ns/op
BenchmarkInterpolators/Akima-32              	  155307	      7629 ns/op
```
