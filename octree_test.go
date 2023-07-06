package octree

import "testing"

func TestNewOctree(t *testing.T) {
	tree := NewOctree(&BoundingBox{
		position: NewPosition(0, 0, 0),
		lengthX:  100,
		widthZ:   200,
		heightY:  300,
	}, 100, 4)

	// 随机插入点
	for _,:=range 
	tree.Root().AddEntity(&Entity{
		Key:      "1",
		position: NewPosition(1, 2, 3),
	})
}
