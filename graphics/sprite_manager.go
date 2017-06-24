package graphics

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
)

// SpriteFile contains the sprite file signature and sprite count
type SpriteFile struct {
	Loaded    bool
	Signature uint32
	Count     uint32
}

// SpriteManager handles the sprite file loading as well as keeping reference to each invididual sprite
type SpriteManager struct {
	SpriteFile *SpriteFile
	Sprites    map[uint32]*Sprite

	filePath string
}

// NewSpriteManager returns a new sprite manager for a given sprite file path
func NewSpriteManager(filePath string) *SpriteManager {
	return &SpriteManager{
		SpriteFile: &SpriteFile{},
		filePath:   filePath,
	}
}

// LoadSpriteByID loads a sprite by its ID
func (spriteManager *SpriteManager) LoadSpriteByID(spriteID int) (*Sprite, error) {
	f, err := os.Open(spriteManager.filePath)

	if err != nil {
		return nil, fmt.Errorf("could not open file %v", err)
	}

	defer f.Close()

	if !spriteManager.SpriteFile.Loaded {
		binary.Read(f, binary.LittleEndian, &spriteManager.SpriteFile.Signature)
		binary.Read(f, binary.LittleEndian, &spriteManager.SpriteFile.Count)
	}

	// Move reader to offset (spriteID - 1 because the sprite ID begins at 2)
	f.Seek((int64(spriteID)+1)*4, io.SeekStart)

	var addr uint32

	if err := binary.Read(f, binary.LittleEndian, &addr); err != nil {
		return nil, fmt.Errorf("could not read sprite address (%v)", err)
	}

	if addr == 0 {
		return nil, errors.New("no sprite found for zero address")
	}

	// Move reader to offset of the sprite data, skipping 3 bytes of transparent data
	f.Seek(int64(addr)+3, io.SeekStart)

	if err != nil {
		return nil, fmt.Errorf("could not seek to sprite add %v", addr)
	}

	sprite := NewSprite()

	if err := binary.Read(f, binary.LittleEndian, &sprite.Size); err == io.EOF {
		return nil, fmt.Errorf("could not read sprite size %v", err)
	}

	var (
		channels   = 3 // Using only 3 channels due to no alpha channel data
		pixelIndex = 0
	)

	for read := 0; read < int(sprite.Size); {
		var transparentPixels uint16
		var coloredPixels uint16

		if err := binary.Read(f, binary.LittleEndian, &transparentPixels); err != nil {
			return nil, fmt.Errorf("could not read transparent pixels (%v)", err)
		}

		if err := binary.Read(f, binary.LittleEndian, &coloredPixels); err != nil {
			return nil, fmt.Errorf("could not read colored pixels (%v)", err)
		}

		for i := 0; i < int(transparentPixels); i++ {
			sprite.Pixels[pixelIndex] = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
			pixelIndex++
		}

		for i := 0; i < int(coloredPixels); i++ {
			rgb := new(color.RGBA)

			if err := binary.Read(f, binary.LittleEndian, &rgb.R); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			if err := binary.Read(f, binary.LittleEndian, &rgb.G); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			if err := binary.Read(f, binary.LittleEndian, &rgb.B); err != nil {
				return nil, fmt.Errorf("could not read colored pixel data: %v", err)
			}

			sprite.Pixels[pixelIndex] = color.RGBA{R: rgb.R, G: rgb.G, B: rgb.B, A: 0xFF}
			pixelIndex++
		}

		read += 4 + (channels * int(coloredPixels))
	}

	return sprite, nil
}

// LoadAll reads the full .spr file and saves the sprites found in it
func (spriteManager *SpriteManager) LoadAll() error {
	f, err := os.Open(spriteManager.filePath)

	if err != nil {
		return fmt.Errorf("could not open file %v", err)
	}

	defer f.Close()

	spriteFile := new(SpriteFile)

	binary.Read(f, binary.LittleEndian, &spriteFile.Signature)
	binary.Read(f, binary.LittleEndian, &spriteFile.Count)

	spriteManager.SpriteFile.Loaded = true

	// Sprites start with the 2 offset
	spriteManager.Sprites = make(map[uint32]*Sprite, spriteFile.Count)

	for spriteID := 0; spriteID <= int(spriteFile.Count); spriteID++ {
		logrus.Infof("Reading sprite ID %v", spriteID)
		if spr, err := spriteManager.LoadSpriteByID(spriteID); err == nil {
			spriteManager.Sprites[uint32(spriteID)] = spr
		}
	}

	return nil
}
