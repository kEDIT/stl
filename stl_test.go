package stl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacetOperations(t *testing.T) {
	v := newFacet(0.0, -1.0, 0.0)
	v.insertVertices([]float32{0.0, 0.0, 0.0},
		[]float32{1.0, 0.0, 0.0},
		[]float32{0.0, 0.0, 1.0})

	g, err := v.getVertex(2)
	if err != nil {
		t.Errorf("error retrieving vertex")
	}
	assert.Equal(t, []float32{0.0, 0.0, 1.0}, g)

	_, err = v.getVertex(4)
	if err == nil {
		t.Errorf("index should be out of range")
	}
}
