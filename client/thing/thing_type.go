package thing

// Type represents a Tibia thing
type Type struct {
	IsGround    bool
	GroundSpeed uint16

	IsGroundBorder,
	IsOnBottom,
	IsOnTop,
	IsContainer,
	IsStackable,
	IsForceUse,
	IsMultiUse,
	IsFluidContainer,
	IsFluid,
	IsUnpassable,
	IsUnmoveable,
	IsBlockMissile,
	IsBlockPathfinding,
	IsNoMoveAnimation,
	IsPickupable,
	IsHangable,
	IsHookSouth,
	IsHookEast,
	IsRotatable,
	IsDontHide,
	IsTranslucent,
	IsLyingObject,
	IsAnimateAlways,
	IsFullGround,
	IsIgnoreLook,
	HasCharges,
	FloorChange,
	Wrappable,
	Unwrappable,
	TopEffect,
	Usable bool

	IsWritable, IsWritableOnce bool
	MaxTextLength              uint16

	HasLight               bool
	LightLevel, LightColor uint16

	HasOffset        bool
	OffsetX, OffsetY uint16

	HasElevation bool
	Elevation    uint16

	IsMiniMap    bool
	MiniMapColor uint16

	IsLensHelp bool
	LensHelp   uint16

	IsCloth   bool
	ClothSlot uint16

	IsMarketItem bool
	MarketCategory,
	MarketTradeAs,
	MarketShowAs,
	MarketNameLength,
	MarketName,
	MarketRestrictProfession,
	MarketRestrictLevel uint16

	HasDefaultAction bool
	DefaultAction    uint16
}
