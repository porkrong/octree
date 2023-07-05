package octree

import "fmt"

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

// ALlPoints 获取实体上连线的所有点
func (e *BoundingBox) ALlPoints() []*Position {
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
