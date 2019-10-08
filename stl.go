package stl

import "fmt"

const (
	nFacetDefault = 10
)

type solid struct {
	name   string
	facets []facet
}

type facet struct {
	n    int
	norm [3]float32
	data [][]float32
}

func newFacet(nx, ny, nz float32) facet {
	f := facet{
		norm: [3]float32{nx, ny, nz},
		data: make([][]float32, 3),
	}
	for i := range f.data {
		f.data[i] = make([]float32, nFacetDefault)
	}
	return f
}

func (f *facet) insertVertices(vs ...[]float32) error {
	if len(vs) != 3 {
		return fmt.Errorf("invalid number of positional arguments (vertices need three)")
	}
	for _, v := range vs {
		for i := 0; i < 3; i++ {
			f.data[i] = append(f.data[i], v[i])
		}
		f.n++
	}
	return nil
}

func (f *facet) getVertex(i int) ([]float32, error) {
	if i >= f.n {
		return []float32{}, fmt.Errorf("index is out of bounds. max: %d, input: %d", f.n, i)
	}
	return []float32{f.data[0][i], f.data[1][i], f.data[2][i]}, nil
}

func newSolid(name string) *solid {
	return &solid{
		name:   name,
		facets: []facet{},
	}
}
