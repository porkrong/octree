package octree

import (
	"encoding/json"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"os"
)

// Octree 八叉树
type Octree struct {
	MaxCap, MaxDeep int
	Length          int
	Width           int
	Height          int
	Root            *Node
}

// NewOctree 构建一颗八叉树
func NewOctree(box *Vector3d, maxCap, maxDeep int) *Octree {
	tree := &Octree{
		MaxCap:  maxCap,
		MaxDeep: maxDeep,
		Length:  box.LengthX,
		Width:   box.WidthZ,
		Height:  box.HeightY,
	}
	root := &Node{
		leaf:     true,
		parent:   nil,
		Children: make(map[int]*Node),
		Bounds:   box,
		Location: -1,
		tree:     tree,
	}
	// 根节点需要自动分裂成八个节点
	root.split()
	tree.Root = root
	return tree
}

// Collision 传入一个建筑，检测当前是否有建筑跟该建筑产生碰撞
func (t *Octree) Collision(building *Building) (result bool, collisionBuilding *Building) {
	return t.Root.collision(building)
}

func (t *Octree) String() string {
	b, _ := json.MarshalIndent(t, "", "   ")
	return string(b)
}

// Export 导出整棵树
func (t *Octree) Export() {
	file, err := os.OpenFile("tree.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(t.String())
	if err != nil {
		panic(err)
	}
	page := components.NewPage()

	surface3D := charts.NewSurface3D()
	surface3D.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "",
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Show:      false,
			Dimension: "2",
			Min:       -1,
			Max:       1,
			InRange: &opts.VisualMapInRange{
				Color: []string{
					//"#313695",
					//"#4575b4",
				},
			},
		}),
		charts.WithXAxis3DOpts(opts.XAxis3D{
			Type: "value",
		}),
		charts.WithYAxis3DOpts(opts.YAxis3D{
			Name: "Z",
			Type: "value",
		}),

		charts.WithZAxis3DOpts(opts.ZAxis3D{
			Name: "Y",
			Type: "value",
		}),
		charts.WithGrid3DOpts(opts.Grid3D{
			ViewControl: &opts.ViewControl{
				AutoRotate: true,
			},
		}),
		//charts.WithToolboxOpts(opts.Toolbox{
		//	Show: true,
		//}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:      true,
			Enterable: true,
			TriggerOn: "click",
			Formatter: `{a}`,
		}),
	)

	t.Root.ExportBuilding(surface3D)
	t.Root.ExportEntity(surface3D)
	page.AddCharts(
		//surface3D1,
		surface3D,
	)
	f, err := os.Create("3d_grid.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
