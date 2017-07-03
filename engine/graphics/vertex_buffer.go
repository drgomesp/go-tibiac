package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VertexBuffer struct {
	ID uint32
}

func NewVertexBuffer(data []float32) *VertexBuffer {
	buffer := &VertexBuffer{}

	gl.GenBuffers(1, &buffer.ID)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.ID)

	return buffer
}
