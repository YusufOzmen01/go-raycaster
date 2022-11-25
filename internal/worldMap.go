package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	EmptyWall = 0
	FullWall  = 1

	SideNone  = 0
	SideUp    = 1
	SideDown  = 2
	SideLeft  = 3
	SideRight = 4
)

type WorldMap interface {
	RenderMap()
	GetWall(x, y float64) int
	GetCellSize() int
	CheckWallCollision(x, y, w, h float64) (bool, int, int)
}

type worldMap struct {
	Map      [][]int
	CellSize int
}

func NewWorldMap(cellSize int) WorldMap {
	return &worldMap{
		Map: [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 0, 0, 0, 1, 0, 1},
			{1, 0, 1, 0, 0, 1, 0, 1},
			{1, 0, 1, 0, 0, 0, 0, 1},
			{1, 1, 1, 0, 0, 1, 0, 1},
			{1, 0, 0, 0, 0, 1, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
		},
		CellSize: cellSize,
	}
}

func (wm *worldMap) RenderMap() {
	rl.DrawRectangle(0, 0, int32(len(wm.Map)*wm.CellSize), int32(len(wm.Map[0])*wm.CellSize), rl.DarkGray)

	for y := 0; y < len(wm.Map); y++ {
		for x := 0; x < len(wm.Map[0]); x++ {
			if wm.Map[y][x] == FullWall {
				rl.DrawRectangle(int32(x*wm.CellSize+1), int32(y*wm.CellSize+1), int32(wm.CellSize-1), int32(wm.CellSize-1), rl.White)
			}
		}
	}
}

func (wm *worldMap) GetWall(x, y float64) int {
	return wm.Map[int(y)/wm.CellSize][int(x)/wm.CellSize]
}

func (wm *worldMap) GetCellSize() int {
	return wm.CellSize
}

func (wm *worldMap) CheckWallCollision(x, y, w, h float64) (bool, int, int) {
	if x < 0 || y < 0 || int(y/float64(wm.CellSize)) > len(wm.Map)-1 || int(x/float64(wm.CellSize)) > len(wm.Map[0])-1 {
		return false, 0, EmptyWall
	}

	if wm.GetWall(x, y) == EmptyWall {
		return false, 0, EmptyWall
	}

	box1 := rl.Rectangle{
		X:      float32(x),
		Y:      float32(y),
		Width:  float32(w),
		Height: float32(h),
	}

	box2 := rl.Rectangle{
		X:      float32(int(x) - int(x)%wm.CellSize),
		Y:      float32(int(y) - int(y)%wm.CellSize),
		Width:  float32(wm.CellSize),
		Height: float32(wm.CellSize),
	}

	collision := rl.CheckCollisionRecs(box1, box2)
	side := SideNone

	cellCenterX, cellCenterY := float64(box2.X+box2.Width/2), float64(box2.Y+box2.Height/2)
	xDistance, yDistance := math.Abs(x-cellCenterX), math.Abs(y-cellCenterY)

	if xDistance > yDistance {
		if cellCenterX > x {
			side = SideRight
		} else {
			side = SideLeft
		}
	} else if yDistance > xDistance {
		if cellCenterY > y {
			side = SideUp
		} else {
			side = SideDown
		}
	}

	return collision, side, wm.GetWall(x, y)
}
