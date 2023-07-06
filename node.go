package octree

import "math"

const (
	bottomOne = 0 + iota
	bottomTwo
	bottomThree
	bottomFour
	topOne
	topTwo
	topThree
	topFour
)

type Node struct {
	leaf bool // 是否为叶子节点
	deep int  // 深度
	tree *Octree

	parent    *Node        // 父节点
	children  [8]*Node     // 子节点
	bounds    *BoundingBox // 矩形的范围
	location  int          // 父节点中的位置
	entities  []*Entity    // 实体
	buildings []*Building  // 建筑
}

// NewNode 新建一个节点
func (n *Node) NewNode(bounds *BoundingBox, location int) *Node {

	return &Node{
		leaf:     true,
		deep:     n.deep + 1,
		tree:     n.tree,
		parent:   n,
		children: [8]*Node{},
		bounds:   bounds,
		location: location,
		entities: []*Entity{},
	}
}

// AddEntity 添加实体
func (n *Node) AddEntity(entity *Entity) {
	// 判断是否需要分割
	if n.leaf && n.needCut() {
		n.split()
	}

	if n.leaf {
		// 直接添加一个实体
		n.entities = append(n.entities, entity)
		return
	}

	// 非叶子节点往下递归
	// 找到对应的区域进行添加
	for _, children := range n.children {
		//检测是否在该节点范围内
		if !children.intersectWithEntity(entity) {
			continue
		}
		// 递归的往下层子节点添加实体
		children.AddEntity(entity)
	}
	return
}

// MoveEntity 移动实体
func (n *Node) MoveEntity(entity *Entity, position *Position) {

}

// DeleteEntity 删除实体
func (n *Node) DeleteEntity(entity *Entity) {

}

// AddBuilding 添加建筑
func (n *Node) AddBuilding(building *Building) {
	if n.leaf {
		// 直接添加一个建筑
		n.buildings = append(n.buildings, building)
		return
	}
	// 非叶子节点往下递归
	// 找到对应的区域进行添加
	for _, children := range n.children {
		//检测是否在该节点范围内
		if !children.intersectWithBuilding(building) {
			continue
		}
		// 递归的往下层子节点添加建筑
		children.AddBuilding(building)
	}
	return
}

// 检测节点是否需要进行拆分
func (n *Node) needCut() bool {
	return len(n.entities)+1 > n.tree.maxCap && n.deep+1 <= n.tree.maxDeep && n.canCut()
}

// canCut 检查节点是否可以分割
func (n *Node) canCut() bool {
	if n.bounds.lengthX >= 2 && n.bounds.widthZ >= 2 && n.bounds.heightY >= 2 {
		return true
	}
	return false
}

// Merge 合并节点 TODO 未完成
func (n *Node) Merge() {

}

// split 对节点进行拆分
func (n *Node) split() {
	if !n.leaf {
		return
	}
	halfLengthX := int(math.Floor(float64(n.bounds.lengthX) / 2))
	halfWidthZ := int(math.Floor(float64(n.bounds.widthZ) / 2))
	halfHeightY := int(math.Floor(float64(n.bounds.heightY) / 2))
	//下面的格子
	n.children[bottomOne] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X + halfLengthX, Y: n.bounds.position.Y, Z: n.bounds.position.Z + halfWidthZ},
		lengthX:  n.bounds.lengthX - halfLengthX,
		widthZ:   n.bounds.widthZ - halfWidthZ,
		heightY:  halfHeightY,
	}, bottomOne)
	n.children[bottomTwo] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y, Z: n.bounds.position.Z + halfWidthZ},
		lengthX:  halfLengthX,
		widthZ:   n.bounds.widthZ - halfWidthZ,
		heightY:  halfHeightY,
	}, bottomTwo)
	n.children[bottomThree] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y, Z: n.bounds.position.Z},
		lengthX:  halfLengthX,
		widthZ:   halfWidthZ,
		heightY:  halfHeightY,
	}, bottomThree)

	n.children[bottomFour] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X + halfLengthX, Y: n.bounds.position.Y, Z: n.bounds.position.Z},
		lengthX:  n.bounds.lengthX - halfLengthX,
		widthZ:   halfWidthZ,
		heightY:  halfHeightY,
	}, bottomFour)

	// 上面的格子
	n.children[topOne] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X + halfLengthX, Y: n.bounds.position.Y + halfHeightY, Z: n.bounds.position.Z + halfWidthZ},
		lengthX:  n.bounds.lengthX - halfLengthX,
		widthZ:   n.bounds.widthZ - halfWidthZ,
		heightY:  n.bounds.heightY - halfHeightY,
	}, topOne)
	n.children[topTwo] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y + halfHeightY, Z: n.bounds.position.Z + halfWidthZ},
		lengthX:  halfLengthX,
		widthZ:   n.bounds.widthZ - halfWidthZ,
		heightY:  n.bounds.heightY - halfHeightY,
	}, topTwo)
	n.children[topThree] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y + halfHeightY, Z: n.bounds.position.Z},
		lengthX:  halfLengthX,
		widthZ:   halfWidthZ,
		heightY:  n.bounds.heightY - halfHeightY,
	}, topThree)

	n.children[topFour] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X + halfLengthX, Y: n.bounds.position.Y + halfHeightY, Z: n.bounds.position.Z},
		lengthX:  n.bounds.lengthX - halfLengthX,
		widthZ:   halfWidthZ,
		heightY:  n.bounds.heightY - halfHeightY,
	}, topFour)

	// 将当前节点上的建筑转移到子节点上
	for _, building := range n.buildings {
		for _, node := range n.children {
			if node.bounds.intersectWithBox(building.bounds) {
				// 对当前实体在该子节点中
				node.AddBuilding(building)
				break
			}

		}
	}

	// 将节点上的实体转移到子节点上
	n.leaf = false
}

// collision 检查当前节点内的建筑是否与该建筑产生碰撞
func (n *Node) collision(building *Building) bool {
	// 判断盒子是否与节点相交
	if !n.intersectWithBuilding(building) {
		return false
	}
	if n.leaf {
		// 如果是叶子节点
		// 检测叶子节点内的建筑是否与该建筑产生碰撞
		for _, b := range n.buildings {
			if b.bounds.intersectWithBox(building.bounds) {
				return true
			}
		}
		return false
	}

	for _, node := range n.children {
		//接着递归检测碰撞
		if node.collision(building) {
			// 检测到碰撞的情况。直接返回
			return true
		}
	}
	return false
}

// 判断实体是否在节点范围内
func (n *Node) intersectWithEntity(entity *Entity) bool {
	return n.bounds.intersectWithPoint(entity.position)
}

// intersectWithBuilding 判断建筑是否在节点范围内
func (n *Node) intersectWithBuilding(building *Building) bool {
	return n.bounds.intersectWithBox(building.bounds)
}
