package stat

import "fmt"

// Level is a value between 0 and 255.
type Level uint8

// Experience is tracked as xp*10 to allow fractional values like 2.6 or 1.1, but cannot be more accurate
// than 0.1. The standard should be to always end an experience literal with _0 if you expect a whole number.
// e.g: 5xp is 5_0 and 5.6xp is 5_6. There should never be more than one digit after the _.
type Experience uint32

type Stat struct {
	Level      Level
	Experience Experience
}

const MaxSkillExperience = 200000000_0

var ExperienceTable = [...]Experience{
	0:  0,
	1:  0,
	2:  83_0,
	3:  174_0,
	4:  276_0,
	5:  388_0,
	6:  512_0,
	7:  650_0,
	8:  801_0,
	9:  969_0,
	10: 1154_0,
	11: 1358_0,
	12: 1584_0,
	13: 1833_0,
	14: 2107_0,
	15: 2411_0,
	16: 2746_0,
	17: 3115_0,
	18: 3523_0,
	19: 3973_0,
	20: 4470_0,
	21: 5018_0,
	22: 5624_0,
	23: 6291_0,
	24: 7028_0,
	25: 7842_0,
	26: 8740_0,
	27: 9730_0,
	28: 10824_0,
	29: 12031_0,
	30: 13363_0,
	31: 14833_0,
	32: 16456_0,
	33: 18247_0,
	34: 20224_0,
	35: 22406_0,
	36: 24815_0,
	37: 27473_0,
	38: 30408_0,
	39: 33648_0,
	40: 37224_0,
	41: 41171_0,
	42: 45529_0,
	43: 50339_0,
	44: 55649_0,
	45: 61512_0,
	46: 67983_0,
	47: 75127_0,
	48: 83014_0,
	49: 91721_0,
	50: 101333_0,
	51: 111945_0,
	52: 123660_0,
	53: 136594_0,
	54: 150872_0,
	55: 166636_0,
	56: 184040_0,
	57: 203254_0,
	58: 224466_0,
	59: 247886_0,
	60: 273742_0,
	61: 302288_0,
	62: 333804_0,
	63: 368599_0,
	64: 407015_0,
	65: 449428_0,
	66: 496254_0,
	67: 547953_0,
	68: 605032_0,
	69: 668051_0,
	70: 737627_0,
	71: 814445_0,
	72: 899257_0,
	73: 992895_0,
	74: 1096278_0,
	75: 1210421_0,
	76: 1336443_0,
	77: 1475581_0,
	78: 1629200_0,
	79: 1798808_0,
	80: 1986068_0,
	81: 2192818_0,
	82: 2421087_0,
	83: 2673114_0,
	84: 2951373_0,
	85: 3258594_0,
	86: 3597792_0,
	87: 3972294_0,
	88: 4385776_0,
	89: 4842295_0,
	90: 5346332_0,
	91: 5902831_0,
	92: 6517253_0,
	93: 7195629_0,
	94: 7944614_0,
	95: 8771558_0,
	96: 9684577_0,
	97: 10692629_0,
	98: 11805606_0,
	99: 13034431_0,
}

// Level uses a binary search to find the level for a given Experience.
func (e Experience) Level() Level {
	var l = 1
	var r = len(ExperienceTable)
	for l < r {
		m := (l + r) / 2
		if ExperienceTable[m] > e {
			r = m
		} else {
			l = m + 1
		}
	}
	return Level(r - 1)
}

func (e Experience) String() string {
	return fmt.Sprintf("%.2f", float64(e)/10)
}