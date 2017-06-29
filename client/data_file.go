package client

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/drgomesp/go-tibiac/client/thing"
)

const (
	ground           = 0x00
	groundBorder     = 0x01
	onBottom         = 0x02
	onTop            = 0x03
	container        = 0x04
	stackable        = 0x05
	forceUse         = 0x06
	multiUse         = 0x07
	writable         = 0x08
	writableOnce     = 0x09
	fluidContainer   = 0x0A
	fluid            = 0x0B
	unpassable       = 0x0C
	unmoveable       = 0x0D
	blockMissile     = 0x0E
	blockPathfinding = 0x0F
	noMoveAnimation  = 0x10
	pickupable       = 0x11
	hangable         = 0x12
	hookSouth        = 0x13
	hookEasth        = 0x14
	rotatable        = 0x15
	hasLight         = 0x16
	dontHide         = 0x17
	translucent      = 0x18
	hasOffset        = 0x19
	hasElevation     = 0x1A
	lyingObject      = 0x1B
	animateAlways    = 0x1C
	miniMap          = 0x1D
	lensHelp         = 0x1E
	fullGround       = 0x1F
	ignoreLook       = 0x20
	cloth            = 0x21
	marketItem       = 0x22
	defaultAction    = 0x23
	wrappable        = 0x24
	unwrappable      = 0x25
	topEffect        = 0x26
	hasCharges       = 0xFC
	floorChange      = 0xFD
	usable           = 0xFE
	end              = 0xFF
)

// Item represents a Tibia Item
type Item struct {
}

// DataFile represents Tibia's .dat file
type DataFile struct {
	Signature                                          uint32
	ItemCount, OutfitCount, EffectCount, DistanceCount uint16
	Things                                             map[uint32]*thing.Type

	path string
}

// NewDataFileFromPath opens, parses and returns a new DataFile from a .dat file path
func NewDataFileFromPath(path string) (*DataFile, error) {
	dataFile := &DataFile{
		path: path,
	}

	f, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("could not open file %v", err)
	}

	defer f.Close()

	binary.Read(f, binary.LittleEndian, &dataFile.Signature)
	binary.Read(f, binary.LittleEndian, &dataFile.ItemCount)
	binary.Read(f, binary.LittleEndian, &dataFile.OutfitCount)
	binary.Read(f, binary.LittleEndian, &dataFile.EffectCount)
	binary.Read(f, binary.LittleEndian, &dataFile.DistanceCount)

	dataFile.Things = make(map[uint32]*thing.Type, dataFile.ItemCount)

	for i := 100; i < 101; i++ {
		logrus.WithFields(logrus.Fields{"ID": i}).Info("item")

		var (
			flag      byte
			thingType = &thing.Type{}
		)

		binary.Read(f, binary.LittleEndian, &flag)

		for flag != end {
			binary.Read(f, binary.LittleEndian, &flag)

			logrus.Infof("reading flag %#X\n", flag)

			switch flag {
			case end:
				{
					continue
				}
			case ground:
				{
					thingType.IsGround = true
					binary.Read(f, binary.LittleEndian, &thingType.GroundSpeed)
				}
			case groundBorder:
				{
					thingType.IsGroundBorder = true
				}
			case onBottom:
				{
					thingType.IsOnBottom = true
				}
			case onTop:
				{
					thingType.IsOnTop = true
				}
			case container:
				{
					thingType.IsContainer = true
				}
			case stackable:
				{
					thingType.IsStackable = true
				}
			case forceUse:
				{
					thingType.IsForceUse = true
				}
			case multiUse:
				{
					thingType.IsMultiUse = true
				}
			case writable:
				{
					thingType.IsWritable = true
					binary.Read(f, binary.LittleEndian, &thingType.MaxTextLength)
				}
			case writableOnce:
				{
					thingType.IsWritableOnce = true
					binary.Read(f, binary.LittleEndian, &thingType.MaxTextLength)
				}
			case fluidContainer:
				{
					thingType.IsFluidContainer = true
				}
			case fluid:
				{
					thingType.IsFluid = true
				}
			case unpassable:
				{
					thingType.IsUnpassable = true
				}
			case unmoveable:
				{
					thingType.IsUnmoveable = true
				}
			case blockMissile:
				{
					thingType.IsBlockMissile = true
				}
			case blockPathfinding:
				{
					thingType.IsBlockPathfinding = true
				}
			case noMoveAnimation:
				{
					thingType.IsNoMoveAnimation = true
				}
			case pickupable:
				{
					thingType.IsPickupable = true
				}
			case hangable:
				{
					thingType.IsHangable = true
				}
			case hookSouth:
				{
					thingType.IsHookSouth = true
				}
			case hookEasth:
				{
					thingType.IsHookEast = true
				}
			case rotatable:
				{
					thingType.IsRotatable = true
				}
			case hasLight:
				{
					thingType.HasLight = true
					binary.Read(f, binary.LittleEndian, &thingType.LightLevel)
					binary.Read(f, binary.LittleEndian, &thingType.LightColor)
				}
			case dontHide:
				{
					thingType.IsDontHide = true
				}
			case translucent:
				{
					thingType.IsTranslucent = true
				}
			case hasOffset:
				{
					thingType.HasOffset = true
					binary.Read(f, binary.LittleEndian, &thingType.OffsetX)
					binary.Read(f, binary.LittleEndian, &thingType.OffsetY)
				}
			case hasElevation:
				{
					thingType.HasElevation = true
				}
			case lyingObject:
				{
					thingType.IsLyingObject = true
				}
			case animateAlways:
				{
					thingType.IsAnimateAlways = true
				}
			case miniMap:
				{
					thingType.IsMiniMap = true
					binary.Read(f, binary.LittleEndian, &thingType.MiniMapColor)
				}
			case lensHelp:
				{
					thingType.IsLensHelp = true
					binary.Read(f, binary.LittleEndian, &thingType.LensHelp)
				}
			case fullGround:
				{
					thingType.IsFullGround = true
				}
			case ignoreLook:
				{
					thingType.IsIgnoreLook = true
				}
			case cloth:
				{
					thingType.IsCloth = true
					binary.Read(f, binary.LittleEndian, &thingType.ClothSlot)
				}
			case marketItem:
				{
					thingType.IsMarketItem = true
					binary.Read(f, binary.LittleEndian, &thingType.MarketCategory)
					binary.Read(f, binary.LittleEndian, &thingType.MarketTradeAs)
					binary.Read(f, binary.LittleEndian, &thingType.MarketShowAs)
					binary.Read(f, binary.LittleEndian, &thingType.MarketNameLength)
					binary.Read(f, binary.LittleEndian, &thingType.MarketName)
					binary.Read(f, binary.LittleEndian, &thingType.MarketRestrictProfession)
					binary.Read(f, binary.LittleEndian, &thingType.MarketRestrictLevel)
				}
			case defaultAction:
				{
					thingType.HasDefaultAction = true
					binary.Read(f, binary.LittleEndian, &thingType.DefaultAction)
				}
			case wrappable:
				{
					thingType.Wrappable = true
				}
			case unwrappable:
				{
					thingType.Unwrappable = true
				}
			case topEffect:
				{
					thingType.TopEffect = true
				}
			case hasCharges:
				{
					thingType.HasCharges = true
				}
			case floorChange:
				{
					thingType.FloorChange = true
				}
			case usable:
				{
					thingType.Usable = true
				}
			default:
				{
					logrus.Panicf("unknown flag %#X\n", flag)
				}
			}
		}

		dataFile.Things[uint32(i)] = thingType
	}

	return dataFile, nil
}
