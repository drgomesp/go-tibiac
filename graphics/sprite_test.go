// +build test

package graphics

import (
	"testing"

	"image/color"
	"unsafe"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewSprite(t *testing.T) {
	Convey("Given a new sprite", t, func() {
		sprite := NewSprite()

		Convey("Then the sprite should have a size equivalent to its width multiplied by its height", func() {
			So(sprite.Pixels, ShouldHaveLength, Width*Height)

			Convey("And the sprite byte size should be the amount of pixels multiplied by the pixel byte size", func() {
				So(sprite.Capacity, ShouldEqual, Width*Height*unsafe.Sizeof(color.RGBA{}))
			})
		})
	})
}
