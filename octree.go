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
		leaf:     true,
		parent:   nil,
		children: [8]*Node{},
		bounds:   box,
		location: -1,
		tree:     tree,
	}
	// 根节点需要自动分裂成八个节点
	root.split()
	tree.root = root
	return tree
}

// Collision 传入一个建筑，检测当前是否有建筑跟该建筑产生碰撞
func (t *Octree) Collision(building *Building) bool {
	return t.root.collision(building)
}
