package vectorizer

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"
)

// Vectorizer is a structure that manages dimensions for vectors.
type Vectorizer[T comparable] struct {
	dim      int
	dimByKey map[T]int
}

// New creates a new Vectorizer with an initial size.
func New[T comparable](size int) *Vectorizer[T] {
	return &Vectorizer[T]{
		dim:      0,
		dimByKey: make(map[T]int, size),
	}
}

// Add applies a value to the vector at the specified key.
func (vz *Vectorizer[T]) Add(v *Vector[T], key T, value float64) {
	dim, ok := vz.dimByKey[key]
	if !ok {
		vz.dim += 1
		dim = vz.dim
		vz.dimByKey[key] = dim
	}
	v.Add(dim, value)
}

// Vector represents a sparse vector.
type Vector[T comparable] struct {
	data            map[int]float64
	cachedMagnitude float64

	// An internal flag to indicate if the magnitude needs to be recalculated,
	// due to a change in the vector's data.
	calculateMagnitude bool
}

// NewVector creates a new empty Vector.
func NewVector[T comparable]() *Vector[T] {
	return &Vector[T]{
		data:               map[int]float64{},
		cachedMagnitude:    0,
		calculateMagnitude: true,
	}
}

// Add adds a value to the vector at the specified dimension.
// NOTE: When the value is zero, it is not added to the vector.
func (v *Vector[T]) Add(dim int, value float64) {
	if value == 0 {
		return
	}

	v.data[dim] += value
	v.calculateMagnitude = true
}

// Delete deletes a dimension from the vector.
func (v *Vector[T]) Delete(dim int) {
	if _, exists := v.data[dim]; exists {
		delete(v.data, dim)
		v.calculateMagnitude = true
	}
}

// Magnitude calculates the magnitude of the vector.
// OPTIMIZATION: The method caches the magnitude, if the vector has not changed since the last calculation.
func (v *Vector[T]) Magnitude() float64 {
	if !v.calculateMagnitude {
		return v.cachedMagnitude
	}

	var magnitude float64
	for _, value := range v.data {
		magnitude += value * value
	}
	magnitude = math.Sqrt(magnitude)

	v.cachedMagnitude = magnitude
	v.calculateMagnitude = false

	return magnitude
}

// DotProduct calculates the dot product with another vector.
func (v *Vector[T]) DotProduct(v2 *Vector[T]) float64 {
	if len(v.data) > len(v2.data) {
		v, v2 = v2, v
	}

	var dot float64
	for dim, value := range v.data {
		if value2, ok := v2.data[dim]; ok {
			dot += value * value2
		}
	}
	return dot
}

// CosineSimilarity calculates the cosine similarity between two vectors.
// It returns an error if either vector has a zero magnitude.
func (v *Vector[T]) CosineSimilarity(v2 *Vector[T]) (float64, error) {
	magnitude1 := v.Magnitude()
	if magnitude1 == 0 {
		return 0, errors.New("Vector: first vector has zero magnitude")
	}

	magnitude2 := v2.Magnitude()
	if magnitude2 == 0 {
		return 0, errors.New("Vector: second vector has zero magnitude")
	}
	return cosineSimilarity(v, v2, magnitude1, magnitude2), nil
}

func cosineSimilarity[T comparable](v1, v2 *Vector[T], magnitude1, magnitude2 float64) float64 {
	return v1.DotProduct(v2) / (magnitude1 * magnitude2)
}

// Normalize scales the vector to have a magnitude of 1.
// NOTE: If the vector has zero magnitude, it remains unchanged.
func (v *Vector[T]) Normalize() {
	magnitude := v.Magnitude()
	if magnitude == 0 {
		return
	}
	v.Scale(1 / magnitude)
}

// Scale multiplies all elements of the vector by a scalar value.
func (v *Vector[T]) Scale(scalar float64) {
	for dim := range v.data {
		v.data[dim] *= scalar
	}
	v.calculateMagnitude = true
}

// String returns a string representation of the vector.
func (v *Vector[T]) String() string {
	var b strings.Builder
	b.WriteString("Vector{")

	dims := make([]int, 0, len(v.data))
	for dim := range v.data {
		dims = append(dims, dim)
	}
	sort.Ints(dims)

	for i, dim := range dims {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d:%.4f", dim, v.data[dim])
	}
	b.WriteString("}")
	return b.String()
}
