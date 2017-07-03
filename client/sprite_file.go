package client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image/color"
	"io"
	"os"
)

// SpriteFile contains the sprite file signature and sprite count
type SpriteFile struct {
	// Extended file contains a 32-bit sprite count. Non-extended contains a 16-bit sprite count
	Extended, Loaded bool
	Signature, Count uint32

	signatureByteSize uint8
	countByteSize     uint8
	descriptor        *os.File
}

// NewSpriteFileFromPath loads a new sprite file from a file path
func NewSpriteFileFromPath(path string) (*SpriteFile, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("could not open file %v", err)
	}

	spriteFile := &SpriteFile{
		descriptor:        file,
		signatureByteSize: 4,
		countByteSize:     4,

		Extended: true, // TODO(daniel): How can this be checked?
	}

	if err := binary.Read(file, binary.LittleEndian, &spriteFile.Signature); err != nil {
		return nil, fmt.Errorf("could not read sprite file's signature: %v", err)
	}

	if spriteFile.Extended {
		if err := binary.Read(file, binary.LittleEndian, &spriteFile.Count); err != nil {
			return nil, fmt.Errorf("could not read sprite file's count: %v", err)
		}
	} else {
		var count uint16

		if err := binary.Read(file, binary.LittleEndian, &count); err != nil {
			return nil, fmt.Errorf("could not read sprite file's count: %v", err)
		}

		spriteFile.Count = uint32(count)
		spriteFile.countByteSize = 2
	}

	spriteFile.Loaded = true

	return spriteFile, nil
}

// LoadSpriteByID loads a sprite by its ID
func (spriteFile *SpriteFile) LoadSpriteByID(spriteID int64) (*Sprite, error) {
	if spriteFile.Loaded {
		spriteFile.descriptor.Seek(int64(spriteFile.signatureByteSize+spriteFile.countByteSize), io.SeekStart) // Skip signature and sprite count
	}

	spriteFile.descriptor.Seek((spriteID+1)*4, io.SeekStart)

	var addr uint32

	if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &addr); err != nil {
		return nil, fmt.Errorf("could not read sprite address (%v)", err)
	}

	if addr == 0 {
		return nil, errors.New("no sprite found for zero address")
	}

	// Move reader to offset of the sprite data, skipping 3 bytes of transparent data
	spriteFile.descriptor.Seek(int64(addr)+3, io.SeekStart)

	sprite := NewSprite()

	if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &sprite.Size); err == io.EOF {
		return nil, fmt.Errorf("could not read sprite size %v", err)
	}

	var (
		channels   = 3 // Using only 3 channels due to no alpha channel data
		pixelIndex = 0
	)

	for read := 0; read < int(sprite.Size); {
		var transparentPixels uint16
		var coloredPixels uint16

		if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &transparentPixels); err != nil {
			return nil, fmt.Errorf("could not read transparent pixels (%v)", err)
		}

		if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &coloredPixels); err != nil {
			return nil, fmt.Errorf("could not read colored pixels (%v)", err)
		}

		for i := 0; i < int(transparentPixels); i++ {
			sprite.Pixels[pixelIndex] = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
			pixelIndex++
		}

		for i := 0; i < int(coloredPixels); i++ {
			rgb := new(color.RGBA)

			if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &rgb.R); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &rgb.G); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			if err := binary.Read(spriteFile.descriptor, binary.LittleEndian, &rgb.B); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			sprite.Pixels[pixelIndex] = color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 0xFF}
			pixelIndex++
		}

		read += 4 + (channels * int(coloredPixels))
	}

	return sprite, nil
}
