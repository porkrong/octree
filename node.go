package octree

import "math"

const (
	bottomOne = 0  + iota
	bottomTwo = 0  + iota
	bottomThree = 0  + iota
	bottomFour = 0  + iota
	topOne = 0  + iota
	topTwo = 0  + iota
	topThree = 0  + iota
	topFour = 0  + iota
)

type Node struct {
	leaf     bool // 是否为叶子节点
	deep     int  // 深度
	tree     *Octree
	position *Position
	parent   *Node        // 父节点
	children [8]*Node     // 子节点
	bounds   *BoundingBox // 矩形的范围
	location int          // 父节点中的位置
	entities []*Entity    // 实体
}

// NewNode 新建一个节点
func (n *Node) NewNode(bounds *BoundingBox, location int) *Node {

	return &Node{
		leaf:     true,
		deep:     n.deep + 1,
		tree:     n.tree,
		position: bounds.position,
		parent:   n,
		children: [8]*Node{},
		bounds:   bounds,
		location: location,
		entities: []*Entity{},
	}
}

func (n *Node) Add(entity *Entity) {
	// 判断是否需要分割
	if n.leaf && n.needCut() {
		n.subdivide()
	}

	// 非叶子节点往下递归
	if !n.leaf {

		n.children[n.findChildrenIndex(x, y)].Add(x, y, name)
		return
	}
}

// findChildrenIndex 根据实体寻找子节点的方位列表
func (n *Node) findChildrenIndex(entity *Entity)  {
	// 获取该实体所有的连接的点
	points:=entity.ALlPoints()
	for _,point:=range points{
		// 获取点所在格子的位置
	}
	return rightDown
}

func (n *Node)find

// 检测节点是否需要进行拆分
func (n *Node) needCut() bool {
	return len(n.entities)+1 > n.tree.maxCap && n.deep+1 <= n.tree.maxDeep && n.canCut()
}

// canCut 检查节点是否可以分割
func (n *Node) canCut() bool {
	if n.bounds.lengthX >= 2 && n.bounds.widthZ>=2 && n.bounds.heightY>=2 {
		return true
	}
	return false
}

// subdivide 对节点进行拆分
func (n *Node) subdivide() {
	if !n.leaf {
		return
	}
	halfLengthX:=int(math.Floor(float64(n.bounds.lengthX)/2))
	halfWidthZ:=int(math.Floor(float64(n.bounds.widthZ)/2))
	halfHeightY:=int(math.Floor(float64(n.bounds.heightY)/2))
	//下面的格子
	n.children[bottomOne] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X+halfLengthX, Y: n.bounds.position.Y, Z: n.bounds.position.Z+halfWidthZ},
		lengthX: n.bounds.lengthX-halfLengthX,
		widthZ: n.bounds.widthZ-halfWidthZ,
		heightY: halfHeightY,
	}, bottomOne)
	n.children[bottomTwo] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y, Z: n.bounds.position.Z+halfWidthZ},
		lengthX: halfLengthX,
		widthZ: n.bounds.widthZ-halfWidthZ,
		heightY: halfHeightY,
	}, bottomTwo)
	n.children[bottomThree] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y, Z: n.bounds.position.Z},
		lengthX: halfLengthX,
		widthZ: halfWidthZ,
		heightY: halfHeightY,
	}, bottomThree)

	n.children[bottomFour] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X+halfLengthX, Y: n.bounds.position.Y, Z: n.bounds.position.Z},
		lengthX: n.bounds.lengthX-halfLengthX,
		widthZ: halfWidthZ,
		heightY: halfHeightY,
	}, bottomFour)


	// 上面的格子
	n.children[topOne] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X+halfLengthX, Y: n.bounds.position.Y+halfHeightY, Z: n.bounds.position.Z+halfWidthZ},
		lengthX: n.bounds.lengthX-halfLengthX,
		widthZ: n.bounds.widthZ-halfWidthZ,
		heightY: n.bounds.heightY-halfHeightY,
	}, topOne)
	n.children[topTwo] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y+halfHeightY, Z: n.bounds.position.Z+halfWidthZ},
		lengthX: halfLengthX,
		widthZ: n.bounds.widthZ-halfWidthZ,
		heightY: n.bounds.heightY-halfHeightY,
	}, topTwo)
	n.children[topThree] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X, Y: n.bounds.position.Y+halfHeightY, Z: n.bounds.position.Z},
		lengthX: halfLengthX,
		widthZ: halfWidthZ,
		heightY: n.bounds.heightY-halfHeightY,
	}, topThree)

	n.children[topFour] = n.NewNode(&BoundingBox{
		position: &Position{X: n.bounds.position.X+halfLengthX, Y: n.bounds.position.Y+halfHeightY, Z: n.bounds.position.Z},
		lengthX: n.bounds.lengthX-halfLengthX,
		widthZ: halfWidthZ,
		heightY: n.bounds.heightY-halfHeightY,
	}, topFour)


	// 将当前节点上的实体转移到子节点上
	for _, entity := range n.entities {
		for _, node := range n.children {
			// 检测当前实体是否在节点中
			if node.intersects(entity) {
				node.Add(entity)
			}
		}
	}
	n.leaf = false
}

// intersects 检查坐标是否在节点范围内
func (n *Node) intersects(entity *Entity) bool {
	switch entity.EntityType() {
	case Point:
		// 实体如果是一个点
		if n.bounds.position.X <= entity.position.X && entity.position.X <= n.bounds.position.X+n.bounds.lengthX &&
			n.bounds.position.Y <= entity.position.Y && entity.position.Y <= n.bounds.position.Y+n.bounds.heightY &&
			n.bounds.position.Z <= entity.position.Z && entity.position.Z <= n.bounds.position.Z +n.bounds.widthZ{
			return true
		}
		return false
	default:
		// 实体是一个盒子
		// 需要检测8个点的位置是否存在一个在该节点中

		return false
	}

}


// GetIndexs 找到某个 box 所在区域列表
func (n *Node)GetIndexs(box *BoundingBox)(list []int){
	if n.leaf{
		return []int{}
	}


	points:=box.ALlPoints()
	var inBottomOne,inBottomTwo,inBottomThree,inBottomFour,inTopOne,inTopTwo,inTopThree,inTopFour bool
	for _,point:=range points{
		//该点在该节点内的位置 下层第几象限 上层第几象限
		if

	}

	if inBottomOne{
		// 在下层第一象限
		list = append(list,bottomOne)
	}

	if inBottomTwo{
		// 在下层第二象限
		list = append(list,bottomTwo)
	}
	if inBottomThree{
		// 在下层第三象限
		list = append(list,bottomThree)
	}
	if inBottomFour{
		// 在下层第四象限
		list = append(list,bottomFour)
	}


	if inTopOne{
		// 在上层第一象限
		list = append(list,topOne)
	}

	if inTopTwo{
		// 在上层第二象限
		list = append(list,topTwo)
	}
	if inTopThree{
		// 在上层第三象限
		list = append(list,topThree)
	}
	if inTopFour{
		// 在上层第四象限
		list = append(list,topFour)
	}
	return
}