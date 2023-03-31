package stat

//go:generate stringer -type Skill
type Skill uint8

const (
	SkillAttack Skill = iota
	SkillDefence
	SkillStrength
	SkillHitpoints
	SkillRanged
	SkillPrayer
	SkillMagic
	SkillCooking
	SkillWoodcutting
	SkillFletching
	SkillFishing
	SkillFiremaking
	SkillCrafting
	SkillSmithing
	SkillMining
	SkillHerblore
	SkillAgility
	SkillThieving
	SkillSlayer
	SkillFarming
	SkillRunecrafting
	SkillCount
)

func (s Skill) Enabled() bool {
	return s != SkillFarming
}