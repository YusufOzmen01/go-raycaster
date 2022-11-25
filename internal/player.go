package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	OneDegreeInRadians = 0.0174533
)

type Player interface {
	RenderPlayer(m WorldMap)
	UpdatePlayer(m WorldMap)
	GetRays() []*Ray
	Position() (float64, float64)
}

type Ray struct {
	X     float64
	Y     float64
	Angle float64
	Side  int
}

type player struct {
	X           float64
	Y           float64
	Angle       float64
	Speed       float64
	Sensitivity float64
	Size        int
	RayCount    int
	Rays        []*Ray
}

func NewPlayer(x, y, size, rayCount int, speed, sensitivity float64) Player {
	return &player{
		X:           float64(x),
		Y:           float64(y),
		Angle:       0,
		Speed:       speed,
		Size:        size,
		Sensitivity: sensitivity,
		RayCount:    rayCount,
		Rays:        make([]*Ray, 0),
	}
}

func (p *player) RenderPlayer(m WorldMap) {
	rad := p.Angle * math.Pi / 180

	tmp := make([]*Ray, 0)

	for i := -OneDegreeInRadians * 30; i < OneDegreeInRadians*30; i += OneDegreeInRadians * 30 * 2 / float64(p.RayCount) {
		startX := float64(p.Size)/2 + p.X
		startY := float64(p.Size)/2 + p.Y

		for d := 0; d < 500; d++ {
			endX := startX + math.Cos(i+rad)*float64(d)
			endY := startY + math.Sin(i+rad)*float64(d)

			collision, side, _ := m.CheckWallCollision(endX, endY, 1, 1)
			if collision {
				endX = startX + math.Cos(i+rad)*float64(d-1)
				endY = startY + math.Sin(i+rad)*float64(d-1)

				tmp = append(tmp, &Ray{
					X:     endX,
					Y:     endY,
					Side:  side,
					Angle: i,
				})

				break
			}

			rl.DrawLine(int32(startX), int32(startY), int32(endX), int32(endY), rl.Blue)
		}
	}

	p.Rays = tmp

	rl.DrawRectangle(int32(p.X), int32(p.Y), int32(p.Size), int32(p.Size), rl.Red)
}

func (p *player) GetRays() []*Ray {
	return p.Rays
}

func (p *player) Position() (float64, float64) {
	return p.X, p.Y
}

func (p *player) UpdatePlayer(m WorldMap) {
	rad := p.Angle * math.Pi / 180

	if rl.IsKeyDown(rl.KeyUp) {
		newX := p.X + math.Cos(rad)*p.Speed*float64(rl.GetFrameTime())
		newY := p.Y + math.Sin(rad)*p.Speed*float64(rl.GetFrameTime())

		collision, _, _ := m.CheckWallCollision(newX, newY, float64(p.Size), float64(p.Size))
		if !collision {
			p.X = newX
			p.Y = newY
		}
	} else if rl.IsKeyDown(rl.KeyDown) {
		newX := p.X - math.Cos(rad)*p.Speed*float64(rl.GetFrameTime())
		newY := p.Y - math.Sin(rad)*p.Speed*float64(rl.GetFrameTime())

		collision, _, _ := m.CheckWallCollision(newX, newY, float64(p.Size), float64(p.Size))
		if !collision {
			p.X = newX
			p.Y = newY
		}
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		p.Angle -= p.Sensitivity * float64(rl.GetFrameTime())

		if p.Angle <= 0 {
			p.Angle = 360
		}
	} else if rl.IsKeyDown(rl.KeyRight) {
		p.Angle += p.Sensitivity * float64(rl.GetFrameTime())

		if p.Angle >= 360 {
			p.Angle = 0
		}
	}
}
