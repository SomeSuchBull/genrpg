package shadowdark

import "github.com/genrpg/utils"

var monsterDistributionMath = []map[int]int{
	{2: 2, 3: 3, 4: 3, 5: 3, 6: 4, 7: 4, 8: 4, 9: 5, 10: 5, 11: 5, 12: 6},
	{2: 6, 3: 9, 4: 10, 5: 11, 6: 12, 7: 12, 8: 12, 9: 13, 10: 14, 11: 15, 12: 18},
	{2: 10, 3: 15, 4: 16, 5: 18, 6: 19, 7: 20, 8: 22, 9: 23, 10: 24, 11: 25, 12: 30},
	{2: 14, 3: 21, 4: 23, 5: 25, 6: 27, 7: 28, 8: 30, 9: 32, 10: 34, 11: 35, 12: 42},
}

func monsterMath(level int, monsterType MonsterType) int {
	switch level {
	case 1:
		return monsterDistributionMath[0][weightedTwoToTwelve(-1, monsterType)]
	case 2:
		return monsterDistributionMath[0][weightedTwoToTwelve(0, monsterType)]
	case 3:
		return monsterDistributionMath[0][weightedTwoToTwelve(1, monsterType)]
	case 4:
		return monsterDistributionMath[1][weightedTwoToTwelve(-1, monsterType)]
	case 5:
		return monsterDistributionMath[1][weightedTwoToTwelve(0, monsterType)]
	case 6:
		return monsterDistributionMath[1][weightedTwoToTwelve(1, monsterType)]
	case 7:
		return monsterDistributionMath[2][weightedTwoToTwelve(-1, monsterType)]
	case 8:
		return monsterDistributionMath[2][weightedTwoToTwelve(0, monsterType)]
	case 9:
		return monsterDistributionMath[2][weightedTwoToTwelve(1, monsterType)]
	case 10:
		return monsterDistributionMath[3][weightedTwoToTwelve(0, monsterType)]
	default:
		return monsterDistributionMath[3][weightedTwoToTwelve(1, monsterType)]
	}
}

func weightedTwoToTwelve(adjustment int, monsterType MonsterType) int {
	res := utils.D(6) + utils.D(6) + adjustment
	if res < 2 {
		res = 2
	}
	if res > 12 {
		res = 12
	}
	switch {
	case monsterType == MobMonsterType && res > 10:
		res = 10
	case monsterType == SoloMonsterType && res < 3:
		res = 3
	case monsterType == SoloMonsterType && res > 11:
		res = 11
	case monsterType == BossMonsterType && res < 7:
		res = 7
	}
	return res
}
