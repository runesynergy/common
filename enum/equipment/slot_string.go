// Code generated by "stringer -type Slot -trimprefix Slot"; DO NOT EDIT.

package equipment

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SlotHead-0]
	_ = x[SlotBack-1]
	_ = x[SlotNeck-2]
	_ = x[SlotRightHand-3]
	_ = x[SlotBody-4]
	_ = x[SlotLeftHand-5]
	_ = x[SlotHairReserved-6]
	_ = x[SlotLegs-7]
	_ = x[SlotBeardReserved-8]
	_ = x[SlotHands-9]
	_ = x[SlotFeet-10]
	_ = x[SlotArmsReserved-11]
	_ = x[SlotRing-12]
	_ = x[SlotQuiver-13]
	_ = x[SlotCount-14]
	_ = x[SlotBothHands-15]
}

const _Slot_name = "HeadBackNeckRightHandBodyLeftHandHairReservedLegsBeardReservedHandsFeetArmsReservedRingQuiverCountBothHands"

var _Slot_index = [...]uint8{0, 4, 8, 12, 21, 25, 33, 45, 49, 62, 67, 71, 83, 87, 93, 98, 107}

func (i Slot) String() string {
	if i < 0 || i >= Slot(len(_Slot_index)-1) {
		return "Slot(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Slot_name[_Slot_index[i]:_Slot_index[i+1]]
}
