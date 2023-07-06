package octree

import (
	"fmt"
)

type BoundingBox struct {
	Position *Position
	LengthX  int // 长 X
	WidthZ   int // 宽 Z
	HeightY  int // 高 Y
}

func (b *BoundingBox) String() string {
	// TODO 未想好打印啥
	return fmt.Sprintf("")
}

// intersectWithPoint 判断某个点是否在盒子里
func (b *BoundingBox) intersectWithPoint(position *Position) bool {
	// 实体如果是一个点
	if b.Position.X <= position.X && position.X < b.Position.X+b.LengthX &&
		b.Position.Y <= position.Y && position.Y < b.Position.Y+b.HeightY &&
		b.Position.Z <= position.Z && position.Z < b.Position.Z+b.WidthZ {
		return true
	}
	return false
}

func (b *BoundingBox) intersectWithBox(box *BoundingBox) bool {
	aMinX := b.Position.X
	aMaxX := b.Position.X + b.LengthX
	aMinY := b.Position.Y
	aMaxY := b.Position.Y + b.HeightY
	aMinZ := b.Position.Z
	aMaxZ := b.Position.Z + b.WidthZ

	bMinX := box.Position.X
	bMaxX := box.Position.X + box.LengthX
	bMinY := box.Position.Y
	bMaxY := box.Position.Y + box.HeightY
	bMinZ := box.Position.Z
	bMaxZ := box.Position.Z + box.WidthZ
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
