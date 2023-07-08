package octree

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math"
	"math/rand"
	"time"
)

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
	leaf     bool // 是否为叶子节点
	tree     *Octree
	Deep     int           // 深度
	Location int           // 父节点中的位置
	parent   *Node         `json:"-"` // 父节点
	Children map[int]*Node // 子节点
	Bounds   *Vector3d     // 矩形的范围

	Entities  []*Entity   // 实体
	Buildings []*Building // 建筑
}

// NewNode 新建一个节点
func (n *Node) NewNode(bounds *Vector3d, location int) *Node {

	return &Node{
		leaf: true,

		Deep:     n.Deep + 1,
		tree:     n.tree,
		parent:   n,
		Children: make(map[int]*Node),
		Bounds:   bounds,
		Location: location,
		Entities: []*Entity{},
	}
}

// AddEntity 添加实体
func (n *Node) AddEntity(entity *Entity) {
	// 判断是否需要分割
	if n.leaf && n.needCut() {
		n.split()
	}
	//fmt.Println("n.Bounds:", n.Bounds)
	if n.leaf {
		// 直接添加一个实体
		n.Entities = append(n.Entities, entity)
		return
	}

	// 非叶子节点往下递归
	// 找到对应的区域进行添加
	for _, children := range n.Children {
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
		n.Buildings = append(n.Buildings, building)
		return
	}
	// 非叶子节点往下递归
	// 找到对应的区域进行添加
	for _, children := range n.Children {
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
	return len(n.Entities)+1 > n.tree.MaxCap && n.Deep+1 <= n.tree.MaxDeep && n.canCut()
}

// canCut 检查节点是否可以分割
func (n *Node) canCut() bool {
	if n.Bounds.LengthX >= 2 && n.Bounds.WidthZ >= 2 && n.Bounds.HeightY >= 2 {
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
	halfLengthX := int(math.Floor(float64(n.Bounds.LengthX) / 2))
	halfWidthZ := int(math.Floor(float64(n.Bounds.WidthZ) / 2))
	halfHeightY := int(math.Floor(float64(n.Bounds.HeightY) / 2))
	//下面的格子
	n.Children[bottomOne] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X + halfLengthX, Y: n.Bounds.Position.Y, Z: n.Bounds.Position.Z + halfWidthZ},
		LengthX:  n.Bounds.LengthX - halfLengthX,
		WidthZ:   n.Bounds.WidthZ - halfWidthZ,
		HeightY:  halfHeightY,
	}, bottomOne)
	n.Children[bottomTwo] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X, Y: n.Bounds.Position.Y, Z: n.Bounds.Position.Z + halfWidthZ},
		LengthX:  halfLengthX,
		WidthZ:   n.Bounds.WidthZ - halfWidthZ,
		HeightY:  halfHeightY,
	}, bottomTwo)
	n.Children[bottomThree] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X, Y: n.Bounds.Position.Y, Z: n.Bounds.Position.Z},
		LengthX:  halfLengthX,
		WidthZ:   halfWidthZ,
		HeightY:  halfHeightY,
	}, bottomThree)

	n.Children[bottomFour] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X + halfLengthX, Y: n.Bounds.Position.Y, Z: n.Bounds.Position.Z},
		LengthX:  n.Bounds.LengthX - halfLengthX,
		WidthZ:   halfWidthZ,
		HeightY:  halfHeightY,
	}, bottomFour)

	// 上面的格子
	n.Children[topOne] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X + halfLengthX, Y: n.Bounds.Position.Y + halfHeightY, Z: n.Bounds.Position.Z + halfWidthZ},
		LengthX:  n.Bounds.LengthX - halfLengthX,
		WidthZ:   n.Bounds.WidthZ - halfWidthZ,
		HeightY:  n.Bounds.HeightY - halfHeightY,
	}, topOne)
	n.Children[topTwo] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X, Y: n.Bounds.Position.Y + halfHeightY, Z: n.Bounds.Position.Z + halfWidthZ},
		LengthX:  halfLengthX,
		WidthZ:   n.Bounds.WidthZ - halfWidthZ,
		HeightY:  n.Bounds.HeightY - halfHeightY,
	}, topTwo)
	n.Children[topThree] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X, Y: n.Bounds.Position.Y + halfHeightY, Z: n.Bounds.Position.Z},
		LengthX:  halfLengthX,
		WidthZ:   halfWidthZ,
		HeightY:  n.Bounds.HeightY - halfHeightY,
	}, topThree)

	n.Children[topFour] = n.NewNode(&Vector3d{
		Position: &Position{X: n.Bounds.Position.X + halfLengthX, Y: n.Bounds.Position.Y + halfHeightY, Z: n.Bounds.Position.Z},
		LengthX:  n.Bounds.LengthX - halfLengthX,
		WidthZ:   halfWidthZ,
		HeightY:  n.Bounds.HeightY - halfHeightY,
	}, topFour)

	// 将当前节点上的建筑转移到子节点上
	for _, building := range n.Buildings {
		for _, node := range n.Children {
			if node.Bounds.intersectWithBox(building.Bounds) {
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
		for _, b := range n.Buildings {
			if b.Bounds.intersectWithBox(building.Bounds) {
				return true
			}
		}
		return false
	}

	for _, node := range n.Children {
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
	return n.Bounds.intersectWithPoint(entity.Position)
}

// intersectWithBuilding 判断建筑是否在节点范围内
func (n *Node) intersectWithBuilding(building *Building) bool {
	return n.Bounds.intersectWithBox(building.Bounds)
}

func (n *Node) ExportEntity(surface3D *charts.Surface3D) {
	var list []opts.Chart3DData
	if n.leaf {
		for _, v := range n.Entities {
			list = append(list, opts.Chart3DData{Value: []interface{}{
				v.Position.X, v.Position.Y, v.Position.Z},
				ItemStyle: &opts.ItemStyle{
					// 实体的点为黑色
					Color: "#000",
				},
			})

		}

		if len(list) == 0 {
			return
		}
		fmt.Println("entity:", list)
		surface3D.AddSeries("entity", list)
		return
	}
	for _, children := range n.Children {
		children.ExportEntity(surface3D)
	}
	return
}

func (n *Node) Export(surface3D *charts.Surface3D, key string) {
	color := "#53fc19" // 节点的点为绿色
	if n.parent == nil {
		//surface3D.AddSeries(key, n.Bounds.cuboid(color))
		surface3D.MultiSeries = append(surface3D.MultiSeries, charts.SingleSeries{
			Name: key,
			Type: surface3D.Type(),
			Data: n.Bounds.cuboid(color),
			//CoordSystem: types.ChartCartesian3D,
		})
	}
	for i := 0; i < 8; i++ {
		children, ok := n.Children[i]
		if !ok {
			continue
		}
		//surface3D.AddSeries(fmt.Sprintf("%v-%v", key, children.Location), children.Bounds.cuboid(color))
		surface3D.MultiSeries = append(surface3D.MultiSeries, charts.SingleSeries{
			Name: fmt.Sprintf("%v-%v", key, children.Location),
			Type: surface3D.Type(),
			Data: children.Bounds.cuboid(color),
			//CoordSystem: types.ChartCartesian3D,
		})
	}
	for _, children := range n.Children {
		children.Export(surface3D, fmt.Sprintf("%v-%v", key, children.Location))
	}
	return
}

func (n *Node) ExportBuilding(surface3D *charts.Surface3D) {
	fmt.Println(n.leaf)
	fmt.Println(n.Children)
	fmt.Println(n.Buildings)
	if n.leaf {
		for _, building := range n.Buildings {
			fmt.Println(building)
			time.Sleep(time.Millisecond * 10)
			list := building.Bounds.cuboid("#53fc19")
			//list := building.Bounds.cuboid(RandColor())
			if len(list) == 0 {
				continue
			}
			//surface3D.AddSeries(
			//	fmt.Sprintf("building:%v", building.Key),
			//	list,
			//)
			surface3D.MultiSeries = append(surface3D.MultiSeries, charts.SingleSeries{
				Name: fmt.Sprintf("building:%v", building.Key),
				Type: surface3D.Type(),
				Data: list,
				//CoordSystem: types.ChartCartesian3D,
			})
		}

		return
	}
	for _, children := range n.Children {
		children.ExportBuilding(surface3D)
	}
	return
}

func RandColor() string {
	rand.Seed(time.Now().UnixNano())

	red := rand.Intn(256)
	green := rand.Intn(256)
	blue := rand.Intn(256)

	color := fmt.Sprintf("#%02X%02X%02X", red, green, blue)
	return color
}
