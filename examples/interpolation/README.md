# Interpolation Example

This example demonstrates all the interpolation methods available in the `github.com/schollz/interpolation` package and visualizes them using Plotly.

## What it does

The example:
1. Creates a sample input curve (sine wave with 8 data points)
2. Interpolates the curve using each interpolation method to 100 samples:
   - **DropSample**: 0th-order B-spline (nearest neighbor)
   - **Linear**: 1st-order B-spline (linear interpolation)
   - **BSpline3**: 3rd-order B-spline (cubic, 4-point)
   - **BSpline5**: 5th-order B-spline (5th degree, 6-point)
3. Exports the data to a JSON file
4. Creates a Python script to visualize the results
5. Generates an interactive HTML plot showing all curves overlaid

## Prerequisites

- Go 1.24.9 or later
- Python 3
- Plotly: `pip install plotly`

## How to run

```bash
cd examples/interpolation
go run main.go
```

This will:
- Generate `interpolation_data.json` with the interpolated data
- Create `plot_interpolation.py` with the plotting script
- Generate `interpolation_plot.html` with the interactive visualization

## Viewing the results

Open `interpolation_plot.html` in your web browser to see an interactive plot comparing all interpolation methods. The original data points are shown as black dots, and each interpolation method is displayed as a colored line.

## Understanding the results

- **Original Data Points** (black dots): The input data being interpolated
- **DropSample** (step-like): Holds each value until the next sample
- **Linear**: Creates straight lines between points
- **BSpline3**: Smooth cubic curve that approximates the data
- **BSpline5**: Even smoother 5th-degree curve with higher quality approximation

The B-spline methods produce smoother curves but may not pass exactly through all original points, as they balance smoothness with fidelity to the data.
