package stl

import (
	"math"
	"testing"
)

func TestFacetOperations(t *testing.T) {
	v := newFacet(0.0, -1.0, 0.0)
	v.insertVertex([]float32{0.0, 0.0, 0.0})
	v.insertVertex([]float32{1.0, 0.0, 0.0})

	g, err := v.getVertex(2)
	if err != nil {
		t.Errorf("error retrieving vertex")
	}

	f32 := float32(7.)/3 - float32(4.)/3 - float32(1.)
	eps := math.Abs(f32)
	// assert.Equal(t, []float32{1.0, 0.0, 0.0}, g)

	_, err = v.getVertex(4)
	if err == nil {
		t.Errorf("index should be out of range")
	}
}
