package octree

import (
	"fmt"
)

type Vector3d struct {
	Position *Position
	LengthX  int // 长 X
	WidthZ   int // 宽 Z
	HeightY  int // 高 Y
}

func (b *Vector3d) String() string {
	// TODO 未想好打印啥
	return fmt.Sprintf("")
}

// intersectWithPoint 判断某个点是否在盒子里
func (b *Vector3d) intersectWithPoint(position *Position) bool {
	// 实体如果是一个点
	if b.Position.X <= position.X && position.X < b.Position.X+b.LengthX &&
		b.Position.Y <= position.Y && position.Y < b.Position.Y+b.HeightY &&
		b.Position.Z <= position.Z && position.Z < b.Position.Z+b.WidthZ {
		return true
	}
	return false
}

func (b *Vector3d) intersectWithBox(box *Vector3d) bool {
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

func (b *Vector3d) cuboid(color string) (result []interface{}) {
	min := []interface{}{b.Position.X, b.Position.Y, b.Position.Z}
	max := []interface{}{b.Position.X + b.LengthX, b.Position.Y + b.HeightY, b.Position.Z + b.WidthZ}

	result = append(result,
		// Front face
		[]interface{}{min[0], min[1], max[2]},
		[]interface{}{max[0], min[1], max[2]},
		[]interface{}{max[0], max[1], max[2]},
		[]interface{}{min[0], max[1], max[2]},

		[]interface{}{min[0], min[1], min[2]},
		[]interface{}{max[0], min[1], min[2]},
		[]interface{}{max[0], max[1], min[2]},
		[]interface{}{min[0], max[1], min[2]},

		[]interface{}{min[0], max[1], min[2]},
		[]interface{}{max[0], max[1], min[2]},
		[]interface{}{max[0], max[1], max[2]},
		[]interface{}{min[0], max[1], max[2]},

		[]interface{}{min[0], min[1], min[2]},
		[]interface{}{max[0], min[1], min[2]},
		[]interface{}{max[0], min[1], max[2]},
		[]interface{}{min[0], min[1], max[2]},

		[]interface{}{max[0], min[1], min[2]},
		[]interface{}{max[0], max[1], min[2]},
		[]interface{}{max[0], max[1], max[2]},
		[]interface{}{max[0], min[1], max[2]},

		[]interface{}{min[0], min[1], min[2]},
		[]interface{}{min[0], max[1], min[2]},
		[]interface{}{min[0], max[1], max[2]},
		[]interface{}{min[0], min[1], max[2]},
	)
	return result
	//list []opts.Chart3DData
	//list = append(list,
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, max.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, min.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{max.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//	opts.Chart3DData{Value: []interface{}{min.X, min.Y, max.Z}, ItemStyle: &opts.ItemStyle{Color: color}},
	//)
	return
}
