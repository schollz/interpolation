package main

import (
	"fmt"

	interpolators "github.com/schollz/interpolation"
)

func main() {
	// Example 1: Simple upsampling with Linear interpolation
	fmt.Println("Example 1: Upsampling integer data with Linear interpolation")
	input1 := []int{0, 10, 20, 30}
	output1, err := interpolators.InterpolateInt(input1, 7, interpolators.Linear)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Input:  %v\n", input1)
	fmt.Printf("Output: %v\n\n", output1)

	// Example 2: Downsampling with Linear interpolation
	fmt.Println("Example 2: Downsampling integer data with Linear interpolation")
	input2 := []int{0, 10, 20, 30, 40, 50}
	output2, err := interpolators.InterpolateInt(input2, 3, interpolators.Linear)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Input:  %v\n", input2)
	fmt.Printf("Output: %v\n\n", output2)

	// Example 3: Using BSpline3 interpolation
	fmt.Println("Example 3: Upsampling with BSpline3 interpolation")
	input3 := []int{1, 5, 10, 15, 20}
	output3, err := interpolators.InterpolateInt(input3, 9, interpolators.BSpline3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Input:  %v\n", input3)
	fmt.Printf("Output: %v\n\n", output3)

	// Example 4: Negative values
	fmt.Println("Example 4: Interpolating negative integer values")
	input4 := []int{-10, -5, 0, 5, 10}
	output4, err := interpolators.InterpolateInt(input4, 9, interpolators.Linear)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Input:  %v\n", input4)
	fmt.Printf("Output: %v\n\n", output4)

	// Example 5: Using Hermite4 (Catmull-Rom) interpolation
	fmt.Println("Example 5: Upsampling with Hermite4 (Catmull-Rom) interpolation")
	input5 := []int{0, 5, 15, 10, 20}
	output5, err := interpolators.InterpolateInt(input5, 9, interpolators.Hermite4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Input:  %v\n", input5)
	fmt.Printf("Output: %v\n\n", output5)

	// Example 6: Comparing different interpolation methods
	fmt.Println("Example 6: Comparing interpolation methods")
	input6 := []int{0, 100, 200}
	outSamples := 5

	methods := []struct {
		name string
		typ  interpolators.InterpolatorType
	}{
		{"Linear", interpolators.Linear},
		{"BSpline3", interpolators.BSpline3},
		{"Hermite4", interpolators.Hermite4},
		{"CubicSpline", interpolators.CubicSpline},
	}

	fmt.Printf("Input: %v\n", input6)
	for _, method := range methods {
		output, err := interpolators.InterpolateInt(input6, outSamples, method.typ)
		if err != nil {
			fmt.Printf("Error with %s: %v\n", method.name, err)
			continue
		}
		fmt.Printf("%-12s: %v\n", method.name, output)
	}
}
