package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/drgomesp/go-tibiac/client"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	m := client.NewSpriteManager("data/tibia986/Tibia.spr")

	spr, err := m.LoadSpriteByID(2)

	if err != nil {
		logrus.Fatal(err)
	}

	SavePNG(2, 32, 32, spr.Pixels)

	// if err := glfw.Init(); err != nil {
	// 	logrus.Fatalf("could not initialize glfw: %v", err)
	// }

	// defer glfw.Terminate()

	// window, err := glfw.CreateWindow(windowWidth, windowHeight, "go-tibiac v0.1.0", nil, nil)

	// if err != nil {
	// 	logrus.Fatalf("could not create window: %v", err)
	// }

	// window.MakeContextCurrent()

	// // Initialize glow
	// if err := gl.Init(); err != nil {
	// 	logrus.Fatalf("could not initialize glow: %v", err)
	// }

	// logrus.Infof("OpenGL version: %v", gl.GoStr(gl.GetString(gl.VERSION)))

	// for !window.ShouldClose() {
	// 	window.SwapBuffers()
	// 	glfw.PollEvents()
	// }
}

// SavePNG saves a PNG image from the image data
func SavePNG(ID uint32, w int, h int, data []color.RGBA) {
	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	for i := 0; i < len(data); i++ {
		x := i % w
		y := (i - x) / w

		img.Set(x, y, data[y*w+x])
	}

	f, err := os.Create(fmt.Sprintf("%d.png", ID))

	if err != nil {
		logrus.Errorf("could not create sprite image: %v", err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		logrus.Errorf("could not encode sprite image: %v", err)
	}

	if err := f.Close(); err != nil {
		logrus.Fatal(err)
	}
}
