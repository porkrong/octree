package octree

type Entity struct {
	position *Position
	bounds   *BoundingBox
	Key      string
}

const (
	Point = 0
	Box   = 1
)

// EntityType 获取到实体类型
func (e *Entity) EntityType() int {
	if e.bounds == nil {
		return Point
	}
	return Box
}
