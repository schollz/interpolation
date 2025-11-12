# interpolation

[![Release](https://img.shields.io/github/v/release/schollz/interpolation.svg)](https://github.com/schollz/interpolation/releases/latest)
[![CI](https://github.com/schollz/interpolation/actions/workflows/CI.yml/badge.svg)](https://github.com/schollz/interpolation/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/schollz/interpolation/branch/main/graph/badge.svg)](https://codecov.io/gh/schollz/interpolation)
[![Go Reference](https://pkg.go.dev/badge/github.com/schollz/interpolation.svg)](https://pkg.go.dev/github.com/schollz/interpolation)


A Go package with an array of polynomial interpolators for resampling audio.

<img width="1584" height="813" alt="newplot(1)" src="https://github.com/user-attachments/assets/da03a1dc-c8fe-48f5-a1bb-16c60d8f96d9" />

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
BenchmarkResampleWAVFile/Linear-32                  1552            835038 ns/op
BenchmarkResampleWAVFile/DropSample-32              1537            849516 ns/op
BenchmarkResampleWAVFile/BSpline3-32                 685           1671143 ns/op
BenchmarkResampleWAVFile/BSpline5-32                 462           2531651 ns/op
BenchmarkResampleWAVFile/Lagrange4-32                679           1730536 ns/op
BenchmarkResampleWAVFile/Lagrange6-32                444           2796961 ns/op
BenchmarkResampleWAVFile/Watte-32                    896           1488571 ns/op
BenchmarkResampleWAVFile/Parabolic2x-32              784           1468642 ns/op
BenchmarkResampleWAVFile/Osculating4-32              646           1990904 ns/op
BenchmarkResampleWAVFile/Osculating6-32              427           2525010 ns/op
BenchmarkResampleWAVFile/Hermite4-32                 570           2215825 ns/op
BenchmarkResampleWAVFile/Hermite6_3-32               411           2687895 ns/op
BenchmarkResampleWAVFile/Hermite6_5-32               403           3145303 ns/op
BenchmarkResampleWAVFile/Lanczos2-32                 213           5560200 ns/op
BenchmarkResampleWAVFile/Lanczos3-32                 147           8437112 ns/op
BenchmarkResampleWAVFile/Bezier-32                   556           2163774 ns/op
```

## License

MIT
