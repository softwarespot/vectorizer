package vectorizer

import (
	"testing"
)

func Test_New(t *testing.T) {
	vz1 := New[string](128)

	v1 := NewVector[string]()
	vz1.Add(v1, "Event 1", 1)
	vz1.Add(v1, "Event 2", 2)
	vz1.Add(v1, "Event 3", 3)

	v2 := NewVector[string]()
	vz1.Add(v2, "Event 2", 2)
	vz1.Add(v2, "Event 3", 3)

	t.Log(v1.CosineSimilarity(v2))

	t.Log("before normalize:", v1.data)
	v1.Normalize()
	t.Log("after normalize:", v1.data)

	type CustomKey struct {
		ID   int
		Name string
	}
	ck1 := CustomKey{ID: 1, Name: "A"}
	ck2 := CustomKey{ID: 1, Name: "A"}
	ck3 := CustomKey{ID: 2, Name: "B"}
	ck4 := CustomKey{ID: 3, Name: "C"}

	vz2 := New[CustomKey](128)

	cv1 := NewVector[CustomKey]()
	vz2.Add(cv1, ck1, 1)
	vz2.Add(cv1, ck3, 2)
	vz2.Add(cv1, ck4, 3)

	cv2 := NewVector[CustomKey]()
	vz2.Add(cv2, ck2, 1)
	vz2.Add(cv2, ck4, 5)

	t.Log(cv1.CosineSimilarity(cv2))
	t.Log(cv1.ToDense())
	t.Log(cv1.String())

	t.Log(cv2.ToDense())
	t.Log(cv2.String())
}
