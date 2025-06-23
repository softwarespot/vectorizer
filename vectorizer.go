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

// ApplyTo applies a value to the vector at the specified key.
func (vz *Vectorizer[T]) ApplyTo(v *Vector[T], key T, value float64) {
	dim, ok := vz.dimByKey[key]
	if !ok {
		vz.dim += 1
		dim = vz.dim
		vz.dimByKey[key] = dim
	}
	v.apply(dim, value)
}

// Vector represents a sparse vector.
type Vector[T comparable] struct {
	data             map[int]float64
	cachedMagnitude  float64
	reCacheMagnitude bool
}

// NewVector creates a new empty Vector.
func NewVector[T comparable]() *Vector[T] {
	return &Vector[T]{
		data:             map[int]float64{},
		reCacheMagnitude: true,
	}
}

func (v *Vector[T]) apply(dim int, value float64) {
	if value == 0 {
		return
	}

	v.data[dim] += value
	v.reCacheMagnitude = true
}

func (v *Vector[T]) magnitude() float64 {
	if !v.reCacheMagnitude {
		return v.cachedMagnitude
	}

	var magnitude float64
	for _, value := range v.data {
		magnitude += value * value
	}
	magnitude = math.Sqrt(magnitude)

	v.cachedMagnitude = magnitude
	v.reCacheMagnitude = false

	return magnitude
}

func (v *Vector[T]) dotProduct(v2 *Vector[T]) float64 {
	// Iterate over the smaller vector
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
// It returns an error if either vector has zero magnitude.
func (v *Vector[T]) CosineSimilarity(v2 *Vector[T]) (float64, error) {
	magnitude1 := v.magnitude()
	if magnitude1 == 0 {
		return 0, errors.New("cosine similarity cannot be calculated for the first vector, as it's 0")
	}

	magnitude2 := v2.magnitude()
	if magnitude2 == 0 {
		return 0, errors.New("cosine similarity cannot be calculated for the second vector, as it's 0")
	}
	return cosineSimilarity(v, v2, magnitude1, magnitude2), nil
}

// Scale multiplies all elements of the vector by a scalar value.
func (v *Vector[T]) Scale(scalar float64) {
	for dim := range v.data {
		v.data[dim] *= scalar
	}
	v.reCacheMagnitude = true
}

// Normalize scales the vector to have a magnitude of 1.
// If the vector has zero magnitude, it remains unchanged.
func (v *Vector[T]) Normalize() {
	magnitude := v.magnitude()
	if magnitude == 0 {
		return
	}
	v.Scale(1 / magnitude)
}

func cosineSimilarity[T comparable](v1, v2 *Vector[T], magnitude1, magnitude2 float64) float64 {
	return v1.dotProduct(v2) / (magnitude1 * magnitude2)
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
