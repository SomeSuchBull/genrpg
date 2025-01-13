package shadowdark

import (
	"fmt"
	"math"
	"math/rand"
	"slices"

	"github.com/genrpg/utils"
)

var stockingEngine = []func(int, ...Biome) string{getEmpty, getEmpty, getTrap, getMinorHazard, getSoloMonster, getNPC, getMonsterMob, getMajorHazard, getTreasure, getBossMonster}

func Stocking(rooms int64, level int) {
	biome := GetBiomes()[utils.TableDie(len(GetBiomes()))]
	fmt.Println("Biome:", biome)
	for i := 1; i < int(rooms)+1; i++ {
		roomContents := ""
		roll := utils.TableDie(10)
		f := stockingEngine[roll]
		roomContents += f(level, biome)
		// roomContents += getNPC(level, biome)
		fmt.Printf("%s %s\n", utils.RoomStyle(fmt.Sprintf("%03d", i)), roomContents)
		// fmt.Printf("%3d %s\n", i, roomContents)
	}
}

func getEmpty(level int, biomes ...Biome) string {
	return utils.B("Empty")
}

func getTrap(level int, biomes ...Biome) string {
	var trap string
	tables := [3]string{"Trap", "Trigger", "Effect"}
	for i, v := range traps {
		roll := rand.Intn(len(v))
		trap += fmt.Sprintf("%s: %s", tables[i], v[roll])
		if i < len(traps)-1 {
			trap += " | "
		}
	}
	var trapDetail string
	for i := 0; i < 2; i++ {
		roll := rand.Intn(len(mapsTrap[i]))
		trapDetail += fmt.Sprintf("%s", mapsTrap[i][roll])
		if i < 1 {
			trapDetail += " "
		}
	}
	return fmt.Sprintf("%s, %s\n%s", utils.B("Trap"), trapDetail, trap)
}

func getTreasure(level int, biomes ...Biome) string {
	return "Treasure"
}

// func getSpecial(level int) string {
// 	return "Special"
// }

func getMinorHazard(level int, biomes ...Biome) string {
	var hazard string
	tables := [3]string{"Movement", "Damage", "Weaken"}
	roll := utils.TableDie(3)
	tableName := tables[roll]
	table := hazards[roll]
	hazard = fmt.Sprintf("%s: %s", tableName, table[utils.TableDie(len(table))])
	hazardDetail := fmt.Sprintf("%s", mapsMinorHazard[rand.Intn(len(mapsMinorHazard))])
	return fmt.Sprintf("%s, %s\n%s", utils.B("Minor Hazard"), hazardDetail, hazard)
}

func getSoloMonster(level int, biomes ...Biome) string {
	var encounter string
	threatLevel := level + 2
	monsters := FilterMonstersByLevelAndBiome(threatLevel, threatLevel, biomes[0])
	var initialMonster IsMonster
	if len(monsters) == 0 {
		initialMonster = MonsterGenerator(level)
	} else {
		initialMonster = monsters[rand.Intn(len(monsters))]
	}
	encounter += fmt.Sprintf("%s", initialMonster)
	return fmt.Sprintf("%s, %s %s\n%s", utils.B("Solo Monster"), mapsSoloMonster[0][rand.Intn(6)], mapsSoloMonster[1][rand.Intn(6)], encounter)
}

func getNPC(level int, biomes ...Biome) string {
	tableRoll := utils.TableDie(6)
	detail := mapsNpc[tableRoll]
	var encounter string
	if tableRoll == 5 {
		encounter = fmt.Sprintf("%s", NewRivalParty(level))
	} else {
		encounter = fmt.Sprintf("%s", NewNPC())
	}
	return fmt.Sprintf("%s, %s\n%s", utils.B("NPC"), detail, encounter)
}

func getMonsterMob(level int, biomes ...Biome) string {
	var encounter string
	wholeThreat := monsterMath(level, MobMonsterType)
	maxLevelThreat := int(math.Floor(float64(wholeThreat) / 2.0))
	if maxLevelThreat > level {
		maxLevelThreat = level
	}
	minLevelThreat := int(math.Ceil(float64(wholeThreat) / 8.0))
	if level == 1 {
		minLevelThreat = 0
	}
	monsters := FilterMonstersByLevelAndBiome(minLevelThreat, maxLevelThreat, biomes[0])
	var initialMonster IsMonster
	if len(monsters) == 0 {
		initialMonster = MonsterGenerator(level)
	} else {
		initialMonster = monsters[rand.Intn(len(monsters))]
	}
	numberOfMonsters := 0
	threatFulfilled := 0
	for {
		numberOfMonsters++
		threatFulfilled += initialMonster.GetLevel()
		if initialMonster.GetLevel() == 0 {
			threatFulfilled++
		}
		if !((threatFulfilled+initialMonster.GetLevel() <= wholeThreat) && (threatFulfilled+1 <= wholeThreat)) {
			break
		}
	}
	encounter += fmt.Sprintf("%d %s", numberOfMonsters, initialMonster)
	if threatFulfilled < wholeThreat {
		threatLevel := wholeThreat - threatFulfilled
		monsters = FilterMonstersByLevelAndBiome(threatLevel, threatLevel, biomes[0])
		var monster IsMonster
		if len(monsters) == 0 {
			monster = MonsterGenerator(level)
		} else {
			monster = monsters[rand.Intn(len(monsters))]
		}
		encounter += fmt.Sprintf(" &\n%s", monster)
	}
	return fmt.Sprintf("%s, %s %s\n%s", utils.B("Monster Mob"), mapsMonsterMob[0][rand.Intn(6)], mapsMonsterMob[1][rand.Intn(6)], encounter)
}

func getMajorHazard(level int, biomes ...Biome) string {
	var hazard string
	tables := [3]string{"Movement", "Damage", "Weaken"}
	roll := utils.TableDie(3)
	tableName := tables[roll]
	table := hazards[roll]
	hazard = fmt.Sprintf("%s: %s", tableName, table[utils.TableDie(len(table))])
	hazardDetail := fmt.Sprintf("%s", mapsMajorHazard[rand.Intn(len(mapsMajorHazard))])
	return fmt.Sprintf("%s, %s\n%s", utils.B("Major Hazard"), hazardDetail, hazard)
}

func getBossMonster(level int, biomes ...Biome) string {
	var encounter string
	minThreatLevel := level + 4
	maxThreatLevel := 2 * level
	if maxThreatLevel > 20 {
		maxThreatLevel = 30
	}
	if maxThreatLevel < minThreatLevel {
		maxThreatLevel = minThreatLevel
	}
	detailsRoll := utils.TableDie(6)
	hasHelp := false
	if slices.Contains([]int{2, 3, 4}, detailsRoll) {
		hasHelp = true
		maxThreatLevel -= 2
		if maxThreatLevel < minThreatLevel {
			minThreatLevel = maxThreatLevel
		}
	}
	monsters := FilterMonstersByLevelAndBiome(minThreatLevel, maxThreatLevel, biomes[0])
	var initialMonster IsMonster
	if len(monsters) == 0 {
		initialMonster = MonsterGenerator(level)
	} else {
		initialMonster = monsters[rand.Intn(len(monsters))]
	}
	encounter += fmt.Sprintf("%s", initialMonster)
	if hasHelp {
		wholeThreat := monsterMath(level, MobMonsterType) - initialMonster.GetLevel()
		if wholeThreat <= 0 {
			wholeThreat = 3
		}
		maxLevelThreat := int(math.Floor(float64(wholeThreat) / 2.0))
		if maxLevelThreat > level {
			maxLevelThreat = level
		}
		minLevelThreat := int(math.Ceil(float64(wholeThreat) / 8.0))
		if level == 1 {
			minLevelThreat = 0
		}
		monsters = FilterMonstersByLevelAndBiome(minLevelThreat, maxLevelThreat, biomes[0])
		if len(monsters) == 0 {
			initialMonster = MonsterGenerator(level)
		} else {
			initialMonster = monsters[rand.Intn(len(monsters))]
		}
		numberOfMonsters := 0
		threatFulfilled := 0
		for {
			numberOfMonsters++
			threatFulfilled += initialMonster.GetLevel()
			if initialMonster.GetLevel() == 0 {
				threatFulfilled++
			}
			if !((threatFulfilled+initialMonster.GetLevel() <= wholeThreat) && (threatFulfilled+1 <= wholeThreat)) {
				break
			}
		}
		encounter += fmt.Sprintf(" &\n%d %s", numberOfMonsters, initialMonster)
	}
	return fmt.Sprintf("%s, %s\n%s", utils.B("Boss Monster"), mapsBossMonster[detailsRoll], encounter)
}
