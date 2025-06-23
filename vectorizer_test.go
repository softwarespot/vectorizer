package vectorizer

import (
	"testing"
)

func Test_New(t *testing.T) {
	vz1 := New[string](128)

	v1 := NewVector[string]()
	vz1.ApplyTo(v1, "Event 1", 1)
	vz1.ApplyTo(v1, "Event 2", 2)
	vz1.ApplyTo(v1, "Event 3", 3)

	v2 := NewVector[string]()
	vz1.ApplyTo(v2, "Event 2", 2)
	vz1.ApplyTo(v2, "Event 3", 3)

	t.Log(v1.CosineSimilarity(v2))

	t.Log("before normalize:", v1.data)
	v1.Normalize()
	t.Log("after normalize:", v1.data)

	type CustomKey struct {
		ID   int
		Name string
	}
	ck1 := CustomKey{
		ID:   1,
		Name: "A",
	}
	ck2 := CustomKey{
		ID:   1,
		Name: "A",
	}

	vz2 := New[CustomKey](128)

	cv1 := NewVector[CustomKey]()
	vz2.ApplyTo(cv1, ck1, 1)

	cv2 := NewVector[CustomKey]()
	vz2.ApplyTo(cv2, ck2, 1)

	t.Log(cv1.CosineSimilarity(cv2))
	t.Log(cv1.String())
}
