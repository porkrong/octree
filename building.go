package octree

import "fmt"

type Building struct {
	Bounds *Vector3d
	Key    string
}

func (b *Building) String() string {
	return fmt.Sprintf(
		`
{
	Id: %v 
	Bounds:{
		X: %v, Y: %v, Z: %v
		LengthX: %v,
		WidthZ: %v,
		HeightY: %v,
	}, 
}`, b.Key, b.Bounds.Position.X, b.Bounds.Position.Y, b.Bounds.Position.Z, b.Bounds.LengthX, b.Bounds.HeightY, b.Bounds.WidthZ)
}
