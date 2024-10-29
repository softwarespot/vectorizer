# Vectorizer

[![Go Reference](https://pkg.go.dev/badge/github.com/softwarespot/vectorizer.svg)](https://pkg.go.dev/github.com/softwarespot/vectorizer) ![Go Tests](https://github.com/softwarespot/vectorizer/actions/workflows/go.yml/badge.svg)

**Vectorizer** is a generic compatible module, which provides an efficient way to manage and manipulate sparse vectors. It supports operations such as applying values to vectors, calculating magnitudes, and computing cosine similarity.

## Features

-   **Sparse Vector Representation**: Efficiently store vectors with non-zero values only.
-   **Vector Operations**: Includes methods for scaling, normalizing, and calculating cosine similarity.

## Prerequisites

-   Go 1.23.0 or above

## Installation

```bash
go get -u github.com/softwarespot/vectorizer
```

## Usage

A basic example of using **Vectorizer**.

```Go
package main

import (
	"fmt"

	"github.com/softwarespot/vectorizer"
)

func main() {
	// Create a vectorizer with the type "string"
   	vz := vectorizer.New[string](128)

	vec1 := vectorizer.NewVector[string]()
	vz.ApplyTo(vec1, "Event 1", 1)
	vz.ApplyTo(vec1, "Event 2", 2)
	vz.ApplyTo(vec1, "Event 3", 3)

	vec2 := vectorizer.NewVector[string]()
	vz.ApplyTo(vec2, "Event 2", 2)
	vz.ApplyTo(vec2, "Event 3", 3)

	fmt.Println(vec1.CosineSimilarity(vec2)) // outputs: 0.9636241116594316
}
```

## License

The code has been licensed under the [MIT](https://opensource.org/license/mit) license.
