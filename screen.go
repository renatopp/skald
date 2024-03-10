package sk

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renatopp/skald/utils"
)

type Screen struct {
	Surface       *ebiten.Image
	PixelsPerUnit float64

	position Vec2 // In units
	size     Vec2 // In units
	zoom     float64
	rotation float64
	color    color.Color
	minZoom  float64
	maxZoom  float64
}

func NewScreen() *Screen {
	s := &Screen{
		Surface:       ebiten.NewImage(int(SCREEN_WIDTH*PIXELS_PER_UNIT), int(SCREEN_HEIGHT*PIXELS_PER_UNIT)),
		PixelsPerUnit: PIXELS_PER_UNIT,
		position:      Vec2{X: -SCREEN_WIDTH / 2, Y: -SCREEN_HEIGHT / 2},
		size:          Vec2{X: SCREEN_WIDTH, Y: SCREEN_HEIGHT},
		zoom:          1,
		rotation:      0,
		color:         color.White,
		minZoom:       0.1,
		maxZoom:       10,
	}

	s.SetColor(color.RGBA{110, 164, 191, 255})

	return s
}

func (s *Screen) GetPosition() Vec2 {
	return s.position
}

func (s *Screen) SetPosition(x, y float64) {
	s.position = Vec2{X: x, Y: y}
}

func (s *Screen) GetSize() Vec2 {
	return s.size
}

func (s *Screen) SetSize(w, h float64) {
	s.size = Vec2{X: w, Y: h}
	s.Surface = ebiten.NewImage(int(w*s.PixelsPerUnit), int(h*s.PixelsPerUnit))
}

func (s *Screen) GetZoom() float64 {
	return s.zoom
}

func (s *Screen) SetZoom(zoom float64) {
	s.zoom = zoom
}

func (s *Screen) GetZoomLimits() (float64, float64) {
	return s.minZoom, s.maxZoom
}

func (s *Screen) SetZoomLimits(min, max float64) {
	s.minZoom = min
	s.maxZoom = max
	s.zoom = utils.Clamp(s.zoom, min, max)
}

func (s *Screen) GetRotation() float64 {
	return s.rotation
}

func (s *Screen) SetRotation(rotation float64) {
	s.rotation = rotation
}

func (s *Screen) GetColor() color.Color {
	return s.color
}

func (s *Screen) SetColor(color color.Color) {
	s.color = color
}

func (s *Screen) Clear() {
	s.Surface.Fill(s.color)
}