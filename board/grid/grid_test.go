package grid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrid_GetMedialCells(t *testing.T) {
	type TestCase struct {
		Name     string
		C1       Cell
		C2       Cell
		Expected []Cell
	}

	testCases := []TestCase{
		{
			Name: "Row",
			C1: Cell{
				X: 0,
				Y: 0,
			},
			C2: Cell{
				X: 3,
				Y: 0,
			},
			Expected: []Cell{
				{
					X: 1,
					Y: 0,
				},

				{
					X: 2,
					Y: 0,
				},
			},
		},
		{
			Name: "Column",
			C1: Cell{
				X: 0,
				Y: 0,
			},
			C2: Cell{
				X: 0,
				Y: 3,
			},
			Expected: []Cell{
				{
					X: 0,
					Y: 1,
				},
				{
					X: 0,
					Y: 2,
				},
			},
		},
		{
			Name: "Ascending diagonal",
			C1: Cell{
				X: 0,
				Y: 0,
			},
			C2: Cell{
				X: 3,
				Y: 3,
			},
			Expected: []Cell{
				{
					X: 1,
					Y: 1,
				},
				{
					X: 2,
					Y: 2,
				},
			},
		},
		{
			Name: "Descending diagonal",
			C1: Cell{
				X: 0,
				Y: 7,
			},
			C2: Cell{

				X: 3,
				Y: 4,
			},
			Expected: []Cell{
				{
					X: 1,
					Y: 6,
				},
				{
					X: 2,
					Y: 5,
				},
			},
		},
		{
			Name: "Directionally non-connected",
			C1: Cell{
				X: 0,
				Y: 7,
			},
			C2: Cell{
				X: 1,
				Y: 5,
			},
			Expected: []Cell{},
		},
		{
			Name: "Identical cells",
			C1: Cell{
				X: 0,
				Y: 0,
			},
			C2: Cell{
				X: 0,
				Y: 0,
			},
			Expected: []Cell{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := GetMedialCells(tc.C1, tc.C2)
			assert.ElementsMatch(t, tc.Expected, actual)
		})
	}
}

func TestGrid_IsValidCell(t *testing.T) {
	g := NewGrid(3)

	assert.True(t, g.IsValidCell(Cell{X: 2, Y: 2}))

	assert.False(t, g.IsValidCell(Cell{X: 3, Y: 2}))

	assert.False(t, g.IsValidCell(Cell{X: 2, Y: 3}))

}
