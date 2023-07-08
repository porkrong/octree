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
	for i := 0; i < 10; i++ {
		x := rand.Intn(lengthX)
		z := rand.Intn(widthZ)
		y := rand.Intn(heightY)
		tree.Root.AddEntity(&Entity{
			Key:      fmt.Sprintf("%v", i+1),
			Position: NewPosition(x, y, z),
		})
	}

	// 随机插入一个建筑
	for i := 0; i < 10; i++ {
		xMin := rand.Intn(lengthX)
		zMin := rand.Intn(widthZ)
		yMin := rand.Intn(heightY)
		width := rand.Intn(50) + 1

		tree.Root.AddBuilding(&Building{
			Key: fmt.Sprintf("%v", i+1),
			Bounds: &Vector3d{
				Position: NewPosition(xMin, yMin, zMin),
				LengthX:  width,
				HeightY:  width,
				WidthZ:   width,
			},
		})
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
