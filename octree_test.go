package octree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewOctree(t *testing.T) {
	lenghtX := 100
	widthZ := 200
	heightY := 300
	tree := NewOctree(&BoundingBox{
		Position: NewPosition(0, 0, 0),
		LengthX:  lenghtX,
		WidthZ:   widthZ,
		HeightY:  heightY,
	}, 100, 4)

	// 随机插入点
	for i := 0; i < 100; i++ {
		x := rand.Intn(lenghtX)
		z := rand.Intn(widthZ)
		y := rand.Intn(heightY)
		tree.Root().AddEntity(&Entity{
			Key:      fmt.Sprintf("%v", i+1),
			position: NewPosition(x, y, z),
		})
	}
	fmt.Println(tree)
}
