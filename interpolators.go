package interpolators

// InterpolatorType defines the type of interpolation to use
type InterpolatorType int

const (
	// None returns the input data as-is without any interpolation
	None InterpolatorType = iota
)

// Interpolate performs interpolation on the input data based on the specified type
func Interpolate(in []float64, interpolatorType InterpolatorType) (out []float64, err error) {
	switch interpolatorType {
	case None:
		// None type returns input exactly as it was
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	default:
		out = make([]float64, len(in))
		copy(out, in)
		return out, nil
	}
}
