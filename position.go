package octree

import "fmt"

type Position struct {
	X int
	Y int
	Z int
}

func (p Position) String() string {
	return fmt.Sprintf("(%v,%v,%v)", p.X, p.Y, p.Z)
}
