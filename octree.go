package octree

// Octree 八叉树
type Octree struct {
	maxCap, maxDeep int
	root            *Node
}

// NewOctree 构建一颗八叉树
func NewOctree(box *BoundingBox, maxCap, maxDeep int) *Octree {
	tree := &Octree{
		maxCap:  maxCap,
		maxDeep: maxDeep,
	}
	root := &Node{
		position: box.position,
		parent:   nil,
		children: [8]*Node{},
		bounds:   box,
		location: -1,
		tree:     tree,
	}
	tree.root = root
	return tree
}

// Retrieve 传入一个盒子，检测当前是否有实体跟该盒子产生碰撞
func (t *Octree) Retrieve(entity *Entity) {
	for _, node := range t.root.children {

	}
}
