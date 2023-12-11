package grid

type Coordinates struct {
	X int
	Y int
}

type Cell struct {
	Coordinates
}

func (c Cell) Up() Cell {
	c.X++
	return c
}

func (c Cell) Down() Cell {
	c.X--
	return c
}

func (c Cell) Right() Cell {
	c.Y++
	return c
}

func (c Cell) Left() Cell {
	c.Y--
	return c
}
