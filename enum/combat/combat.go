package combat

//go:generate stringer -type Action -trimprefix Action
type Action uint8

const (
	ActionInvalid Action = iota
	ActionStab
	ActionSlash
	ActionCrush
	ActionCast
	ActionFire
	ActionDefend
	ActionEquip
	ActionAttack
	ActionDie
)

func (a Action) DamageType() DamageType {
	switch a {
	case ActionStab:
		return DamageTypeStab
	case ActionSlash:
		return DamageTypeSlash
	case ActionCrush:
		return DamageTypeCrush
	case ActionCast:
		return DamageTypeMagic
	case ActionFire:
		return DamageTypeRanged
	}
	return DamageTypeNone
}

//go:generate stringer -type Style -trimprefix Style
type Style uint8

const (
	StyleInvalid Style = iota
	StyleAccurate
	StyleAggressive
	StyleDefensive
	StyleControlled
	StyleRapid
	StyleLongrange
	StyleMagic
	StyleAutocast
)

//go:generate stringer -type DamageType -trimprefix DamageType
type DamageType uint16

const (
	DamageTypeNone DamageType = iota
	DamageTypeStab DamageType = 1 << (iota - 1)
	DamageTypeSlash
	DamageTypeCrush
	DamageTypeMagic
	DamageTypeRanged
	DamageTypePoison
	DamageTypeAir
	DamageTypeWater
	DamageTypeEarth
	DamageTypeFire
	DamageTypeDragonfire
)

func (d DamageType) Contains(other DamageType) bool {
	return (d & other) == other
}

//go:generate stringer -type Splat -trimprefix Splat
type Splat uint8

const (
	SplatMiss Splat = iota
	SplatHit
	SplatPoison
)

// Option is used to represent the combat settings of a clicked combat option.
type Option struct {
	// Name is only for debugging and documentation purposes.
	Name string
	// Style can be used to modify the speed or accuracy of a weapon.
	Style Style
	// Action determines the animation, sound, and damage type of this option.
	Action Action
	// DamageType is an addition to the damage type provided by the Action.DamageType().
	DamageType DamageType
}