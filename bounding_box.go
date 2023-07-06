package octree

import (
	"fmt"
)

type BoundingBox struct {
	position *Position
	lengthX  int // 长 X
	widthZ   int // 宽 Z
	heightY  int // 高 Y
}

func (b *BoundingBox) String() string {
	// TODO 未想好打印啥
	return fmt.Sprintf("")
}

// AllPoints 获取实体上连线的所有点
func (e *BoundingBox) AllPoints() []*Position {
	//   Z  底部的点位置
	//   2 ^------------. 3
	//     |            |
	//     |            |
	//     |            |
	//     '            '
	//     |            |
	//  0  '----------->' 1    X
	//

	//    Z  底部的点位置
	//   6 ^------------. 7
	//     |            |
	//     |            |
	//     |            |
	//     '            '
	//     |            |
	//  4  '------------> 5  X
	//
	return []*Position{
		// 下面四个点
		&Position{X: e.position.X, Y: e.position.Y, Z: e.position.Z},
		&Position{X: e.position.X + e.lengthX, Y: e.position.Y, Z: e.position.Z},
		&Position{X: e.position.X, Y: e.position.Y, Z: e.position.Z + e.widthZ},
		&Position{X: e.position.X + e.lengthX, Y: e.position.Y, Z: e.position.Z + e.widthZ},

		// 上面四个点
		&Position{X: e.position.X, Y: e.position.Y + e.heightY, Z: e.position.Z},
		&Position{X: e.position.X + e.lengthX, Y: e.position.Y + e.heightY, Z: e.position.Z},
		&Position{X: e.position.X, Y: e.position.Y + e.heightY, Z: e.position.Z + e.widthZ},
		&Position{X: e.position.X + e.lengthX, Y: e.position.Y + e.heightY, Z: e.position.Z + e.widthZ},
	}
}

// intersectWithPoint 判断某个点是否在盒子里
func (b *BoundingBox) intersectWithPoint(position *Position) bool {
	// 实体如果是一个点
	if b.position.X <= position.X && position.X < b.position.X+b.lengthX &&
		b.position.Y <= position.Y && position.Y < b.position.Y+b.heightY &&
		b.position.Z <= position.Z && position.Z < b.position.Z+b.widthZ {
		return true
	}
	return false
}

func (b *BoundingBox) intersectWithBox(box *BoundingBox) bool {
	aMinX := b.position.X
	aMaxX := b.position.X + b.lengthX
	aMinY := b.position.Y
	aMaxY := b.position.Y + b.heightY
	aMinZ := b.position.Z
	aMaxZ := b.position.Z + b.widthZ

	bMinX := box.position.X
	bMaxX := box.position.X + box.lengthX
	bMinY := box.position.Y
	bMaxY := box.position.Y + box.heightY
	bMinZ := box.position.Z
	bMaxZ := box.position.Z + box.widthZ
	// 如果两个长方体在某个坐标轴上的投影没有重叠，那么这两个长方体就不可能相交
	if aMinX > bMaxX || aMaxX < bMinX {
		return false
	}
	if aMinY > bMaxY || aMaxY < bMinY {
		return false
	}
	if aMinZ > bMaxZ || aMaxZ < bMinZ {
		return false
	}

	return true
}
