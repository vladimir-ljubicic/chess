package grid

type Coordinates struct {
	X int
	Y int
}

type Cell struct {
	Coordinates
}

func (c Cell) GoUp() Cell {
	c.X++
	return c
}

func (c Cell) GoDown() Cell {
	c.X--
	return c
}

func (c Cell) GoRight() Cell {
	c.Y++
	return c
}

func (c Cell) GoLeft() Cell {
	c.Y--
	return c
}
