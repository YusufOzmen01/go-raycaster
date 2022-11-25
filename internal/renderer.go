package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Renderer interface {
	Render(world WorldMap, player Player)
}

type renderer struct {
	StartX int
	StartY int
	Width  int
	Height int
}

func NewRenderer(startX, startY, width, height int) Renderer {
	return &renderer{
		StartX: startX,
		StartY: startY,
		Width:  width,
		Height: height,
	}
}

func (r *renderer) Render(world WorldMap, player Player) {
	rays := player.GetRays()

	if len(rays) > 0 {
		for i, ray := range rays {
			if i+1 != len(rays) && i > 0 {
				if rays[i-1].Side != ray.Side && rays[i+1].Side != ray.Side {
					ray.Side = rays[i-1].Side
				}
			}

			playerX, playerY := player.Position()

			diffX, diffY := playerX-ray.X, playerY-ray.Y
			dist := math.Sqrt(diffX*diffX+diffY*diffY) * math.Cos(ray.Angle)

			columnWidth := float64(r.Width / len(rays))
			columnHeight := 50 / dist * float64(r.Height)

			columnX := float64(r.StartX) + float64(i+1)*columnWidth
			columnY := float64(r.StartY+r.Height)/2 - columnHeight/2

			r, g, b := 252, 143, 0

			switch ray.Side {
			case SideUp, SideDown:
				r, g, b = 161, 91, 0
			}

			rl.DrawRectangle(int32(columnX), int32(columnY), int32(columnWidth), int32(columnHeight), rl.NewColor(uint8(r), uint8(g), uint8(b), 255))
		}
	}
}
