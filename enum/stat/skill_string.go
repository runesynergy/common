// Code generated by "stringer -type Skill"; DO NOT EDIT.

package stat

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SkillAttack-0]
	_ = x[SkillDefence-1]
	_ = x[SkillStrength-2]
	_ = x[SkillHitpoints-3]
	_ = x[SkillRanged-4]
	_ = x[SkillPrayer-5]
	_ = x[SkillMagic-6]
	_ = x[SkillCooking-7]
	_ = x[SkillWoodcutting-8]
	_ = x[SkillFletching-9]
	_ = x[SkillFishing-10]
	_ = x[SkillFiremaking-11]
	_ = x[SkillCrafting-12]
	_ = x[SkillSmithing-13]
	_ = x[SkillMining-14]
	_ = x[SkillHerblore-15]
	_ = x[SkillAgility-16]
	_ = x[SkillThieving-17]
	_ = x[SkillSlayer-18]
	_ = x[SkillFarming-19]
	_ = x[SkillRunecrafting-20]
	_ = x[SkillCount-21]
}

const _Skill_name = "SkillAttackSkillDefenceSkillStrengthSkillHitpointsSkillRangedSkillPrayerSkillMagicSkillCookingSkillWoodcuttingSkillFletchingSkillFishingSkillFiremakingSkillCraftingSkillSmithingSkillMiningSkillHerbloreSkillAgilitySkillThievingSkillSlayerSkillFarmingSkillRunecraftingSkillCount"

var _Skill_index = [...]uint16{0, 11, 23, 36, 50, 61, 72, 82, 94, 110, 124, 136, 151, 164, 177, 188, 201, 213, 226, 237, 249, 266, 276}

func (i Skill) String() string {
	if i >= Skill(len(_Skill_index)-1) {
		return "Skill(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Skill_name[_Skill_index[i]:_Skill_index[i+1]]
}