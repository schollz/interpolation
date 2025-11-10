package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"

	"github.com/schollz/interpolation"
)

// PlotData represents data for a single curve
type PlotData struct {
	Name string    `json:"name"`
	X    []float64 `json:"x"`
	Y    []float64 `json:"y"`
}

func main() {
	// Create sample input data - a simple sine-like curve with a few points
	inputSamples := 8
	input := make([]float64, inputSamples)
	for i := 0; i < inputSamples; i++ {
		x := float64(i) / float64(inputSamples-1) * 2 * math.Pi
		input[i] = math.Sin(x)
	}

	// Number of output samples for interpolation (upsampling)
	outputSamples := 100

	// Store all plot data
	var allPlots []PlotData

	// Add original data points
	originalX := make([]float64, len(input))
	for i := range input {
		originalX[i] = float64(i) / float64(len(input)-1) * 2 * math.Pi
	}
	allPlots = append(allPlots, PlotData{
		Name: "Original Data Points",
		X:    originalX,
		Y:    input,
	})

	// Test each interpolation type
	interpolationTypes := []struct {
		name string
		typ  interpolators.InterpolatorType
	}{
		{"None", interpolators.None},
		{"DropSample", interpolators.DropSample},
		{"Linear", interpolators.Linear},
		{"BSpline3", interpolators.BSpline3},
		{"BSpline5", interpolators.BSpline5},
		{"Lagrange4", interpolators.Lagrange4},
		{"Lagrange6", interpolators.Lagrange6},
	}

	for _, interp := range interpolationTypes {
		// Skip "None" for upsampling since it just returns the input
		if interp.typ == interpolators.None {
			continue
		}

		output, err := interpolators.Interpolate(input, outputSamples, interp.typ)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error interpolating with %s: %v\n", interp.name, err)
			continue
		}

		// Create x values for the interpolated points
		x := make([]float64, len(output))
		for i := range output {
			x[i] = float64(i) / float64(len(output)-1) * 2 * math.Pi
		}

		allPlots = append(allPlots, PlotData{
			Name: interp.name,
			X:    x,
			Y:    output,
		})

		fmt.Printf("Generated %d samples using %s interpolation\n", len(output), interp.name)
	}

	// Write data to JSON file
	jsonData, err := json.MarshalIndent(allPlots, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("interpolation_data.json", jsonData, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing JSON file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nData written to interpolation_data.json")
	fmt.Println("Running Python Plotly visualization...")

	// Create and run Python script
	pythonScript := `import json
import plotly.graph_objects as go

# Read the data
with open('interpolation_data.json', 'r') as f:
    data = json.load(f)

# Create figure
fig = go.Figure()

# Add each curve
for i, curve in enumerate(data):
    if curve['name'] == 'Original Data Points':
        # Plot original points as markers only
        fig.add_trace(go.Scatter(
            x=curve['x'],
            y=curve['y'],
            mode='markers',
            name=curve['name'],
            marker=dict(size=10, color='black', symbol='circle')
        ))
    else:
        # Plot interpolated curves as lines
        fig.add_trace(go.Scatter(
            x=curve['x'],
            y=curve['y'],
            mode='lines',
            name=curve['name'],
            line=dict(width=2)
        ))

# Update layout
fig.update_layout(
    title='Interpolation Methods Comparison',
    xaxis_title='X (radians)',
    yaxis_title='Y',
    hovermode='closest',
    legend=dict(
        yanchor="top",
        y=0.99,
        xanchor="left",
        x=0.01
    )
)

# Save to HTML file
fig.write_html('interpolation_plot.html')
print('\nPlot saved to interpolation_plot.html')
print('Open this file in a web browser to view the interactive plot.')
`

	err = os.WriteFile("plot_interpolation.py", []byte(pythonScript), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing Python script: %v\n", err)
		os.Exit(1)
	}

	// Run the Python script
	cmd := exec.Command("python3", "plot_interpolation.py")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError running Python script: %v\n", err)
		fmt.Println("Make sure plotly is installed: pip install plotly")
		os.Exit(1)
	}
}
