package client

import (
	"encoding/binary"
	"fmt"
	"os"
)

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
func (spriteManager *SpriteManager) LoadSpriteByID(spriteID int64) (*Sprite, error) {
	file, err := NewSpriteFileFromPath(spriteManager.filePath)

	if err != nil {
		return nil, fmt.Errorf("could not create sprite file from path: %v", err)
	}

	return file.LoadSpriteByID(spriteID)
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
		if spr, err := spriteManager.LoadSpriteByID(int64(spriteID)); err == nil {
			spriteManager.Sprites[uint32(spriteID)] = spr
		}
	}

	return nil
}
