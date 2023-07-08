package octree

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"io"
	"math/rand"
	"os"
	"testing"
)

func TestNewOctree(t *testing.T) {
	lengthX := 100
	widthZ := 200
	heightY := 300
	tree := NewOctree(&Vector3d{
		Position: NewPosition(0, 0, 0),
		LengthX:  lengthX,
		WidthZ:   widthZ,
		HeightY:  heightY,
	}, 2, 4)

	// 随机插入点
	for i := 0; i < 100; i++ {
		x := rand.Intn(lengthX)
		z := rand.Intn(widthZ)
		y := 0
		tree.Root.AddEntity(&Entity{
			Key:      fmt.Sprintf("%v", i+1),
			Position: NewPosition(x, y, z),
		})
	}

	// 随机插入一个建筑

	for i := 0; i < 1; i++ {
		for j := 0; j < 10; j++ {
			xMin := rand.Intn(lengthX)
			zMin := rand.Intn(widthZ)
			yMin := (i - 0) * 10
			width := rand.Intn(50) + 1
			building := &Building{
				Key: fmt.Sprintf("%v", (i+1)*100000+j),
				Bounds: &Vector3d{
					Position: NewPosition(xMin, yMin, zMin),
					LengthX:  width,
					HeightY:  width,
					WidthZ:   width,
				},
			}
			//fmt.Println("当前建筑:", building)
			// 判断是否碰撞
			result, collisionBuilding := tree.Root.collision(building)
			if result {
				fmt.Println("产生碰撞:")

				fmt.Println("被碰撞的建筑:", collisionBuilding)
				continue
			}
			tree.Root.AddBuilding(building)
		}

	}
	tree.Export()
}

// 测试检测建筑碰撞
func TestBuildingCollision(t *testing.T) {
	lengthX := 100
	widthZ := 200
	heightY := 300
	page := components.NewPage()
	tree := NewOctree(&Vector3d{
		Position: NewPosition(0, 0, 0),
		LengthX:  lengthX,
		WidthZ:   widthZ,
		HeightY:  heightY,
	}, 2, 4)
	surface3D := charts.NewSurface3D()
	surface3D.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "",
		}),
	)
	tree.Root.ExportEntity(surface3D)
	tree.Root.ExportBuilding(surface3D)
	f, err := os.Create("3d_grid.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))

}
