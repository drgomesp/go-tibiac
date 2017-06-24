package graphics

import (
	"image/color"
	"unsafe"
)

const (
	// Width of the sprite
	Width = 32
	// Height of the sprite
	Height = 32
)

// Sprite defined by pixel data
type Sprite struct {
	Size     uint16
	Capacity uint16
	Pixels   []color.RGBA
}

// NewSprite returns a new sprite with proper pixel data capacity
func NewSprite() *Sprite {
	return &Sprite{
		Capacity: uint16((Width * Height) * unsafe.Sizeof(color.RGBA{})),
		Pixels:   make([]color.RGBA, Width*Height),
	}
}
