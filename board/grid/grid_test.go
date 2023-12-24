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
				Coordinates{
					X: 0,
					Y: 0,
				}},
			C2: Cell{
				Coordinates{
					X: 3,
					Y: 0,
				}},
			Expected: []Cell{
				{
					Coordinates{
						X: 1,
						Y: 0,
					},
				},
				{
					Coordinates{
						X: 2,
						Y: 0,
					},
				},
			},
		},
		{
			Name: "Column",
			C1: Cell{
				Coordinates{
					X: 0,
					Y: 0,
				}},
			C2: Cell{
				Coordinates{
					X: 0,
					Y: 3,
				}},
			Expected: []Cell{
				{
					Coordinates{
						X: 0,
						Y: 1,
					},
				},
				{
					Coordinates{
						X: 0,
						Y: 2,
					},
				},
			},
		},
		{
			Name: "Ascending diagonal",
			C1: Cell{
				Coordinates{
					X: 0,
					Y: 0,
				}},
			C2: Cell{
				Coordinates{
					X: 3,
					Y: 3,
				}},
			Expected: []Cell{
				{
					Coordinates{
						X: 1,
						Y: 1,
					},
				},
				{
					Coordinates{
						X: 2,
						Y: 2,
					},
				},
			},
		},
		{
			Name: "Descending diagonal",
			C1: Cell{
				Coordinates{
					X: 0,
					Y: 7,
				}},
			C2: Cell{
				Coordinates{
					X: 3,
					Y: 4,
				}},
			Expected: []Cell{
				{
					Coordinates{
						X: 1,
						Y: 6,
					},
				},
				{
					Coordinates{
						X: 2,
						Y: 5,
					},
				},
			},
		},
		{
			Name: "Directionally non-connected",
			C1: Cell{
				Coordinates{
					X: 0,
					Y: 7,
				}},
			C2: Cell{
				Coordinates{
					X: 1,
					Y: 5,
				}},
			Expected: []Cell{},
		},
		{
			Name: "Identical cells",
			C1: Cell{
				Coordinates{
					X: 0,
					Y: 0,
				}},
			C2: Cell{
				Coordinates{
					X: 0,
					Y: 0,
				}},
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
