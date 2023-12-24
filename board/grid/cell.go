package grid

type Coordinates struct {
	X int
	Y int
}

type Cell struct {
	Coordinates
}

type Direction string

const (
	Up    Direction = "Up"
	Down  Direction = "Down"
	Right Direction = "Right"
	Left  Direction = "Left"
)

var Directions = []Direction{Up, Down, Right, Left}

type DiagonalDirection string

const (
	RightAscending  DiagonalDirection = "RightAscending"
	RightDescending DiagonalDirection = "RightDescending"
	LeftAscending   DiagonalDirection = "LeftAscending"
	LeftDescending  DiagonalDirection = "LeftDescending"
)

var DiagonalDirections = []DiagonalDirection{RightAscending, RightDescending, LeftAscending, LeftDescending}

func (c Cell) Up() Cell {
	c.Y++
	return c
}

func (c Cell) Down() Cell {
	c.Y--
	return c
}

func (c Cell) Right() Cell {
	c.X++
	return c
}

func (c Cell) Left() Cell {
	c.X--
	return c
}

var Movements = map[Direction]func(c Cell) Cell{
	Up:    func(c Cell) Cell { return c.Up() },
	Down:  func(c Cell) Cell { return c.Down() },
	Right: func(c Cell) Cell { return c.Right() },
	Left:  func(c Cell) Cell { return c.Left() },
}

var DiagonalMovements = map[DiagonalDirection]func(r Cell) Cell{
	RightAscending:  func(c Cell) Cell { return c.Right().Up() },
	RightDescending: func(c Cell) Cell { return c.Right().Down() },
	LeftAscending:   func(c Cell) Cell { return c.Left().Up() },
	LeftDescending:  func(c Cell) Cell { return c.Left().Down() },
}
