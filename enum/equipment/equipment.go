package equipment

//go:generate stringer -type Slot -trimprefix Slot
type Slot int

const (
	SlotHead Slot = iota
	SlotBack
	SlotNeck
	SlotRightHand
	SlotBody
	SlotLeftHand
	SlotHairReserved
	SlotLegs
	SlotBeardReserved
	SlotHands
	SlotFeet
	SlotArmsReserved
	SlotRing
	SlotQuiver
	SlotCount

	SlotBothHands // used for 2h weapons

	SlotCountWithAppearance = SlotCount - 2 // - ring - quiver
)

func (s Slot) HasAppearance() bool {
	return !(s == SlotRing || s == SlotQuiver)
}