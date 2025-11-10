# interpolation

[![CI](https://github.com/schollz/interpolation/actions/workflows/CI.yml/badge.svg)](https://github.com/schollz/interpolation/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/schollz/interpolation/branch/main/graph/badge.svg)](https://codecov.io/gh/schollz/interpolation)
[![Go Reference](https://pkg.go.dev/badge/github.com/schollz/interpolation.svg)](https://pkg.go.dev/github.com/schollz/interpolation)
[![Release](https://img.shields.io/github/v/release/schollz/interpolation.svg)](https://github.com/schollz/interpolation/releases/latest)


A Go package with an array of polynomial interpolators for resampling audio.

<img width="1584" height="813" alt="newplot(1)" src="https://github.com/user-attachments/assets/64e2a7ae-3606-4456-9acf-337fc91e72ea" />



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
