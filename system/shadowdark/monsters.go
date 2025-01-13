package shadowdark

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/genrpg/utils"
)

type MonsterTag string

func (t MonsterTag) String() string {
	return string(t)
}

const (
	AberrationTag  MonsterTag = "aberration"
	AngelTag       MonsterTag = "angel"
	AnimalTag      MonsterTag = "animal"
	CelestialTag   MonsterTag = "celestial"
	ConstructTag   MonsterTag = "construct"
	DemonTag       MonsterTag = "demon"
	DevilTag       MonsterTag = "devil"
	DinosaurTag    MonsterTag = "dinosaur"
	DireTag        MonsterTag = "dire"
	DragonTag      MonsterTag = "dragon"
	ElementalTag   MonsterTag = "elemental"
	FeyTag         MonsterTag = "fey"
	FiendTag       MonsterTag = "fiend"
	GiantTag       MonsterTag = "giant"
	GolemTag       MonsterTag = "golem"
	HumanoidTag    MonsterTag = "humanoid"
	InsectTag      MonsterTag = "insect"
	LegendaryTag   MonsterTag = "legendary"
	MonstrosityTag MonsterTag = "monstrosity"
	OozeTag        MonsterTag = "ooze"
	OutsiderTag    MonsterTag = "outsider"
	PlantTag       MonsterTag = "plant"
	SwarmTag       MonsterTag = "swarm"
	UndeadTag      MonsterTag = "undead"
)

type MonsterQuality string

func (q MonsterQuality) String() string {
	return string(q)
}

const (
	BeastlikeQuality   MonsterQuality = "Beastlike"
	AvianQuality       MonsterQuality = "Avian"
	AmphibiousQuality  MonsterQuality = "Amphibious"
	DemonicQuality     MonsterQuality = "Demonic"
	ArachnidQuality    MonsterQuality = "Arachnid"
	OozeQuality        MonsterQuality = "Ooze"
	InsectoidQuality   MonsterQuality = "Insectoid"
	DraconicQuality    MonsterQuality = "Draconic"
	PlantlikeQuality   MonsterQuality = "Plantlike"
	ElephantineQuality MonsterQuality = "Elephantine"
	UndeadQuality      MonsterQuality = "Undead"
	CrystallineQuality MonsterQuality = "Crystalline"
	HumanoidQuality    MonsterQuality = "Humanoid"
	AngelicQuality     MonsterQuality = "Angelic"
	SpectralQuality    MonsterQuality = "Spectral"
	StonecarvedQuality MonsterQuality = "Stonecarved"
	SerpentineQuality  MonsterQuality = "Serpentine"
	ElementalQuality   MonsterQuality = "Elemental"
	PiscineQuality     MonsterQuality = "Piscine"
	ReptilianQuality   MonsterQuality = "Reptilian"
)

func GetMonsterQuality() MonsterQuality {
	return []MonsterQuality{BeastlikeQuality, AvianQuality, AmphibiousQuality, DemonicQuality, ArachnidQuality, OozeQuality, InsectoidQuality, DraconicQuality, PlantlikeQuality, ElephantineQuality, UndeadQuality, CrystallineQuality, HumanoidQuality, AngelicQuality, SpectralQuality, StonecarvedQuality, SerpentineQuality, ElementalQuality, PiscineQuality, ReptilianQuality}[utils.TableDie(20)]
}

type MonsterStrength string

func (q MonsterStrength) String() string {
	return string(q)
}

const (
	Attack1Strength           MonsterStrength = "+1 attack"
	AbsorbsMagicStrength      MonsterStrength = "Absorbs magic"
	SwarmStrength             MonsterStrength = "Swarm"
	D10DamageStrength         MonsterStrength = "1d10 damage"
	PoisonStingStrength       MonsterStrength = "Poison sting"
	ConfusingGazeStrength     MonsterStrength = "Confusing gaze"
	EatsMetalStrength         MonsterStrength = "Eats metal"
	RangedAttacksStrength     MonsterStrength = "Ranged attacks"
	HighlyIntelligentStrength MonsterStrength = "Highly intelligent"
	CrushingGraspStrength     MonsterStrength = "Crushing grasp"
	PsychicBlastStrength      MonsterStrength = "Psychic blast"
	StealthyStrength          MonsterStrength = "Stealthy"
	PetrifyingGazeStrength    MonsterStrength = "Petrifying gaze"
	D12DamageStrength         MonsterStrength = "1d12 damage"
	ImpersonationStrength     MonsterStrength = "Impersonation"
	BlindingAuraStrength      MonsterStrength = "Blinding aura"
	TurnsInvisibleStrength    MonsterStrength = "Turns invisible"
	TwoD6DamageStrength       MonsterStrength = "2d6 damage"
	SwallowsWholeStrength     MonsterStrength = "Swallows whole"
	Attack2Strength           MonsterStrength = "+2 attacks"
)

func GetMonsterStrength() MonsterStrength {
	return []MonsterStrength{Attack1Strength, AbsorbsMagicStrength, SwarmStrength, D10DamageStrength, PoisonStingStrength, ConfusingGazeStrength, EatsMetalStrength, RangedAttacksStrength, HighlyIntelligentStrength, CrushingGraspStrength, PsychicBlastStrength, StealthyStrength, PetrifyingGazeStrength, D12DamageStrength, ImpersonationStrength, BlindingAuraStrength, TurnsInvisibleStrength, TwoD6DamageStrength, SwallowsWholeStrength, Attack2Strength}[utils.TableDie(20)]
}

type MonsterWeakness string

func (q MonsterWeakness) String() string {
	return string(q)
}

const (
	ColdWeakness        MonsterWeakness = "Cold"
	GreedyWeakness      MonsterWeakness = "Greedy"
	LightWeakness       MonsterWeakness = "Light"
	SaltWeakness        MonsterWeakness = "Salt"
	VainWeakness        MonsterWeakness = "Vain"
	MirrorsWeakness     MonsterWeakness = "Mirrors"
	ElectricityWeakness MonsterWeakness = "Electricity"
	FragileBodyWeakness MonsterWeakness = "Fragile body"
	SunlightWeakness    MonsterWeakness = "Sunlight"
	SilverWeakness      MonsterWeakness = "Silver"
	FireWeakness        MonsterWeakness = "Fire"
	FoodWeakness        MonsterWeakness = "Food"
	AcidWeakness        MonsterWeakness = "Acid"
	GarlicWeakness      MonsterWeakness = "Garlic"
	IronWeakness        MonsterWeakness = "Iron"
	WaterWeakness       MonsterWeakness = "Water"
	ItsTrueNameWeakness MonsterWeakness = "Its True Name"
	LoudSoundsWeakness  MonsterWeakness = "Loud sounds"
	HolyWaterWeakness   MonsterWeakness = "Holy water"
	MusicWeakness       MonsterWeakness = "Music"
)

func GetMonsterWeakness() MonsterWeakness {
	return []MonsterWeakness{ColdWeakness, GreedyWeakness, LightWeakness, SaltWeakness, VainWeakness, MirrorsWeakness, ElectricityWeakness, FragileBodyWeakness, SunlightWeakness, SilverWeakness, FireWeakness, FoodWeakness, AcidWeakness, GarlicWeakness, IronWeakness, WaterWeakness, ItsTrueNameWeakness, LoudSoundsWeakness, HolyWaterWeakness, MusicWeakness}[utils.TableDie(20)]
}

type MonsterMutation string

func (q MonsterMutation) String() string {
	return string(q)
}

const (
	ShapechangerMutation      MonsterMutation = "Shapechanger"
	DoubleDamageMutation      MonsterMutation = "Double damage"
	SpeaksCommonMutation      MonsterMutation = "Speaks Common"
	FinsGillsMutation         MonsterMutation = "Fins and gills"
	BreathesFireMutation      MonsterMutation = "Breathes fire"
	SpellsMutation            MonsterMutation = "Knows 1d4 spells"
	InsulatingFurMutation     MonsterMutation = "Insulating fur"
	FastHealingMutation       MonsterMutation = "Fast healing"
	TelepathicMutation        MonsterMutation = "Telepathic"
	IronlikeScalesMutation    MonsterMutation = "Ironlike scales"
	Attack1Mutation           MonsterMutation = "+1 attack"
	ToxicSporesMutation       MonsterMutation = "Toxic spores"
	ExtraLimbsMutation        MonsterMutation = "Extra limbs"
	AC2Mutation               MonsterMutation = "+2 AC"
	SonicBlastsMutation       MonsterMutation = "Sonic blasts"
	TentaclesMutation         MonsterMutation = "Tentacles"
	Levels2Mutation           MonsterMutation = "+2 levels"
	BurstTeleportsMutation    MonsterMutation = "Can teleport in bursts"
	BonelessMutation          MonsterMutation = "Boneless"
	PlusDamageMutation        MonsterMutation = "+1d6 damage"
	ParalyticMutation         MonsterMutation = "Paralytic touch"
	GiganticMutation          MonsterMutation = "Gigantic"
	LifeDrainingMutation      MonsterMutation = "Life-draining touch"
	GeniusMutation            MonsterMutation = "Genius intellect"
	FlingsSpikesMutation      MonsterMutation = "Flings spikes"
	VeryFastMutation          MonsterMutation = "Very fast"
	AntimagicMutation         MonsterMutation = "Antimagic field"
	TwoHeadsMutation          MonsterMutation = "Two heads"
	ReflectsSpellsMutation    MonsterMutation = "Reflects spells"
	BloodDrainingMutation     MonsterMutation = "Blood-draining bite"
	BurrowsMutation           MonsterMutation = "Burrows"
	ElectrifiedWeaponMutation MonsterMutation = "Electrified weapon"
	FeverMutation             MonsterMutation = "Has swamp fever"
	WingsMutation             MonsterMutation = "Wings"
	AcidicMutation            MonsterMutation = "Acidic saliva"
	BlessedMutation           MonsterMutation = "Blessed by a god"
)

func GetMonsterMutation(num int) MonsterMutation {
	switch num {
	case 1:
		return []MonsterMutation{ShapechangerMutation, FinsGillsMutation, InsulatingFurMutation, IronlikeScalesMutation, ExtraLimbsMutation, TentaclesMutation, BonelessMutation, GiganticMutation, FlingsSpikesMutation, TwoHeadsMutation, BurrowsMutation, WingsMutation}[utils.TableDie(12)]
	case 2:
		return []MonsterMutation{DoubleDamageMutation, BreathesFireMutation, FastHealingMutation, Attack1Mutation, AC2Mutation, Levels2Mutation, PlusDamageMutation, LifeDrainingMutation, VeryFastMutation, ReflectsSpellsMutation, ElectrifiedWeaponMutation, AcidicMutation}[utils.TableDie(12)]
	default:
		return []MonsterMutation{SpeaksCommonMutation, SpellsMutation, TelepathicMutation, ToxicSporesMutation, SonicBlastsMutation, BurstTeleportsMutation, ParalyticMutation, GeniusMutation, AntimagicMutation, BloodDrainingMutation, FeverMutation, BlessedMutation}[utils.TableDie(12)]
	}
}

type MonsterType int

const (
	MobMonsterType MonsterType = iota - 1
	SoloMonsterType
	BossMonsterType
)

type Biome string

func (b Biome) String() string {
	return string(b)
}

const (
	ArcticBiome     Biome = "arctic"
	CaveBiome       Biome = "cave"
	DeepsBiome      Biome = "deeps"
	DesertBiome     Biome = "desert"
	ForestBiome     Biome = "forest"
	GrasslandBiome  Biome = "grassland"
	JungleBiome     Biome = "jungle"
	MountainBiome   Biome = "mountain"
	OceanBiome      Biome = "ocean"
	RiverCoastBiome Biome = "river/coast"
	RuinsBiome      Biome = "ruins"
	SwampBiome      Biome = "swamp"
	TombBiome       Biome = "tomb"
)

func GetBiomes() []Biome {
	return []Biome{ArcticBiome, CaveBiome, DeepsBiome, DesertBiome, ForestBiome, GrasslandBiome, JungleBiome, MountainBiome, OceanBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}
}

type MonsterAlignment string

func (a MonsterAlignment) String() string {
	return string(a)
}

const (
	LawfulAlignment  MonsterAlignment = "l"
	NeutralAlignment MonsterAlignment = "n"
	ChaoticAlignment MonsterAlignment = "c"
)

type IsMonster interface {
	String() string
	GetLevel() int
}

type Monster struct {
	Name             string           `json:"name"`
	Level            int              `json:"level"`
	MonsterAlignment MonsterAlignment `json:"alignment"`
	Move             string           `json:"move"`
	Attack           string           `json:"attack"`
	Page             string           `json:"page"`
	StatBlock        string           `json:"stat_block"`
	Tags             []MonsterTag     `json:"tags"`
	Biomes           []Biome          `json:"biomes"`
}

func (m Monster) String() string {
	return fmt.Sprintf("%s | LV: %d (p. %s)", m.Name, m.Level, m.Page)
}

func (m Monster) GetLevel() int {
	return m.Level
}

func ShowDistribution() {
	monsters := GetAllMonsters()
	distribution := make(map[int]int)
	for _, monster := range monsters {
		if _, ok := distribution[monster.Level]; !ok {
			distribution[monster.Level] = 0
		}
		distribution[monster.Level]++
	}
	orderedLevels := make([][2]int, len(distribution), len(distribution))
	for level, count := range distribution {
		for i, orderedLevel := range orderedLevels {
			if orderedLevel[1] == 0 {
				orderedLevels[i] = [2]int{level, count}
				break
			}
			if level < orderedLevel[0] {
				for ii := len(orderedLevels) - 1; ii > i; ii-- {
					orderedLevels[ii] = orderedLevels[ii-1]
				}
				orderedLevels[i] = [2]int{level, count}
				break
			}
		}
	}
	for i, orderedLevel := range orderedLevels {
		fmt.Printf("%02d -> Level %d: %d\n", i, orderedLevel[0], orderedLevel[1])
	}
}

type RandomMonster struct {
	AC, Level, Combat, AttackNumber int
	DamageDie                       string
	Quality                         MonsterQuality
	Strength                        MonsterStrength
	Weakness                        MonsterWeakness
	Mutation1, Mutation2, Mutation3 MonsterMutation
}

func (m RandomMonster) String() string {
	output := fmt.Sprintf("%s | LV: %d | AC: %d | ATK: %d × %d (%s) | Quality: %s | Strength: %s | Weakness: %s",
		utils.B("Random Monster"), m.Level, m.AC, m.AttackNumber, m.Combat, m.DamageDie, m.Quality, m.Strength, m.Weakness)
	if m.Mutation1 != "" {
		output += fmt.Sprintf("\nMutations: %s", m.Mutation1)
	}
	if m.Mutation2 != "" {
		output += fmt.Sprintf(" | %s", m.Mutation2)
	}
	if m.Mutation3 != "" {
		output += fmt.Sprintf(" | %s", m.Mutation3)
	}
	return output
}

func (m RandomMonster) GetLevel() int {
	return m.Level
}

func MonsterGenerator(level int) RandomMonster {
	monster := RandomMonster{
		AC: 10 + level,
	}
	monster.Level = []int{level - 3, level - 3, level - 2, level - 2, level - 1, level - 1, level, level, level, level,
		level, level, level, level + 1, level + 1, level + 2, level + 2, level + 3, level + 3, level + 4}[utils.TableDie(20)]
	monster.Combat = monster.Level
	if monster.Level < 0 {
		monster.Level = 0
	}
	switch {
	case level <= 3:
		monster.Mutation1 = GetMonsterMutation(1)
		if utils.D(3) == 3 {
			monster.AttackNumber = 2
		} else {
			monster.AttackNumber = 1
		}
		if utils.D(3) == 1 {
			monster.DamageDie = "d4"
		} else {
			monster.DamageDie = "d6"
		}
	case level <= 6:
		monster.Mutation1 = GetMonsterMutation(1)
		monster.Mutation2 = GetMonsterMutation(2)
		if utils.D(3) == 3 {
			monster.AttackNumber = 3
		} else {
			monster.AttackNumber = 2
		}
		if utils.D(3) == 1 {
			monster.DamageDie = "d6"
		} else {
			monster.DamageDie = "d8"
		}
	case level <= 9:
		monster.Mutation1 = GetMonsterMutation(1)
		monster.Mutation2 = GetMonsterMutation(2)
		monster.Mutation3 = GetMonsterMutation(3)
		if utils.D(3) == 3 {
			monster.AttackNumber = 3
		} else {
			monster.AttackNumber = 4
		}
		if utils.D(3) == 1 {
			monster.DamageDie = "d8"
		} else {
			monster.DamageDie = "d10"
		}
	case level >= 10:
		monster.Mutation1 = GetMonsterMutation(1)
		monster.Mutation2 = GetMonsterMutation(2)
		monster.Mutation3 = GetMonsterMutation(3)
		if utils.D(3) == 3 {
			monster.AttackNumber = 4
		} else {
			monster.AttackNumber = 5
		}
		monster.DamageDie = "d12"
	}
	monster.Quality = GetMonsterQuality()
	monster.Strength = GetMonsterStrength()
	monster.Weakness = GetMonsterWeakness()
	return monster
}

func FilterMonstersByLevel(level int) []Monster {
	monsters := GetAllMonsters()
	filteredMonsters := make([]Monster, 0)
	for _, monster := range monsters {
		if monster.Level == level {
			filteredMonsters = append(filteredMonsters, monster)
		}
	}
	return filteredMonsters
}

func FilterMonstersByBiome(biome Biome) []Monster {
	monsters := GetAllMonsters()
	filteredMonsters := make([]Monster, 0)
	for _, monster := range monsters {
		if monster.Biomes == nil {
			filteredMonsters = append(filteredMonsters, monster)
			continue
		}
		for _, monsterBiome := range monster.Biomes {
			if monsterBiome == biome {
				filteredMonsters = append(filteredMonsters, monster)
				break
			}
		}
	}
	return filteredMonsters
}

func FilterMonstersByTag(tag MonsterTag) []Monster {
	monsters := GetAllMonsters()
	filteredMonsters := make([]Monster, 0)
	for _, monster := range monsters {
		for _, monsterTag := range monster.Tags {
			if monsterTag == tag {
				filteredMonsters = append(filteredMonsters, monster)
				break
			}
		}
	}
	return filteredMonsters
}

func FilterMonstersByAlignment(alignment MonsterAlignment) []Monster {
	monsters := GetAllMonsters()
	filteredMonsters := make([]Monster, 0)
	for _, monster := range monsters {
		if monster.MonsterAlignment == alignment {
			filteredMonsters = append(filteredMonsters, monster)
		}
	}
	return filteredMonsters
}

func FilterMonstersByLevelAndBiome(minLevel, maxLevel int, biome Biome) []Monster {
	monsters := GetAllMonsters()
	filteredMonsters := make([]Monster, 0)
	for _, monster := range monsters {
		if monster.Level >= minLevel && monster.Level <= maxLevel {
			if monster.Biomes == nil {
				filteredMonsters = append(filteredMonsters, monster)
				continue
			}
			for _, monsterBiome := range monster.Biomes {
				if monsterBiome == biome {
					filteredMonsters = append(filteredMonsters, monster)
					break
				}
			}
		}
	}
	return filteredMonsters
}

func GetAllMonsters() []Monster {
	// ShadowDark
	monsters := []Monster{Aboleth, Acolyte, AngelDomini, AngelPrincipi, AngelSeraph, AnimatedArmor, Ankheg, Ape, ApeSnow, Apprentice, Archangel, Archdevil, Archmage, Assassin, Azer, Badger, Bandit, Basilisk, BatGiant, BatSwarm, BearBrown, BearPolar, Beastman, Berserker, BlackPudding, Boar, Brachiosaurus, BrainEater, Bugbear, Bulette, Camel, CaveBrute, CaveCreeper, Centaur, CentipedeGiant, CentipedeSwarm, Chimera, Chuul, Cloaker, Cockatrice, Couatl, CrabGiant, Crocodile, Cultist, Cyclops, Darkmantle, DeepOne, DemonBalor, DemonDretch, DemonGlabrezu, DemonMarilith, DemonVrock, DevilBarbed, DevilCubi, DevilErinyes, DevilHorned, DevilImp, Djinni, Doppelganger, DragonDesert, DragonFire, DragonForest, DragonFrost, DragonSea, DragonSwamp, Drow, DrowDrider, DrowPriestess, Druid, Dryad, Duergar, DungBeetleGiant, Efreeti, ElementalAir6HD, ElementalAir9HD, ElementalEarth6HD, ElementalEarth9HD, ElementalFire6HD, ElementalFire9HD, ElementalWater6HD, ElementalWater9HD, Elephant, Elf, Ettercap, Fairy, FrogGiant, Gargoyle, GelatinousCube, Ghast, Ghost, Ghoul, GiantCloud, GiantFire, GiantFrost, GiantGoat, GiantHill, GiantStone, GiantStorm, GibberingMouther, Gladiator, Gnoll, GnomeDeep, Goblin, GoblinBoss, GoblinShaman, GolemClay, GolemFlesh, GolemIron, GolemStone, Gorgon, Gorilla, GrayOoze, Grick, Griffon, Grimlow, Guard, HagNight, HagSea, HagWeald, Harpy, HellHound, Hippogriff, Hippopotamus, Hobgoblin, Horse, Hydra2HD, Hydra4HD, Hydra6HD, Hydra8HD, Hydra10HD, Hydra12HD, Hydra14HD, Hydra16HD, InvisibleStalker, Jellyfish, Knight, Kobold, KoboldSorcerer, Kraken, LeechGiant, Leprechaun, Lich, Lion, Lizardfolk, Mage, Mammoth, MantaRayGiant, Manticore, Mastiff, Medusa, Merfolk, Mimic, Minotaur, Moose, MordanticusTheFlayed, Mummy, Mushroomfolk, Naga, NagaBone, Nightmare, ObeIxxOfAzarumme, OchreJelly, OctopusGiant, Ogre, Oni, Orc, OrcChieftain, Otyugh, Owlbear, Panther, Peasant, Pegasus, Phoenix, PiranhaSwarm, Pirate, Plesiosaurus, Priest, PrimordialSlime, Pterodactyl, PurpleWorm, Rakshasa, Rat, RatDire, RatGiant, RatSwarm, Rathgamnon, Reaver, Remorhaz, Rhinoceros, RimeWalker, Roc, Roper, RotFlower, RustMonster, Sahuagin, Salamander, ScarabSwarm, Scarecrow, Scorpion, ScorpionGiant, Shadow, ShamblingMound, Shark, SharkMegalodon, Siren, Skeleton, Smilodon, SnakeCobra, SnakeGiant, SnakeSwarm, Soldier, Sphinx, Spider, SpiderGiant, SpiderSwarm, Stingbat, Strangler, Tarrasque, TenEyedOracle, WanderingMerchant, Thief, Thug, Treant, Triceratops, Troll, TrollFrost, Tyrannosaurus, Unicorn, Vampire, VampireSpawn, Velociraptor, VioletFungus, Viperian, ViperianOphid, ViperianWizard, VoidSpawn, VoidSpider, Vulture, WaspGiant, Wererat, Werewolf, Wight, WillOWisp, Wolf, WolfDire, WolfWinter, Worg, Wraith, Wyvern, Zombie}
	// Cursed Scroll 1-3
	monsters = append(monsters, []Monster{Bittermold, Bogthorn, Dralech, GordockBreeg, Hexling, Howler, IchorOoze, MarrowFiend, Mugdulblub, MutantCatfish, PlogrinaBittermold, Skrell, TarBat, Willowman, CamelSilver, CanyonApe, Donkey, Dunefiend, DustDevil, HeroGladiator, HorseWar, Mirage, RasGodai, RookiePitFighter, Scrag, ScragWar, Siruul, Scourge, DrakeGreater, DrakeLesser, Draugr, Dverg, Nord, TrollDeep, SeaSerpent, SeaNymph, Orca, Oracle, Werebear, Valkyrie}...)
	return monsters
}

func LoadMonsters(w io.Writer) error {
	monsters, ok := monsters["entries"].([]map[string]string)
	if !ok {
		return errors.New("could not load monsters")
	}
	cursedScrollMonsters, ok := cursedScrollMonsters["entries"].([]map[string]string)
	if !ok {
		return errors.New("could not load monsters")
	}
	for _, monsters := range [][]map[string]string{monsters, cursedScrollMonsters} {
		for _, monster := range monsters {
			monsterNameParts := strings.Split(monster["name"], ", ")
			var monsterName string
			var monsterVarName string
			for i, partComaSeparated := range monsterNameParts {
				if i != 0 {
					monsterName += ", "
				}
				partsSpaceSeparated := strings.Split(partComaSeparated, " ")
				for ii, partSpaceSeparated := range partsSpaceSeparated {
					if ii != 0 {
						monsterName += " "
					}
					partsMinusSeparated := strings.Split(partSpaceSeparated, "-")
					for iii, part := range partsMinusSeparated {
						if iii != 0 {
							monsterName += "-"
						}
						if i != 0 || ii != 0 || iii != 0 || part != "THE" {
							monsterVarName += fmt.Sprintf("%c%s", part[0], strings.ToLower(part[1:]))
						}
						monsterName += fmt.Sprintf("%c%s", part[0], strings.ToLower(part[1:]))
					}
				}
			}
			// monsterName = string(monster["name"][0]) + strings.ToLower(monster["name"][1:])
			// monsterVarName = string(monster["name"][0]) + strings.ToLower(monster["name"][1:])
			var alignment string
			switch monster["alignment"] {
			case "L":
				alignment = "LawfulAlignment"
			case "N":
				alignment = "NeutralAlignment"
			case "C":
				alignment = "ChaoticAlignment"
			}
			monsterLevel := monster["level"]
			if monsterLevel == "*" {
				monsterLevel = "0"
			}
			inputString := fmt.Sprintf(
				"var %s = Monster{Name: \"%s\",Level: %s,MonsterAlignment: %s,Move: \"%s\",Attack: \"%s\",Page: \"%s\",StatBlock: \"%s\",",
				monsterVarName, monsterName, monster["level"], alignment, monster["move"], monster["attack"], monster["page"], monster["statblock"])
			biomes := strings.Split(monster["biome"], ",")
			for i, biome := range biomes {
				if i == 0 && biome != "*" {
					inputString += "Biomes: []Biome{"
				}
				switch biome {
				case "arctic":
					inputString += fmt.Sprintf("%s, ", "ArcticBiome")
				case "cave":
					inputString += fmt.Sprintf("%s, ", "CaveBiome")
				case "deeps":
					inputString += fmt.Sprintf("%s, ", "DeepsBiome")
				case "desert":
					inputString += fmt.Sprintf("%s, ", "DesertBiome")
				case "forest":
					inputString += fmt.Sprintf("%s, ", "ForestBiome")
				case "grassland":
					inputString += fmt.Sprintf("%s, ", "GrasslandBiome")
				case "jungle":
					inputString += fmt.Sprintf("%s, ", "JungleBiome")
				case "mountain":
					inputString += fmt.Sprintf("%s, ", "MountainBiome")
				case "ocean":
					inputString += fmt.Sprintf("%s, ", "OceanBiome")
				case "river/coast":
					inputString += fmt.Sprintf("%s, ", "RiverCoastBiome")
				case "ruins":
					inputString += fmt.Sprintf("%s, ", "RuinsBiome")
				case "swamp":
					inputString += fmt.Sprintf("%s, ", "SwampBiome")
				case "tomb":
					inputString += fmt.Sprintf("%s, ", "TombBiome")
				}
				if i == len(biomes)-1 && biome != "*" {
					inputString += "},"
				}
			}
			tags := strings.Split(monster["tags"], ",")
			for i, tag := range tags {
				if i == 0 && tag != "*" {
					inputString += "Tags: []MonsterTag{"
				}
				switch tag {
				case "aberration":
					inputString += fmt.Sprintf("%s, ", "AberrationTag")
				case "angel":
					inputString += fmt.Sprintf("%s, ", "AngelTag")
				case "animal":
					inputString += fmt.Sprintf("%s, ", "AnimalTag")
				case "celestial":
					inputString += fmt.Sprintf("%s, ", "CelestialTag")
				case "construct":
					inputString += fmt.Sprintf("%s, ", "ConstructTag")
				case "demon":
					inputString += fmt.Sprintf("%s, ", "DemonTag")
				case "devil":
					inputString += fmt.Sprintf("%s, ", "DevilTag")
				case "dinosaur":
					inputString += fmt.Sprintf("%s, ", "DinosaurTag")
				case "dire":
					inputString += fmt.Sprintf("%s, ", "DireTag")
				case "dragon":
					inputString += fmt.Sprintf("%s, ", "DragonTag")
				case "elemental":
					inputString += fmt.Sprintf("%s, ", "ElementalTag")
				case "fey":
					inputString += fmt.Sprintf("%s, ", "FeyTag")
				case "fiend":
					inputString += fmt.Sprintf("%s, ", "FiendTag")
				case "giant":
					inputString += fmt.Sprintf("%s, ", "GiantTag")
				case "golem":
					inputString += fmt.Sprintf("%s, ", "GolemTag")
				case "humanoid":
					inputString += fmt.Sprintf("%s, ", "HumanoidTag")
				case "insect":
					inputString += fmt.Sprintf("%s, ", "InsectTag")
				case "legendary":
					inputString += fmt.Sprintf("%s, ", "LegendaryTag")
				case "monstrosity":
					inputString += fmt.Sprintf("%s, ", "MonstrosityTag")
				case "ooze":
					inputString += fmt.Sprintf("%s, ", "OozeTag")
				case "outsider":
					inputString += fmt.Sprintf("%s, ", "OutsiderTag")
				case "plant":
					inputString += fmt.Sprintf("%s, ", "PlantTag")
				case "swarm":
					inputString += fmt.Sprintf("%s, ", "SwarmTag")
				case "undead":
					inputString += fmt.Sprintf("%s, ", "UndeadTag")
				}
				if i == len(tags)-1 && tag != "*" {
					inputString += "},"
				}
			}
			inputString += "}"
			fmt.Fprintln(w, inputString)
		}
	}
	return nil
}

// ShadowDark
var Aboleth = Monster{Name: "Aboleth", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "194", StatBlock: "AC 16, HP 39, ATK 2 tentacle (near) +5 (1d8 + curse) or 1 tail +5 (3d6), MV near (swim), S +4, D -1, C +3, I +4, W +2, Ch +2, AL C, LV 8", Biomes: []Biome{DeepsBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AberrationTag}}
var Acolyte = Monster{Name: "Acolyte", Level: 1, MonsterAlignment: LawfulAlignment, Move: "", Attack: "spell", Page: "194", StatBlock: "AC 12, HP 4, ATK 1 mace +1 (1d6) or 1 spell +2, MV near, S +1, D -1, C+0, I -1, W +2, Ch +0, AL L, LV 1", Biomes: []Biome{RuinsBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var AngelDomini = Monster{Name: "Angel, Domini", Level: 9, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "", Page: "195", StatBlock: "AC 17 (plate mail + shield), HP 42, ATK 3 bastard sword +7 (1d8) or 1 horn, MV near (fly), S +4, D +1, C+2, I +3, W +4, Ch +4, AL L, LV 9", Tags: []MonsterTag{CelestialTag, AngelTag}}
var AngelPrincipi = Monster{Name: "Angel, Principi", Level: 11, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "", Page: "195", StatBlock: "AC 16 (+1 plate mail), HP 53, ATK 3 silvered bastard sword +9 (1d10), MV double near (fly), S +4, D +2, C +4, I +4, W +4, Ch +4, AL L, LV 11", Tags: []MonsterTag{CelestialTag, AngelTag}}
var AngelSeraph = Monster{Name: "Angel, Seraph", Level: 3, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "", Page: "195", StatBlock: "AC 14 (chainmail), HP 14, ATK 2 longsword +3 (1d8), MV near (fly), S +3, D +1, C +1, I +2, W +3, Ch +3, AL L, LV 3", Tags: []MonsterTag{CelestialTag, AngelTag}}
var AnimatedArmor = Monster{Name: "Animated Armor", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "196", StatBlock: "AC 15, HP 11, ATK 1 longsword +3 (1d8), MV near, S +3, D -1, C +2, I -1, W +1, Ch +0, AL C, LV 2", Biomes: []Biome{RuinsBiome, TombBiome}, Tags: []MonsterTag{ConstructTag}}
var Ankheg = Monster{Name: "Ankheg", Level: 3, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "196", StatBlock: "AC 14, HP 14, ATK 1 bite +4 (1d6) or 1 acid spray (near) +4 (2d6), MV near (burrow), S +2, D +2, C+1, I -2, W +1, Ch -2, AL N, LV 3", Biomes: []Biome{CaveBiome, DeepsBiome, DesertBiome, GrasslandBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Ape = Monster{Name: "Ape", Level: 2, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "196", StatBlock: "AC 12, HP 10, ATK 1 fist +2 (1d6) or 1 rock (far) +2 (1d4), MV near (climb), S +2, D +2, C +1, I -2, W +1, Ch +0, AL N, LV 2", Biomes: []Biome{ForestBiome, MountainBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag}}
var ApeSnow = Monster{Name: "Ape, Snow", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "196", StatBlock: "AC 13, HP 19, ATK 2 fist +4 (1d6) or 1 rock (far) +4 (2d6), MV near (climb), S +3, D +1, C +1, I -2, W +1, Ch +0, AL N, LV 4", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{AnimalTag}}
var Apprentice = Monster{Name: "Apprentice", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "spell", Page: "196", StatBlock: "AC 11, HP 3, ATK 1 dagger (close/near) +1 (1d4) or 1 spell +2, MV near, S -1, D +1, C -1, I +2, W +0, Ch+0, AL N, LV 1", Biomes: []Biome{ArcticBiome, CaveBiome, DesertBiome, ForestBiome, GrasslandBiome, JungleBiome, MountainBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Archangel = Monster{Name: "Archangel", Level: 16, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "fire", Page: "195", StatBlock: "AC 18 (+3 plate mail), HP 76, ATK3 flaming greatsword +10 (2d12), MV double near (fly), S +5, D +2, C+4, I +4, W +5, Ch +5, AL L, LV 16", Tags: []MonsterTag{CelestialTag, AngelTag}}
var Archdevil = Monster{Name: "Archdevil", Level: 16, MonsterAlignment: ChaoticAlignment, Move: "teleport", Attack: "", Page: "206", StatBlock: "AC 19, HP 76, ATK 4 iron scepter +10 (3d10) or 1 soulbind, MV far (teleport), S +5, D +4, C +4, I +5, W +4, Ch +7, AL C, LV 16", Tags: []MonsterTag{FiendTag, DevilTag}}
var Archmage = Monster{Name: "Archmage", Level: 10, MonsterAlignment: LawfulAlignment, Move: "", Attack: "spell", Page: "197", StatBlock: "AC 12, HP 44, ATK 2 spell +7, MV near, S -1, D +2, C -1, I +4, W +2, Ch+1, AL L, LV 10", Tags: []MonsterTag{HumanoidTag}}
var Assassin = Monster{Name: "Assassin", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "poison", Page: "197", StatBlock: "AC 15 (leather), HP 38, ATK 2 poisoned dagger (close/near) +6 (2d4), MV near (climb), S +2, D +4, C +2, I +2, W +3, Ch +3, AL C, LV 8", Tags: []MonsterTag{HumanoidTag}}
var Azer = Monster{Name: "Azer", Level: 3, MonsterAlignment: LawfulAlignment, Move: "", Attack: "fire", Page: "197", StatBlock: "AC 15, HP 15, ATK 2 flaming warhammer +3 (1d10, ignites flammables) or 1 crossbow (far) +0 (1d6), MV near, S +3, D +0, C+2, I +0, W +0, Ch +0, AL L, LV 3", Biomes: []Biome{CaveBiome, DeepsBiome, MountainBiome}, Tags: []MonsterTag{ElementalTag}}
var Badger = Monster{Name: "Badger", Level: 1, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "197", StatBlock: "AC 11, HP 5, ATK 2 claw +2 (1d4), MV near (burrow), S +2, D +0, C+1, I -3, W +1, Ch -2, AL N, LV 1", Biomes: []Biome{ForestBiome}, Tags: []MonsterTag{AnimalTag}}
var Bandit = Monster{Name: "Bandit", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "197", StatBlock: "AC 13 (leather + shield), HP 4, ATK 1 club +1 (1d4) or 1 shortbow (far) +0 (1d4), MV near, S +1, D +0, C +0, I -1, W +0, Ch -1, AL C, LV 1", Biomes: []Biome{ArcticBiome, DesertBiome, MountainBiome, RiverCoastBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Basilisk = Monster{Name: "Basilisk", Level: 5, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "198", StatBlock: "AC 14, HP 25, ATK 2 bite +4 (2d6 + petrify), MV near, S +3, D +1, C +3, I -3, W +1, Ch -3, AL N, LV 5", Biomes: []Biome{CaveBiome, GrasslandBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var BatGiant = Monster{Name: "Bat, Giant", Level: 2, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "198", StatBlock: "AC 12, HP 9, ATK 1 bite +2 (1d6), MV near (fly), S -1, D +2, C +0, I -3, W +1, Ch -3, AL N, LV 2", Biomes: []Biome{CaveBiome, DeepsBiome, MountainBiome, RuinsBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var BatSwarm = Monster{Name: "Bat, Swarm", Level: 4, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "198", StatBlock: "AC 12, HP 18, ATK 3 bite +2 (1d6), MV near (fly), S -3, D +2, C +0, I -3, W +1, Ch -3, AL N, LV 4", Biomes: []Biome{CaveBiome, MountainBiome, RuinsBiome}, Tags: []MonsterTag{AnimalTag, SwarmTag}}
var BearBrown = Monster{Name: "Bear, Brown", Level: 5, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "198", StatBlock: "AC 13, HP 25, ATK 2 claw +4 (1d8), MV near (climb) S +4, D +1, C +3, I-2, W +1, Ch -2, AL N, LV 5", Biomes: []Biome{ForestBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var BearPolar = Monster{Name: "Bear, Polar", Level: 7, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "198", StatBlock: "AC 13, HP 34, ATK 2 claw +6 (2d6), MV near (climb), S +4, D +1, C +3, I-2, W +1, Ch -2, AL N, LV 7", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{AnimalTag}}
var Beastman = Monster{Name: "Beastman", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "198", StatBlock: "AC 12 (leather), HP 5, ATK 1 spear (close/near) +2 (1d6 + 1), MV near, S +2, D +1, C +1, I -2, W +1, Ch -1, AL C, LV 1", Biomes: []Biome{CaveBiome, DeepsBiome, MountainBiome, RuinsBiome}, Tags: []MonsterTag{HumanoidTag}}
var Berserker = Monster{Name: "Berserker", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "199", StatBlock: "AC 12 (leather), HP 10, ATK 1 greataxe +2 (1d10) or 1 spear (close/near) +2 (1d6), MV near, S+2, D +1, C +1, I +0, W +1, Ch +0, AL N, LV 2", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome, JungleBiome}, Tags: []MonsterTag{HumanoidTag}}
var BlackPudding = Monster{Name: "Black Pudding", Level: 6, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "199", StatBlock: "AC 9, HP 30, ATK 3 tentacle +4 (2d6), MV near (climb), S +2, D -1, C +3, I -4, W -3, Ch -4, AL N, LV 6", Biomes: []Biome{ArcticBiome, CaveBiome, DeepsBiome, RiverCoastBiome, RuinsBiome}, Tags: []MonsterTag{OozeTag}}
var Boar = Monster{Name: "Boar", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "199", StatBlock: "AC 12, HP 14, ATK 2 tusk +3 (1d6), MV near, S +3, D +0, C +1, I -2, W+1, Ch -2, AL N, LV 3", Biomes: []Biome{ForestBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var Brachiosaurus = Monster{Name: "Brachiosaurus", Level: 12, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "208", StatBlock: "AC 13, HP 57, ATK 3 stomp +7 (2d10), MV double near, S +6, D -1, C +3, I -3, W +1, Ch -3, AL N, LV 12", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{DinosaurTag}}
var BrainEater = Monster{Name: "Brain Eater", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "199", StatBlock: "AC 14 (leather), HP 36, ATK 4 tentacle +5 (1d8 + latch) or 1 mind blast or 1 mind control, MV near, S +2, D +3, C +0, I +4, W +2, Ch +4, AL C, LV 8", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{AberrationTag}}
var Bugbear = Monster{Name: "Bugbear", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "200", StatBlock: "AC 13 (leather + shield), HP 14, ATK 2 spiked mace +3 (1d6), MV near, S +3, D +0, C +1, I -1, W +0, Ch -2, AL C, LV 3", Biomes: []Biome{CaveBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Bulette = Monster{Name: "Bulette", Level: 8, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "200", StatBlock: "AC 17, HP 40, ATK 3 bite +5 (2d6) or 1 leap, MV near (burrow), S +5, D +1, C +4, I -3, W +1, Ch -2, AL N, LV 8", Biomes: []Biome{CaveBiome, DesertBiome, GrasslandBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Camel = Monster{Name: "Camel", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "200", StatBlock: "AC 10, HP 12, ATK 1 hoof +3 (1d6) or 1 spit (near) +0 (1d4), MV double near, S +3, D +0, C +3, I -2, W +1, Ch -3, AL N, LV 2", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{AnimalTag}}
var CaveBrute = Monster{Name: "Cave Brute", Level: 6, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "200", StatBlock: "AC 14, HP 28, ATK 2 claw +5 (1d8) and 1 mandible +5 (1d10), MV near (burrow), S +4, D +1, C +1, I-3, W +1, Ch -3, AL N, LV 6", Biomes: []Biome{CaveBiome, DeepsBiome, RuinsBiome}, Tags: []MonsterTag{InsectTag}}
var CaveCreeper = Monster{Name: "Cave Creeper", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "toxin", Page: "200", StatBlock: "AC 12, HP 18, ATK 1 bite +3 (1d6) and 1 tentacles +3 (1d8 + toxin), MV near (climb), S +2, D +2, C +0, I -3, W +1, Ch -3, AL N, LV 4", Biomes: []Biome{CaveBiome, DeepsBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{InsectTag}}
var Centaur = Monster{Name: "Centaur", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "201", StatBlock: "AC 12 (leather), HP 14, ATK 2 spear (close/near) +2 (1d6) or 1 longbow (far) +1 (1d8), MV double near, S +2, D +1, C +1, I +0, W +2, Ch +1, AL N, LV 3", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{MonstrosityTag}}
var CentipedeGiant = Monster{Name: "Centipede, Giant", Level: 1, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "201", StatBlock: "AC 11, HP 4, ATK 1 bite +1 (1d4 + poison), MV near (climb), S -3, D+1, C +0, I -4, W -3, Ch -4, AL N, LV 1", Biomes: []Biome{CaveBiome, ForestBiome, JungleBiome}, Tags: []MonsterTag{InsectTag, GiantTag}}
var CentipedeSwarm = Monster{Name: "Centipede, Swarm", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "201", StatBlock: "AC 11, HP 18, ATK 3 bite +1 (1d4 + poison), MV near (climb), S -3, D+1, C +0, I -4, W -3, Ch -4, AL N, LV 4", Biomes: []Biome{CaveBiome, ForestBiome, JungleBiome, TombBiome}, Tags: []MonsterTag{InsectTag, SwarmTag}}
var Chimera = Monster{Name: "Chimera", Level: 10, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "fire", Page: "201", StatBlock: "AC 16, HP 49, ATK 4 rend +7 (2d8) and 1 fire breath, MV double near (fly), S +5, D +4, C +4, I -3, W+2, Ch -1, AL C, LV 10", Biomes: []Biome{CaveBiome, GrasslandBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Chuul = Monster{Name: "Chuul", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "201", StatBlock: "AC 15, HP 25, ATK 2 pincer +4 (1d8 + grab), MV near (swim), S+3, D -1, C +3, I -1, W +1, Ch -2, AL C, LV 5", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{InsectTag}}
var Cloaker = Monster{Name: "Cloaker", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "202", StatBlock: "AC 13, HP 28, ATK 3 lash +4 (1d8) or 1 screech, MV near (fly), S +2, D +3, C +1, I +1, W +1, Ch +0, AL C, LV 6", Biomes: []Biome{CaveBiome, DeepsBiome}, Tags: []MonsterTag{AberrationTag}}
var Cockatrice = Monster{Name: "Cockatrice", Level: 3, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "202", StatBlock: "AC 11, HP 14, ATK 1 bite +1 (1d4 + petrify), MV near (fly), S -2, D +1, C+1, I -3, W +1, Ch -3, AL N, LV 3", Biomes: []Biome{CaveBiome, DeepsBiome, DesertBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Couatl = Monster{Name: "Couatl", Level: 9, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "poison", Page: "202", StatBlock: "AC 16, HP 42, ATK 3 bite +6 (2d6 + poison), MV near (fly), S +2, D +3, C +2, I +4, W +4, Ch +5, AL L, LV 9", Biomes: []Biome{JungleBiome}, Tags: []MonsterTag{CelestialTag}}
var CrabGiant = Monster{Name: "Crab, Giant", Level: 5, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "202", StatBlock: "AC 15, HP 24, ATK 2 pincer +4 (1d8 + crush), MV near (swim), S+3, D +0, C +2, I -3, W +0, Ch -3, AL N, LV 5", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var Crocodile = Monster{Name: "Crocodile", Level: 4, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "203", StatBlock: "AC 14, HP 20, ATK 2 bite +3 (1d8), MV near (swim), S +3, D +1, C +2, I-2, W +1, Ch -2, AL N, LV 4", Biomes: []Biome{JungleBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag}}
var Cultist = Monster{Name: "Cultist", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "203", StatBlock: "AC 14 (chainmail + shield), HP9, ATK 1 longsword +1 (1d8) or 1 spell +2, MV near, S +1, D -1, C +0, I-1, W +2, Ch +0, AL C, LV 2", Tags: []MonsterTag{HumanoidTag}}
var Cyclops = Monster{Name: "Cyclops", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "203", StatBlock: "AC 11 (leather), HP 38, ATK 2 greatclub +7 (2d8) or 1 rock (far) +5 (1d12), MV double near, S +5, D +0, C +2, I -1, W -2, Ch +0, AL C, LV 8", Biomes: []Biome{ForestBiome, GrasslandBiome, JungleBiome, MountainBiome}, Tags: []MonsterTag{HumanoidTag, GiantTag}}
var Darkmantle = Monster{Name: "Darkmantle", Level: 1, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "203", StatBlock: "AC 13, HP 4 ATK 1 bite +3 (1d4) or 1 darkness, MV near (fly), S -2, D+3, C +0, I -3, W +0, Ch -3, AL N, LV 1", Biomes: []Biome{CaveBiome, DeepsBiome, RuinsBiome}, Tags: []MonsterTag{AberrationTag}}
var DeepOne = Monster{Name: "Deep One", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "203", StatBlock: "AC 13, HP 10, ATK 2 spear (close/near) +2 (1d6), MV near (swim), S+2, D +1, C +1, I -2, W +0, Ch -2, AL C, LV 2", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{AberrationTag}}
var DemonBalor = Monster{Name: "Demon, Balor", Level: 16, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "fire", Page: "204", StatBlock: "AC 19, HP 77, ATK 3 greatsword +10 (2d12 + hellfire) and 1 fire whip (near) +10 (2d6 + grab), MV double near (fly), S +6, D +2, C +5, I +4, W +3, Ch +4, AL C, LV 16", Tags: []MonsterTag{FiendTag, DemonTag}}
var DemonDretch = Monster{Name: "Demon, Dretch", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "204", StatBlock: "AC 12, HP 11, ATK 1 claw +2 (1d6) or 1 gas, MV near, S +2, D +0, C +2, I -2, W -1, Ch -3, AL C, LV 2", Tags: []MonsterTag{FiendTag, DemonTag}}
var DemonGlabrezu = Monster{Name: "Demon, Glabrezu", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "204", StatBlock: "AC 15, HP 40, ATK 2 pincer +7 (2d8 + crush), MV near, S +4, D +1, C +4, I +3, W +2, Ch +2, AL C, LV 8", Tags: []MonsterTag{FiendTag, DemonTag}}
var DemonMarilith = Monster{Name: "Demon, Marilith", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "205", StatBlock: "AC 17 (plate mail), HP 43, ATK6 longsword +7 (1d8), MV near (climb), S +5, D +4, C +3, I +3, W+3, Ch +4, AL C, LV 9", Tags: []MonsterTag{FiendTag, DemonTag}}
var DemonVrock = Monster{Name: "Demon, Vrock", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "205", StatBlock: "AC 14, HP 24, ATK 2 talons +4 (1d8) or 1 screech, MV near (fly), S+2, D +2, C +2, I -1, W +1, Ch +0, AL C, LV 5", Tags: []MonsterTag{FiendTag, DemonTag}}
var DevilBarbed = Monster{Name: "Devil, Barbed", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "fire", Page: "206", StatBlock: "AC 13, HP 14, ATK 2 spine (near) +3 (1d6 + barb) or 1 fire blast (far) +3 (1d8), MV near, S +2, D +3, C +1, I +1, W +1, Ch +1, AL C, LV 3", Tags: []MonsterTag{FiendTag, DevilTag}}
var DevilCubi = Monster{Name: "Devil, Cubi", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "207", StatBlock: "AC 14, HP 29, ATK 1 kiss +4 (1d6 + drain) or 1 charm, MV near (fly), S +2, D +4, C +2, I +3, W +2, Ch +5, AL C, LV 6", Tags: []MonsterTag{FiendTag, DevilTag}}
var DevilErinyes = Monster{Name: "Devil, Erinyes", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "poison", Page: "207", StatBlock: "AC 17 (+1 plate mail), HP 43, ATK 3 greatsword +8 (1d12) or 2 longbow (far) +8 (1d8 + poison), MV double near (fly), S +4, D +4, C +3, I +4, W +4, Ch +5, AL C, LV 9", Tags: []MonsterTag{FiendTag, DevilTag}}
var DevilHorned = Monster{Name: "Devil, Horned", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "fire", Page: "207", StatBlock: "AC 16, HP 35, ATK 2 burning trident (near) +7 (2d6) or 1 fire blast (far) +4 (2d8), MV double near (fly), S +5, D +2, C +4, I +2, W+1, Ch +2, AL C, LV 7", Tags: []MonsterTag{FiendTag, DevilTag}}
var DevilImp = Monster{Name: "Devil, Imp", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "poison", Page: "207", StatBlock: "AC 13, HP 9, ATK 1 stinger +3 (1d4 + poison), MV near (fly), S -2, D +3, C +0, I +1, W +0, Ch +2, AL C, LV 2", Tags: []MonsterTag{FiendTag, DevilTag}}
var Djinni = Monster{Name: "Djinni", Level: 10, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "209", StatBlock: "AC 14, HP 48, ATK 3 scimitar +7 (1d12) or 1 whirlwind, MV double near (fly), S +4, D +4, C +3, I +4, W+3, Ch +3, AL N, LV 10", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{ElementalTag}}
var Doppelganger = Monster{Name: "Doppelganger", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "209", StatBlock: "AC 12, HP 20, ATK 1 dagger (close/near) +2 (1d4), MV near, S+1, D +2, C +2, I +1, W +0, Ch +4, AL C, LV 4", Tags: []MonsterTag{MonstrosityTag}}
var DragonDesert = Monster{Name: "Dragon, Desert", Level: 13, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "electric", Page: "210", StatBlock: "AC 17, HP 61, ATK 3 rend +9 (2d10) or 1 lightning breath, MV double near (fly), S +5, D +3, C +3, I +4, W+5, Ch +5, AL L, LV 13", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{DragonTag}}
var DragonFire = Monster{Name: "Dragon, Fire", Level: 17, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "fire", Page: "210", StatBlock: "AC 18, HP 80, ATK 4 rend +11 (2d12) or 1 fire breath, MV double near (fly), S +6, D +5, C +4, I +4, W+4, Ch +5, AL C, LV 17", Biomes: []Biome{CaveBiome, DesertBiome, MountainBiome}, Tags: []MonsterTag{DragonTag}}
var DragonForest = Monster{Name: "Dragon, Forest", Level: 12, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "poison", Page: "211", StatBlock: "AC 16, HP 58, ATK 3 rend +8 (2d8) or 1 poison breath, MV double near (fly), S +4, D +3, C +4, I +3, W+3, Ch +4, AL N, LV 12", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{DragonTag}}
var DragonFrost = Monster{Name: "Dragon, Frost", Level: 14, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "ice", Page: "211", StatBlock: "AC 17, HP 68, ATK 4 rend +9 (2d10) or 1 ice breath, MV double near (fly), S +4, D +3, C +5, I +3, W+4, Ch +3, AL N, LV 14", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{DragonTag}}
var DragonSea = Monster{Name: "Dragon, Sea", Level: 16, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "211", StatBlock: "AC 17, HP 76, ATK 4 rend +10 (2d10) or 1 steam breath or 1 water spout, MV double near (fly, swim), S +5, D +6, C +4, I +4, W+5, Ch +5, AL L, LV 16", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{DragonTag}}
var DragonSwamp = Monster{Name: "Dragon, Swamp", Level: 12, MonsterAlignment: ChaoticAlignment, Move: "burrow", Attack: "", Page: "211", StatBlock: "AC 16, HP 58, ATK 3 rend +8 (2d10) or 1 smog breath, MV double near (burrow, swim), S +5, D +3, C +4, I +4, W +3, Ch +3, AL C, LV 12", Biomes: []Biome{SwampBiome, RiverCoastBiome}, Tags: []MonsterTag{DragonTag}}
var Drow = Monster{Name: "Drow", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "poison", Page: "212", StatBlock: "AC 16 (mithral chainmail), HP 9, ATK 1 poison dart (near) +3 (1d4 + poison) or 1 longsword +1 (1d8), MV near, S +0, D +3, C +0, I +1, W+1, Ch +1, AL C, LV 2", Biomes: []Biome{CaveBiome, DeepsBiome, JungleBiome}, Tags: []MonsterTag{HumanoidTag}}
var DrowDrider = Monster{Name: "Drow, Drider", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "poison", Page: "212", StatBlock: "AC 16 (mithral chainmail), HP 29, ATK 3 longsword +3 (1d8) or 2 longbow (far) +3 (1d8 + poison), MV near (climb), S +3, D +3, C +2, I +2, W +2, Ch +0, AL C, LV 6", Biomes: []Biome{CaveBiome, DeepsBiome, JungleBiome}, Tags: []MonsterTag{MonstrosityTag}}
var DrowPriestess = Monster{Name: "Drow, Priestess", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell,poison", Page: "212", StatBlock: "AC 16 (mithral chainmail), HP 28, ATK 3 snake whip (near) +4 (1d8 + poison) or 1 spell +4, MV near, S +2, D +3, C +1, I +3, W +4, Ch +3, AL C, LV 6", Biomes: []Biome{CaveBiome, DeepsBiome, JungleBiome}, Tags: []MonsterTag{HumanoidTag}}
var Druid = Monster{Name: "Druid", Level: 7, MonsterAlignment: NeutralAlignment, Move: "", Attack: "spell", Page: "213", StatBlock: "AC 11, HP 31, ATK 1 staff +0 (1d4) or 2 spell +5, MV near, S +0, D +1, C +0, I +4, W +3, Ch +0, AL N, LV 7", Biomes: []Biome{ForestBiome, GrasslandBiome}, Tags: []MonsterTag{HumanoidTag}}
var Dryad = Monster{Name: "Dryad", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "213", StatBlock: "AC 13, HP 19, ATK 1 staff -1 (1d4) or 1 charm, MV near, S -1, D +2, C +1, I +1, W +3, Ch +4, AL N, LV 4", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{FeyTag}}
var Duergar = Monster{Name: "Duergar", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "213", StatBlock: "AC 15 (chainmail + shield), HP 12, ATK 1 war pick +2 (1d6), MV near, S +2, D +0, C +3, I +0, W -1, Ch -1, AL C, LV 2", Biomes: []Biome{CaveBiome, DeepsBiome}, Tags: []MonsterTag{HumanoidTag}}
var DungBeetleGiant = Monster{Name: "Dung Beetle, Giant", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "214", StatBlock: "AC 13, HP 10, ATK 1 horn +1 (1d4 + knock), MV near, S +1, D -1, C +1, I-3, W -1, Ch -3, AL N, LV 2", Biomes: []Biome{CaveBiome, GrasslandBiome, RuinsBiome}, Tags: []MonsterTag{InsectTag, GiantTag}}
var Efreeti = Monster{Name: "Efreeti", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "fire", Page: "214", StatBlock: "AC 15, HP 43, ATK 3 scimitar +8 (2d10) or 2 fire bolt (far) +5 (2d6), MV near (fly), S +5, D +2, C +3, I +3, W +2, Ch +3, AL C, LV 9", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{ElementalTag}}
var ElementalAir6HD = Monster{Name: "Elemental, Air (LV 6)", Level: 6, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "215", StatBlock: "AC 16, HP 29, ATK 3 slam +7 (2d6) or 1 whirlwind, MV double near (fly), S +3, D +5, C +2, I -2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalAir9HD = Monster{Name: "Elemental, Air (LV 9)", Level: 9, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "215", StatBlock: "AC 16, HP 42, ATK 3 slam +7 (3d6) or 1 whirlwind, MV double near (fly), S +3, D +5, C +2, I -2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalEarth6HD = Monster{Name: "Elemental, Earth (LV 6)", Level: 6, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "215", StatBlock: "AC 17, HP 31, ATK 3 slam +7 (2d8) or 1 avalanche, MV near (burrow), S +5, D +0, C +4, I-2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalEarth9HD = Monster{Name: "Elemental, Earth (LV 9)", Level: 9, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "215", StatBlock: "AC 17, HP 44, ATK 3 slam +7 (3d8) or 1 avalanche, MV near (burrow), S +5, D +0, C +4, I-2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalFire6HD = Monster{Name: "Elemental, Fire (LV 6)", Level: 6, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "fire", Page: "215", StatBlock: "AC 15, HP 30, ATK 3 slam +6 (2d10) or 1 inferno, MV near (fly), S +4, D +3, C +3, I -2, W +1, Ch-2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalFire9HD = Monster{Name: "Elemental, Fire (LV 9)", Level: 9, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "fire", Page: "215", StatBlock: "AC 15, HP 43, ATK 3 slam +6 (3d10) or 1 inferno, MV near (fly), S +4, D +3, C +3, I -2, W +1, Ch-2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalWater6HD = Monster{Name: "Elemental, Water (LV 6)", Level: 6, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "215", StatBlock: "AC 15, HP 29, ATK 3 slam +6 (2d6) or 1 whirlpool, MV double near (swim), S +4, D +2, C+2, I -2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var ElementalWater9HD = Monster{Name: "Elemental, Water (LV 9)", Level: 9, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "215", StatBlock: "AC 15, HP 42, ATK 3 slam +6 (3d6) or 1 whirlpool, MV double near (swim), S +4, D +2, C+2, I -2, W +1, Ch -2, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var Elephant = Monster{Name: "Elephant", Level: 7, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "216", StatBlock: "AC 14, HP 34, ATK 2 tusks +6 (1d8), MV near, S +5, D +0, C +3, I-2, W +1, Ch +0, AL N, LV 7", Biomes: []Biome{GrasslandBiome, JungleBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Elf = Monster{Name: "Elf", Level: 2, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "216", StatBlock: "AC 13, HP 9, ATK 1 longbow (far) +3 (1d8) or 1 longsword +1 (1d8), MV near, S +0, D +3, C +0, I +1, W+1, Ch +1, AL L, LV 2", Tags: []MonsterTag{HumanoidTag}}
var Ettercap = Monster{Name: "Ettercap", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "poison", Page: "216", StatBlock: "AC 12, HP 14, ATK 2 bite +2 (1d6) or 1 poison web (near) +2, MV near (climb), S +0, D +2, C +1, I +0, W +0, Ch -1, AL C, LV 3", Biomes: []Biome{CaveBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Fairy = Monster{Name: "Fairy", Level: 1, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "poison", Page: "216", StatBlock: "AC 13, HP 4, ATK 1 needle +3 (1 + poison), MV near (fly), S -2, D +3, C +0, I +1, W +0, Ch +1, AL N, LV 1", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{FeyTag}}
var FrogGiant = Monster{Name: "Frog, Giant", Level: 2, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "216", StatBlock: "AC 12, HP 10, ATK 1 tongue and 1 bite +2 (1d6), MV near (swim), S+2, D +2, C +1, I -3, W +0, Ch -3, AL N, LV 2", Biomes: []Biome{CaveBiome, ForestBiome, JungleBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var Gargoyle = Monster{Name: "Gargoyle", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "216", StatBlock: "AC 16, HP 20, ATK 2 claw +3 (1d6), MV near (fly), S +3, D +1, C +2, I +0, W +1, Ch -1, AL C, LV 4", Tags: []MonsterTag{ElementalTag}}
var GelatinousCube = Monster{Name: "Gelatinous Cube", Level: 5, MonsterAlignment: NeutralAlignment, Move: "", Attack: "toxin", Page: "217", StatBlock: "AC 11, HP 24, ATK 1 touch +4 (1d8 + toxin + engulf), MV near, S +3, D +1, C +2, I -4, W +1, Ch -4, AL N, LV 5", Biomes: []Biome{RuinsBiome, TombBiome}, Tags: []MonsterTag{OozeTag}}
var Ghast = Monster{Name: "Ghast", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "217", StatBlock: "AC 11, HP 20, ATK 2 claw +4 (1d8 + paralyze), MV near, S +3, D +1, C+2, I +0, W +0, Ch +2, AL C, LV 4", Tags: []MonsterTag{UndeadTag}}
var Ghost = Monster{Name: "Ghost", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "life drain", Page: "217", StatBlock: "AC 13, HP 27, ATK 2 death touch +5 (1d8 + life drain) or 1 possess, MV near (fly), S -2, D +3, C +0, I +0, W +0, Ch +4, AL C, LV 6", Tags: []MonsterTag{UndeadTag}}
var Ghoul = Monster{Name: "Ghoul", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "217", StatBlock: "AC 11, HP 11, ATK 1 claw +2 (1d6 + paralyze), MV near, S +2, D +1, C+2, I -3, W -1, Ch +0, AL C, LV 2", Tags: []MonsterTag{UndeadTag}}
var GiantCloud = Monster{Name: "Giant, Cloud", Level: 10, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "218", StatBlock: "AC 15 (leather), HP 48, ATK3 morningstar +9 (2d10), MV double near, S +5, D +4, C +3, I +3, W +3, Ch +3, AL N, LV 10", Biomes: []Biome{ArcticBiome, MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantFire = Monster{Name: "Giant, Fire", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "218", StatBlock: "AC 15 (plate mail), HP 44, ATK 3 greatsword +9 (2d12), MV double near, S +6, D +0, C +4, I +1, W +2, Ch +1, AL C, LV 9", Biomes: []Biome{MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantFrost = Monster{Name: "Giant, Frost", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "219", StatBlock: "AC 14 (chainmail), HP 44, ATK 3 greataxe +8 (2d10), MV double near, S +5, D +1, C +4, I +2, W +3, Ch +2, AL C, LV 9", Biomes: []Biome{ArcticBiome, MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantGoat = Monster{Name: "Giant, Goat", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "219", StatBlock: "AC 12 (leather), HP 39, ATK 2 greatclub +7 (2d8) or 1 boulder (far) +7 (2d10), MV double near (climb), S +4, D +1, C +3, I -2, W+0, Ch -2, AL C, LV 8", Biomes: []Biome{MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantHill = Monster{Name: "Giant, Hill", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "219", StatBlock: "AC 11 (leather), HP 34, ATK 2 greatclub +6 (2d8) or 1 boulder (far) +6 (2d10), MV double near, S+4, D +0, C +3, I -2, W -2, Ch -2, AL C, LV 7", Biomes: []Biome{GrasslandBiome, MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantStone = Monster{Name: "Giant, Stone", Level: 8, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "219", StatBlock: "AC 17, HP 40, ATK 2 greatclub +7 (2d8) or 1 boulder (far) +7 (2d10), MV double near, S +4, D +2, C +4, I +1, W +1, Ch -1, AL N, LV 8", Biomes: []Biome{MountainBiome}, Tags: []MonsterTag{GiantTag}}
var GiantStorm = Monster{Name: "Giant, Storm", Level: 12, MonsterAlignment: LawfulAlignment, Move: "swim", Attack: "electric", Page: "219", StatBlock: "AC 15 (mithral chainmail), HP 58, ATK 3 greatsword +10 (2d12) or 1 lightning bolt, MV double near (swim), S +6, D +2, C +4, I +3, W+4, Ch +4, AL L, LV 12", Biomes: []Biome{MountainBiome, OceanBiome}, Tags: []MonsterTag{GiantTag}}
var GibberingMouther = Monster{Name: "Gibbering Mouther", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "220", StatBlock: "AC 8, HP 21, ATK 2 bite +3 (1d8 + latch), MV near (climb, swim), S+2, D -2, C +3, I -3, W +0, Ch -3, AL N, LV 4", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{AberrationTag}}
var Gladiator = Monster{Name: "Gladiator", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "220", StatBlock: "AC 16 (chainmail + shield), HP15, ATK 2 longsword +3 (1d8) or 1 spear (close/near) +3 (1d6), MV near, S +2, D +1, C +2, I +0, W +0, Ch +1, AL N, LV 3", Tags: []MonsterTag{HumanoidTag}}
var Gnoll = Monster{Name: "Gnoll", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "220", StatBlock: "AC 12 (leather), HP 10, ATK 1 spear (close/near) +1 (1d6) or 1 longbow (far) +1 (1d8), MV near, S+1, D +1, C +1, I -1, W +0, Ch -1, AL C, LV 2", Biomes: []Biome{CaveBiome, GrasslandBiome, SwampBiome}, Tags: []MonsterTag{HumanoidTag}}
var GnomeDeep = Monster{Name: "Gnome, Deep", Level: 3, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "220", StatBlock: "AC 14 (leather + shield), HP 14, ATK 1 pick +3 (1d6) or 1 dart (near) +2 (1d4), MV near, S +2, D +1, C +1, I +1, W +1, Ch +1, AL L, LV 3", Biomes: []Biome{CaveBiome, DeepsBiome, RuinsBiome}, Tags: []MonsterTag{HumanoidTag}}
var Goblin = Monster{Name: "Goblin", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "221", StatBlock: "AC 11, HP 5, ATK 1 club +0 (1d4) or 1 shortbow (far) +1 (1d4), MV near, S +0, D +1, C +1, I -1, W -1, Ch -2, AL C, LV 1", Tags: []MonsterTag{HumanoidTag}}
var GoblinBoss = Monster{Name: "Goblin, Boss", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "221", StatBlock: "AC 14 (chainmail), HP 20, ATK 1 spear (close/near) +3 (1d6), MV near, S +2, D +1, C +2, I -1, W +0, Ch +1, AL C, LV 4", Tags: []MonsterTag{HumanoidTag}}
var GoblinShaman = Monster{Name: "Goblin, Shaman", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "221", StatBlock: "AC 12 (leather), HP 19, ATK 1 staff +0 (1d4) or 1 spell +3, MV near, S+0, D +1, C +1, I +0, W +2, Ch +1, AL C, LV 4", Tags: []MonsterTag{HumanoidTag}}
var GolemClay = Monster{Name: "Golem, Clay", Level: 8, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "222", StatBlock: "AC 14, HP 40, ATK 3 slam +6 (1d8), MV near, S +4, D +0, C +4, I-2, W +0, Ch -2, AL N, LV 8", Tags: []MonsterTag{GolemTag, ConstructTag}}
var GolemFlesh = Monster{Name: "Golem, Flesh", Level: 7, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "222", StatBlock: "AC 9, HP 35, ATK 3 slam +6 (1d8), MV near, S +4, D -1, C +4, I -1, W+1, Ch -3, AL N, LV 7", Tags: []MonsterTag{GolemTag, ConstructTag}}
var GolemIron = Monster{Name: "Golem, Iron", Level: 10, MonsterAlignment: NeutralAlignment, Move: "", Attack: "poison", Page: "222", StatBlock: "AC 19, HP 49, ATK 3 slam +8 (2d8) or 1 poison breath, MV near, S +5, D -1, C +4, I -2, W +0, Ch -2, AL N, LV 10", Tags: []MonsterTag{GolemTag, ConstructTag}}
var GolemStone = Monster{Name: "Golem, Stone", Level: 8, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "222", StatBlock: "AC 18, HP 40, ATK 3 slam +6 (1d10) and 1 slow, MV near, S +4, D -1, C +4, I -2, W +0, Ch -2, AL N, LV 8", Tags: []MonsterTag{GolemTag, ConstructTag}}
var Gorgon = Monster{Name: "Gorgon", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "223", StatBlock: "AC 18, HP 33, ATK 2 gore +6 (2d8) or 1 charge or 1 petrifying breath, MV double near, S +4, D +0, C +2, I -3, W +1, Ch -3, AL C, LV 7", Biomes: []Biome{RuinsBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Gorilla = Monster{Name: "Gorilla", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "223", StatBlock: "AC 12, HP 20, ATK 2 rend +5 (2d6), MV near (climb), S +4, D +2, C +2, I -1, W +1, Ch -1, AL N, LV 4", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{AnimalTag}}
var GrayOoze = Monster{Name: "Gray Ooze", Level: 2, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "223", StatBlock: "AC 11, HP 9, ATK 1 tentacle +2 (1d6), MV near (climb), S +1, D +1, C +0, I -4, W -3, Ch -4, AL N, LV 2", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{OozeTag}}
var Grick = Monster{Name: "Grick", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "223", StatBlock: "AC 14, HP 19, ATK 1 beak +3 (1d8) and 1 tentacle +3 (1d6 + grab), MV near (climb), S +3, D +2, C +1, I -3, W +1, Ch -3, AL N, LV 4", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Griffon = Monster{Name: "Griffon", Level: 4, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "224", StatBlock: "AC 12, HP 19, ATK 2 rend +4 (1d10), MV double near (fly), S +4, D +2, C +1, I -3, W +1, Ch -1, AL N, LV 4", Biomes: []Biome{GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Grimlow = Monster{Name: "Grimlow", Level: 9, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "224", StatBlock: "AC 12, HP 43, ATK 1 grab and 3 bite +6 (2d8), MV near, S +4, D +2, C +3, I -3, W +1, Ch -2, AL N, LV 9", Biomes: []Biome{CaveBiome, ForestBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Guard = Monster{Name: "Guard", Level: 1, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "224", StatBlock: "AC 15 (chainmail + shield), HP 4, ATK 1 spear (close/near) +1 (1d6) or 1 longsword +1 (1d8), MV near, S +1, D +0, C +0, I +0, W +1, Ch +0, AL L, LV 1", Tags: []MonsterTag{HumanoidTag}}
var HagNight = Monster{Name: "Hag, Night", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "225", StatBlock: "AC 14, HP 37, ATK 2 bite +6 (1d10) and 1 blind, MV near, S +4, D +2, C +1, I +2, W +3, Ch +3, AL C, LV 8", Biomes: []Biome{DeepsBiome, ForestBiome, SwampBiome}, Tags: []MonsterTag{FeyTag}}
var HagSea = Monster{Name: "Hag, Sea", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "225", StatBlock: "AC 15, HP 28, ATK 2 claw +4 (1d8), MV near (swim), S +2, D +3, C +1, I+1, W +2, Ch +2, AL C, LV 6", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{FeyTag}}
var HagWeald = Monster{Name: "Hag, Weald", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "225", StatBlock: "AC 14, HP 28, ATK 2 claw +4 (1d8) or 1 drink pain, MV near, S +3, D+2, C +1, I +1, W +2, Ch +3, AL C, LV 6", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{FeyTag}}
var Harpy = Monster{Name: "Harpy", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "226", StatBlock: "AC 13, HP 14, ATK 2 claw +3 (1d6) or 1 song, MV near (fly), S +1, D +3, C +1, I +0, W +0, Ch +1, AL C, LV 3", Biomes: []Biome{ArcticBiome, ForestBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag}}
var HellHound = Monster{Name: "Hell Hound", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "fire", Page: "226", StatBlock: "AC 13, HP 19, ATK 2 bite +4 (1d8) or 1 fire breath, MV double near, S +2, D +1, C +1, I -2, W +1, Ch -3, AL C, LV 4", Tags: []MonsterTag{FiendTag}}
var Hippogriff = Monster{Name: "Hippogriff", Level: 3, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "226", StatBlock: "AC 13, HP 14, ATK 2 rend +3 (1d8), MV double near (fly), S +3, D +3, C+1, I -3, W +1, Ch -2, AL N, LV 3", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hippopotamus = Monster{Name: "Hippopotamus", Level: 5, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "226", StatBlock: "AC 12, HP 24, ATK 2 bite +4 (1d10), MV near (swim), S +4, D +0, C +2, I -3, W +0, Ch -3, AL N, LV 5", Biomes: []Biome{RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag}}
var Hobgoblin = Monster{Name: "Hobgoblin", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "227", StatBlock: "AC 15 (chainmail + shield), HP10, ATK 1 longsword +3 (1d8) or 1 longbow (far) +0 (1d8), MV near, S +3, D +0, C +1, I +2, W +1, Ch +1, AL C, LV 2", Tags: []MonsterTag{HumanoidTag}}
var Horse = Monster{Name: "Horse", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "227", StatBlock: "AC 11, HP 11, ATK 1 hooves +3 (1d6), MV double near, S +3, D +1, C +2, I -3, W +1, Ch -2, AL N, LV 2", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Hydra2HD = Monster{Name: "Hydra, 1 head", Level: 2, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra4HD = Monster{Name: "Hydra, 2 heads", Level: 4, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra6HD = Monster{Name: "Hydra, 3 heads", Level: 6, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra8HD = Monster{Name: "Hydra, 4 heads", Level: 8, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra10HD = Monster{Name: "Hydra, 5 heads", Level: 10, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra12HD = Monster{Name: "Hydra, 6 heads", Level: 12, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra14HD = Monster{Name: "Hydra, 7 heads", Level: 14, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Hydra16HD = Monster{Name: "Hydra, 8 heads", Level: 16, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "227", StatBlock: "AC 15, HP *, ATK 1 bite (near) +6 (1d8), MV near (swim), S +5, D +1, C +2, I -2, W +1, Ch -2, AL N, LV *", Biomes: []Biome{DeepsBiome, OceanBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var InvisibleStalker = Monster{Name: "Invisible Stalker", Level: 6, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "227", StatBlock: "AC 13, HP 29, ATK 3 pummel +4 (1d6), MV near (fly), S +2, D +3, C+2, I +2, W +1, Ch +0, AL N, LV 6", Tags: []MonsterTag{ElementalTag}}
var Jellyfish = Monster{Name: "Jellyfish", Level: 0, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "toxin", Page: "228", StatBlock: "AC 11, HP 1, ATK 1 sting +1 (1 + toxin), MV close (swim), S -4, D+1, C +0, I -4, W +1, Ch -4, AL N, LV 0", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{AnimalTag}}
var Knight = Monster{Name: "Knight", Level: 3, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "228", StatBlock: "AC 17 (plate mail + shield), HP 14, ATK 2 bastard sword +3 (1d8), MV near, S +3, D +0, C +1, I +0, W +0, Ch +1, AL L, LV 3", Tags: []MonsterTag{HumanoidTag}}
var Kobold = Monster{Name: "Kobold", Level: 0, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "228", StatBlock: "AC 13 (leather), HP 1, ATK 1 spear (close/near) +0 (1d6), MV near, S-2, D +2, C +0, I -1, W +0, Ch -1, AL C, LV 0", Tags: []MonsterTag{HumanoidTag}}
var KoboldSorcerer = Monster{Name: "Kobold, Sorcerer", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "228", StatBlock: "AC 13 (leather), HP 13, ATK 1 club +1 (1d4) or 1 spell +2, MV near, S-2, D +2, C +0, I -1, W +1, Ch +2, AL C, LV 3", Tags: []MonsterTag{HumanoidTag}}
var Kraken = Monster{Name: "Kraken", Level: 17, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "electric", Page: "229", StatBlock: "AC 18, HP 80, ATK 4 tentacle (near) +9 (2d12) or 1 storm or 1d4 lightning bolt, MV double near (swim), S +6, D +3, C +4, I +4, W+3, Ch +4, AL C, LV 17", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{MonstrosityTag, GiantTag}}
var LeechGiant = Monster{Name: "Leech, Giant", Level: 2, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "229", StatBlock: "AC 9, HP 10, ATK 1 bite +1 (1d4 + attach), MV near (swim), S +1, D -1, C +1, I -3, W -1, Ch -3, AL N, LV 2", Biomes: []Biome{DeepsBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var Leprechaun = Monster{Name: "Leprechaun", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "spell", Page: "230", StatBlock: "AC 13, HP 19, ATK 1 spell +4, MV near, S +1, D +3, C +1, I +2, W +1, Ch +3, AL N, LV 4", Biomes: []Biome{ForestBiome}, Tags: []MonsterTag{FeyTag}}
var Lich = Monster{Name: "Lich", Level: 13, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "230", StatBlock: "AC 16, HP 62, ATK 2 touch +6 (2d8 + paralysis) and 2 spell +7, MV near, S +3, D +1, C +4, I +4, W+3, Ch +3, AL C, LV 13", Tags: []MonsterTag{UndeadTag}}
var Lion = Monster{Name: "Lion", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "231", StatBlock: "AC 12, HP 15, ATK 2 rend +4 (1d8), MV near, S +4, D +2, C +2, I -3, W+1, Ch -3, AL N, LV 3", Biomes: []Biome{GrasslandBiome, MountainBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Lizardfolk = Monster{Name: "Lizardfolk", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "231", StatBlock: "AC 14 (leather + shield), HP 11, ATK 1 spear (close/near) +2 (1d6), MV near (swim), S +1, D +1, C +2, I-1, W +1, Ch -2, AL C, LV 2", Biomes: []Biome{CaveBiome, DeepsBiome, DesertBiome, JungleBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{HumanoidTag}}
var Mage = Monster{Name: "Mage", Level: 6, MonsterAlignment: LawfulAlignment, Move: "", Attack: "spell", Page: "231", StatBlock: "AC 11, HP 27, ATK 1 spell +5, MV near, S -1, D +1, C +0, I +3, W +1, Ch+0, AL L, LV 6", Tags: []MonsterTag{HumanoidTag}}
var Mammoth = Monster{Name: "Mammoth", Level: 9, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "231", StatBlock: "AC 15, HP 44, ATK 2 tusks +7 (1d12), MV near, S +5, D +0, C +4, I-2, W +1, Ch +0, AL N, LV 9", Biomes: []Biome{ArcticBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var MantaRayGiant = Monster{Name: "Manta Ray, Giant", Level: 8, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "poison", Page: "232", StatBlock: "AC 13, HP 37, ATK 2 sting +5 (1d12 + poison), MV double near (swim), S +3, D +3, C +1, I -2, W +1, Ch -3, AL N, LV 8", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var Manticore = Monster{Name: "Manticore", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "232", StatBlock: "AC 14, HP 29, ATK 2 rend +6 (2d6) or 2 tail spike (far) +4 (1d8), MV double near (fly), S +4, D +2, C +2, I -2, W +1, Ch -2, AL C, LV 6", Biomes: []Biome{CaveBiome, DesertBiome, GrasslandBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Mastiff = Monster{Name: "Mastiff", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "232", StatBlock: "AC 11, HP 4, ATK 1 bite +1 (1d6), MV near, S +1, D +1, C +0, I -2, W+1, Ch -2, AL N, LV 1", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Medusa = Monster{Name: "Medusa", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "poison", Page: "232", StatBlock: "AC 14, HP 38, ATK 1 snake bite +6 (1d6 + poison), MV near, S +2, D+1, C +2, I +2, W +3, Ch +4, AL C, LV 8", Biomes: []Biome{JungleBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Merfolk = Monster{Name: "Merfolk", Level: 2, MonsterAlignment: LawfulAlignment, Move: "swim", Attack: "", Page: "232", StatBlock: "AC 11, HP 9, ATK 1 spear (close/near) +2 (1d6), MV near (swim), S+1, D +1, C +0, I +0, W +1, Ch +1, AL L, LV 2", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{HumanoidTag}}
var Mimic = Monster{Name: "Mimic", Level: 5, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "232", StatBlock: "AC 12, HP 23, ATK 2 bite +5 (1d8 + stick), MV near, S +2, D +0, C +1, I-2, W +0, Ch -3, AL N, LV 5", Biomes: []Biome{CaveBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Minotaur = Monster{Name: "Minotaur", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "233", StatBlock: "AC 14 (chainmail), HP 34, ATK 2 greataxe +6 (1d10) and 1 horns +6 (1d12), MV near, S +4, D +1, C +3, I+1, W +2, Ch +1, AL C, LV 7", Biomes: []Biome{CaveBiome, DeepsBiome, MountainBiome, RuinsBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Moose = Monster{Name: "Moose", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "233", StatBlock: "AC 11, HP 19, ATK 2 antler +3 (1d6), MV double near, S +3, D +0, C +1, I-2, W +0, Ch -2, AL N, LV 4", Biomes: []Biome{ArcticBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var MordanticusTheFlayed = Monster{Name: "Mordanticus The Flayed", Level: 19, MonsterAlignment: NeutralAlignment, Move: "", Attack: "spell", Page: "234", StatBlock: "AC 17, HP 89, ATK 1 rot touch +8 (1d10 + necrosis) and 3 spell +8, MV near, S +4, D +4, C +4, I +5, W +4, Ch +5, AL N, LV 19", Tags: []MonsterTag{LegendaryTag}}
var Mummy = Monster{Name: "Mummy", Level: 10, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "236", StatBlock: "AC 13, HP 47, ATK 3 rot touch +8 (1d10 + necrosis), MV near, S +3, D+0, C +2, I +3, W +2, Ch +3, AL C, LV 10", Tags: []MonsterTag{UndeadTag}}
var Mushroomfolk = Monster{Name: "Mushroomfolk", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "236", StatBlock: "AC 13, HP 15, ATK 2 slam +2 (1d6), MV near, S +2, D -1, C +2, I +0, W+1, Ch +0, AL N, LV 3", Biomes: []Biome{CaveBiome, DeepsBiome, JungleBiome, RuinsBiome, SwampBiome}, Tags: []MonsterTag{HumanoidTag, PlantTag}}
var Naga = Monster{Name: "Naga", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "spell,poison", Page: "237", StatBlock: "AC 16, HP 43, ATK 2 bite +7 (2d6 + poison) and 1 spell +7, MV near (climb), S +4, D +1, C +3, I +2, W+2, Ch +4, AL C, LV 9", Biomes: []Biome{DesertBiome, JungleBiome, SwampBiome}, Tags: []MonsterTag{MonstrosityTag}}
var NagaBone = Monster{Name: "Naga, Bone", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "burrow", Attack: "", Page: "237", StatBlock: "AC 13, HP 31, ATK 2 bite +5 (2d6), MV near (burrow, climb), S +3, D+2, C +4, I -3, W +0, Ch +4, AL C, LV 6", Biomes: []Biome{DesertBiome, JungleBiome, OceanBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Nightmare = Monster{Name: "Nightmare", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "237", StatBlock: "AC 13, HP 29, ATK 2 hooves +5 (1d8), MV double near (fly), S +3, D +3, C +2, I -1, W +1, Ch -2, AL C, LV 6", Tags: []MonsterTag{FiendTag}}
var ObeIxxOfAzarumme = Monster{Name: "Obe-Ixx Of Azarumme", Level: 16, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "blood drain", Page: "238", StatBlock: "AC 18 (+3 plate mail), HP 76, ATK 4 greatsword (near) +11 (1d12 + 2 + Moonbite properties) and 1 bite +9 (1d8 + blood drain) and 1 charm, MV near (climb, fly), S +5, D +3, C +4, I +3, W +4, Ch +5, AL C, LV 16", Tags: []MonsterTag{LegendaryTag}}
var OchreJelly = Monster{Name: "Ochre Jelly", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "239", StatBlock: "AC 9, HP 20, ATK 2 tentacle +3 (1d6), MV near (climb), S +2, D -1, C +2, I -4, W -3, Ch -4, AL N, LV 4", Biomes: []Biome{CaveBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{OozeTag}}
var OctopusGiant = Monster{Name: "Octopus, Giant", Level: 5, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "239", StatBlock: "AC 13, HP 23, ATK 2 tentacle (near) +4 (1d8 + grab), MV near (swim), S +3, D +3, C +1, I -2, W +1, Ch -3, AL N, LV 5", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var Ogre = Monster{Name: "Ogre", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "239", StatBlock: "AC 9, HP 30, ATK 2 greatclub +6 (2d6), MV near, S +4, D -1, C +3, I-2, W -2, Ch -2, AL C, LV 6", Biomes: []Biome{ArcticBiome, ForestBiome, MountainBiome, RuinsBiome, SwampBiome}, Tags: []MonsterTag{GiantTag}}
var Oni = Monster{Name: "Oni", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "239", StatBlock: "AC 11, HP 33, ATK 1 glaive (near) +6 (1d10) or 1 spell +5, MV near, S+5, D +1, C +2, I +2, W +1, Ch +3, AL C, LV 7", Tags: []MonsterTag{GiantTag}}
var Orc = Monster{Name: "Orc", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "240", StatBlock: "AC 15 (chainmail + shield), HP 4, ATK 1 greataxe +2 (1d8), MV near, S +2, D +0, C +0, I -1, W +0, Ch -1, AL C, LV 1", Tags: []MonsterTag{HumanoidTag}}
var OrcChieftain = Monster{Name: "Orc, Chieftain", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "240", StatBlock: "AC 14 (chainmail), HP 19, ATK 2 greataxe +4 (1d10), MV near, S +2, D +1, C +1, I -1, W +0, Ch -1, AL C, LV 4", Tags: []MonsterTag{HumanoidTag}}
var Otyugh = Monster{Name: "Otyugh", Level: 7, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "240", StatBlock: "AC 13, HP 35, ATK 2 tentacle +5 (1d8) and 1 bite +5 (1d10 + disease), MV near, S +4, D -1, C+4, I -2, W +0, Ch -3, AL N, LV 7", Biomes: []Biome{CaveBiome, ForestBiome, SwampBiome}, Tags: []MonsterTag{AberrationTag}}
var Owlbear = Monster{Name: "Owlbear", Level: 6, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "242", StatBlock: "AC 13, HP 30, ATK 2 claw +5 (1d10), MV near (climb), S +4, D +1, C +3, I -2, W +2, Ch -3, AL N, LV 6", Biomes: []Biome{ArcticBiome, ForestBiome, RuinsBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Panther = Monster{Name: "Panther", Level: 3, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "242", StatBlock: "AC 14, HP 14, ATK 2 rend +3 (1d6), MV near (climb), S +3, D +4, C +1, I-2, W +1, Ch -3, AL N, LV 3", Biomes: []Biome{ForestBiome, JungleBiome}, Tags: []MonsterTag{AnimalTag}}
var Peasant = Monster{Name: "Peasant", Level: 1, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "242", StatBlock: "AC 10, HP 4, ATK 1 club +0 (1d4), MV near, S +0, D +0, C +0, I +0, W+0, Ch +0, AL L, LV 1", Tags: []MonsterTag{HumanoidTag}}
var Pegasus = Monster{Name: "Pegasus", Level: 3, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "242", StatBlock: "AC 12, HP 15, ATK 2 hooves +3 (1d6), MV double near (fly), S +3, D +2, C +2, I -3, W +1, Ch +0, AL N, LV 3", Biomes: []Biome{GrasslandBiome, MountainBiome, RiverCoastBiome}, Tags: []MonsterTag{CelestialTag}}
var Phoenix = Monster{Name: "Phoenix", Level: 13, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "", Page: "243", StatBlock: "AC 16, HP 60, ATK 4 rend +8 (2d12), MV double near (fly), S +3, D +4, C +2, I +3, W +3, Ch +3, AL L, LV 13", Tags: []MonsterTag{ElementalTag}}
var PiranhaSwarm = Monster{Name: "Piranha, Swarm", Level: 3, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "243", StatBlock: "AC 12, HP 13, ATK 2 bite +2 (1d6), MV near (swim), S -2, D +2, C +0, I-3, W +0, Ch -3, AL N, LV 3", Biomes: []Biome{RiverCoastBiome}, Tags: []MonsterTag{SwarmTag}}
var Pirate = Monster{Name: "Pirate", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "243", StatBlock: "AC 12 (leather), HP 4, ATK 1 cutlass +1 (1d6) or 1 dagger (close/near) +1 (1d4), MV near, S +1, D +1, C +0, I +0, W +0, Ch +0, AL C, LV 1", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{HumanoidTag}}
var Plesiosaurus = Monster{Name: "Plesiosaurus", Level: 6, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "208", StatBlock: "AC 13, HP 30, ATK 2 bite +5 (2d8), MV double near (swim), S +4, D +3, C +3, I -3, W +1, Ch -3, AL N, LV 6", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{DinosaurTag}}
var Priest = Monster{Name: "Priest", Level: 5, MonsterAlignment: LawfulAlignment, Move: "", Attack: "spell", Page: "243", StatBlock: "AC 15 (chainmail + shield), HP 23, ATK 2 mace +3 (1d6) or 1 spell +3, MV near, S +1, D +0, C +1, I +0, W+2, Ch +1, AL L, LV 5", Tags: []MonsterTag{HumanoidTag}}
var PrimordialSlime = Monster{Name: "Primordial Slime", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "241", StatBlock: "AC 9, HP 30, ATK 2 tentacle +4 (1d10 + dissolve), MV near (climb), S +3, D +2, C +3, I -4, W -3, Ch -4, AL C, LV 6", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{OutsiderTag}}
var Pterodactyl = Monster{Name: "Pterodactyl", Level: 4, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "208", StatBlock: "AC 14, HP 20, ATK 2 beak +4 (1d8 + grab), MV double near (fly), S+2, D +4, C +2, I -2, W +1, Ch -3, AL N, LV 4", Biomes: []Biome{GrasslandBiome, JungleBiome}, Tags: []MonsterTag{DinosaurTag}}
var PurpleWorm = Monster{Name: "Purple Worm", Level: 12, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "poison", Page: "244", StatBlock: "AC 18, HP 57, ATK 2 bite +9 (2d12 + swallow) and 1 sting +9 (1d10 + poison), MV double near (burrow), S +5, D +1, C +3, I -3, W+1, Ch -3, AL N, LV 12", Biomes: []Biome{ArcticBiome, DeepsBiome, DesertBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Rakshasa = Monster{Name: "Rakshasa", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "244", StatBlock: "AC 16, HP 39, ATK 2 claw +6 (1d8), MV near, S +1, D +3, C +3, I +3, W+3, Ch +4, AL C, LV 8", Tags: []MonsterTag{FiendTag}}
var Rat = Monster{Name: "Rat", Level: 0, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "245", StatBlock: "AC 10, HP 1, ATK 1 bite +0 (1 + disease), MV near, S -3, D +0, C +1, I -3, W +1, Ch -3, AL N, LV 0", Tags: []MonsterTag{AnimalTag}}
var RatDire = Monster{Name: "Rat, Dire", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "245", StatBlock: "AC 12, HP 10, ATK 1 bite +2 (1d6 + disease), MV near, S +1, D +2, C +1, I -2, W +1, Ch -2, AL N, LV 2", Tags: []MonsterTag{AnimalTag, DireTag}}
var RatGiant = Monster{Name: "Rat, Giant", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "245", StatBlock: "AC 11, HP 5, ATK 1 bite +1 (1d4 + disease), MV near, S -2, D +1, C +1, I -2, W +1, Ch -2, AL N, LV 1", Tags: []MonsterTag{AnimalTag, GiantTag}}
var RatSwarm = Monster{Name: "Rat, Swarm", Level: 6, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "245", StatBlock: "AC 10, HP 28, ATK 4 bite +0 (1 + disease), MV near, S -3, D +0, C +1, I -3, W +1, Ch -3, AL N, LV 6", Tags: []MonsterTag{AnimalTag, SwarmTag}}
var Rathgamnon = Monster{Name: "Rathgamnon", Level: 19, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "spell", Page: "246", StatBlock: "AC 17, HP 89, ATK 2 rend (near) +9 (2d10) and 2 spell +8, MV double near (fly), S +5, D +3, C +4, I +5, W +6, Ch +5, AL L, LV 19", Tags: []MonsterTag{LegendaryTag}}
var Reaver = Monster{Name: "Reaver", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "247", StatBlock: "AC 17 (plate mail + shield), HP 28, ATK 3 bastard sword +4 (1d8 + 2), MV near, S +3, D +0, C +1, I +0, W+0, Ch +2, AL C, LV 6", Biomes: []Biome{ArcticBiome, MountainBiome}, Tags: []MonsterTag{HumanoidTag}}
var Remorhaz = Monster{Name: "Remorhaz", Level: 10, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "247", StatBlock: "AC 16, HP 47, ATK 3 bite +7 (2d6 + swallow), MV near (burrow), S+5, D +1, C +2, I -3, W +1, Ch -3, AL N, LV 10", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{InsectTag}}
var Rhinoceros = Monster{Name: "Rhinoceros", Level: 5, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "248", StatBlock: "AC 14, HP 25, ATK 2 horn +4 (1d8), MV near, S +4, D -1, C +3, I -3, W+0, Ch -3, AL N, LV 5", Biomes: []Biome{GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var RimeWalker = Monster{Name: "Rime Walker", Level: 9, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "241", StatBlock: "AC 16, HP 43, ATK 4 claw +8 (1d12), MV near (fly), S +4, D +4, C+3, I +2, W +2, Ch +2, AL C, LV 9", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{OutsiderTag}}
var Roc = Monster{Name: "Roc", Level: 15, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "248", StatBlock: "AC 15, HP 69, ATK 4 rend +9 (2d10 + grab), MV double near (fly), S+5, D +3, C +2, I -2, W +2, Ch -2, AL N, LV 15", Biomes: []Biome{ArcticBiome, GrasslandBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag, GiantTag}}
var Roper = Monster{Name: "Roper", Level: 6, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "248", StatBlock: "AC 14, HP 31, ATK 4 tendril (double near) +4 (1d6 + grab) and 1 bite +4 (2d8), MV close (climb), S +3, D -2, C +4, I -1, W +2, Ch +1, AL N, LV 6", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{MonstrosityTag}}
var RotFlower = Monster{Name: "Rot Flower", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "toxin", Page: "249", StatBlock: "AC 9, HP 10, ATK 1 bite +1 (1d4 + toxin), MV none, S +1, D -3, C +1, I-4, W -3, Ch -4, AL N, LV 2", Biomes: []Biome{JungleBiome, TombBiome}, Tags: []MonsterTag{PlantTag}}
var RustMonster = Monster{Name: "Rust Monster", Level: 4, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "249", StatBlock: "AC 13, HP 19, ATK 2 claw +3 (1d6), MV near (climb), S +2, D +3, C +1, I-3, W +1, Ch -3, AL N, LV 4", Biomes: []Biome{CaveBiome, DeepsBiome, DesertBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Sahuagin = Monster{Name: "Sahuagin", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "swim", Attack: "", Page: "249", StatBlock: "AC 14 (leather + shield), HP 9, ATK 2 trident (near) +1 (1d6), MV near (swim), S +1, D +1, C +0, I -1, W +0, Ch -1, AL C, LV 2", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{HumanoidTag}}
var Salamander = Monster{Name: "Salamander", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "fire", Page: "249", StatBlock: "AC 13, HP 24, ATK 2 flaming spear (close/near) +4 (1d6, ignites flammables) or 1 iron longbow (far) +2 (1d8), MV near, S +2, D +0, C +2, I -1, W +1, Ch -1, AL C, LV 5", Biomes: []Biome{DesertBiome, MountainBiome}, Tags: []MonsterTag{ElementalTag}}
var ScarabSwarm = Monster{Name: "Scarab, Swarm", Level: 3, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "249", StatBlock: "AC 13, HP 14, ATK 2 bite +3 (1d6), MV near (fly), S -1, D +3, C +1, I -3, W +0, Ch -3, AL N, LV 3", Biomes: []Biome{DesertBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{InsectTag, SwarmTag}}
var Scarecrow = Monster{Name: "Scarecrow", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "249", StatBlock: "AC 12, HP 15, ATK 2 claws +2 (1d6) or 1 scream, MV near, S +2, D +2, C +2, I +0, W +0, Ch +2, AL C, LV 3", Tags: []MonsterTag{UndeadTag}}
var Scorpion = Monster{Name: "Scorpion", Level: 0, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "250", StatBlock: "AC 11, HP 1, ATK 1 sting +1 (1 + poison), MV near (climb), S -4, D+1, C +0, I -4, W +0, Ch -4, AL N, LV 0", Biomes: []Biome{DesertBiome, JungleBiome}, Tags: []MonsterTag{InsectTag}}
var ScorpionGiant = Monster{Name: "Scorpion, Giant", Level: 3, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "250", StatBlock: "AC 14, HP 13, ATK 1 claw +2 (1d6 + grab) and 1 sting +2 (1d4 + poison), MV near (climb), S +2, D+2, C +0, I -4, W +0, Ch -4, AL N, LV 3", Biomes: []Biome{DeepsBiome, DesertBiome, JungleBiome}, Tags: []MonsterTag{InsectTag, GiantTag}}
var Shadow = Monster{Name: "Shadow", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "250", StatBlock: "AC 12, HP 15, ATK 2 touch +2 (1d4 + drain), MV near (fly), S -4, D +2, C +2, I -2, W +0, Ch -1, AL C, LV 3", Tags: []MonsterTag{UndeadTag}}
var ShamblingMound = Monster{Name: "Shambling Mound", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "251", StatBlock: "AC 14, HP 20, ATK 2 slam +3 (1d6 + engulf), MV near, S +3, D -2, C+2, I -3, W +0, Ch -3, AL N, LV 4", Biomes: []Biome{ForestBiome, SwampBiome}, Tags: []MonsterTag{PlantTag}}
var Shark = Monster{Name: "Shark", Level: 3, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "251", StatBlock: "AC 11, HP 15, ATK 1 bite +3 (1d10), MV near (swim), S +3, D +1, C +2, I-3, W +1, Ch -3, AL N, LV 3", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{AnimalTag}}
var SharkMegalodon = Monster{Name: "Shark, Megalodon", Level: 8, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "251", StatBlock: "AC 13, HP 38, ATK 3 bite +7 (2d8), MV double near (swim), S +5, D +1, C +2, I -3, W +1, Ch -3, AL N, LV 8", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{DinosaurTag}}
var Siren = Monster{Name: "Siren", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "251", StatBlock: "AC 12, HP 18, ATK 2 claw +2 (1d6) or 1 song, MV near (swim, fly), S+0, D +2, C +0, I +2, W +2, Ch +4, AL C, LV 4", Biomes: []Biome{OceanBiome}, Tags: []MonsterTag{FeyTag}}
var Skeleton = Monster{Name: "Skeleton", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "251", StatBlock: "AC 13 (chainmail), HP 11, ATK 1 shortsword +1 (1d6) or 1 shortbow (far) +0 (1d4), MV near, S +1, D +0, C +2, I -2, W +0, Ch -1, AL C, LV 2", Tags: []MonsterTag{UndeadTag}}
var Smilodon = Monster{Name: "Smilodon", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "251", StatBlock: "AC 12, HP 14, ATK 2 bite +3 (1d6), MV near, S +3, D +2, C +1, I -3, W+1, Ch -3, AL N, LV 3", Biomes: []Biome{ArcticBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var SnakeCobra = Monster{Name: "Snake, Cobra", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "poison", Page: "252", StatBlock: "AC 12, HP 4, ATK 1 bite +2 (1 + poison), MV near, S -3, D +2, C +0, I -3, W +0, Ch -3, AL N, LV 1", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{AnimalTag}}
var SnakeGiant = Monster{Name: "Snake, Giant", Level: 5, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "252", StatBlock: "AC 12, HP 23, ATK 2 bite +4 (1d6) and 1 constrict (near), MV near (climb), S +3, D +2, C +1, I -2, W +0, Ch -2, AL N, LV 5", Biomes: []Biome{CaveBiome, DesertBiome, ForestBiome, GrasslandBiome, JungleBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag, GiantTag}}
var SnakeSwarm = Monster{Name: "Snake, Swarm", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "poison", Page: "252", StatBlock: "AC 12, HP 19, ATK 3 bite +2 (1d4 + poison), MV near, S -3, D +2, C +1, I -3, W +0, Ch -3, AL N, LV 4", Biomes: []Biome{JungleBiome, RiverCoastBiome, SwampBiome}, Tags: []MonsterTag{AnimalTag, SwarmTag}}
var Soldier = Monster{Name: "Soldier", Level: 2, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "252", StatBlock: "AC 15 (chainmail + shield), HP10, ATK 1 longsword +2 (1d8) or 1 crossbow (far) +1 (1d6), MV near, S +1, D +0, C +1, I +0, W +0, Ch +0, AL L, LV 2", Tags: []MonsterTag{HumanoidTag}}
var Sphinx = Monster{Name: "Sphinx", Level: 9, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "spell", Page: "253", StatBlock: "AC 16, HP 42, ATK 3 claw +7 (1d10) or 2 spell +5, MV double near (fly), S +4, D +1, C +2, I +4, W+4, Ch +3, AL L, LV 9", Biomes: []Biome{DesertBiome, JungleBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Spider = Monster{Name: "Spider", Level: 0, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "254", StatBlock: "AC 11, HP 1, ATK 2 bite +1 (1 + poison), MV near (climb), S -4, D+1, C +0, I -4, W +0, Ch -4, AL N, LV 0", Tags: []MonsterTag{InsectTag}}
var SpiderGiant = Monster{Name: "Spider, Giant", Level: 3, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "254", StatBlock: "AC 13, HP 13, ATK 1 bite +3 (1d4 + poison), MV near (climb), S +2, D +3, C +0, I -2, W +1, Ch -2, AL N, LV 3", Biomes: []Biome{ArcticBiome, CaveBiome, DeepsBiome, DesertBiome, ForestBiome, JungleBiome, RuinsBiome}, Tags: []MonsterTag{InsectTag, GiantTag}}
var SpiderSwarm = Monster{Name: "Spider, Swarm", Level: 2, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "poison", Page: "254", StatBlock: "AC 13, HP 9, ATK 1 bite +3 (1d4 + poison), MV near (climb), S -1, D+3, C +0, I -3, W +1, Ch -3, AL N, LV 2", Biomes: []Biome{CaveBiome, ForestBiome, JungleBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{InsectTag, SwarmTag}}
var Stingbat = Monster{Name: "Stingbat", Level: 1, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "blood drain", Page: "254", StatBlock: "AC 12, HP 4, ATK 1 beak +2 (1d4 + blood drain), MV near (fly), S -2, D +2, C +0, I -2, W +0, Ch -2, AL N, LV 1", Biomes: []Biome{CaveBiome, ForestBiome, JungleBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Strangler = Monster{Name: "Strangler", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "254", StatBlock: "AC 12, HP 14, ATK 2 claws +2 (1d6), MV near (climb), S -2, D +2, C +1, I -2, W +0, Ch -2, AL C, LV 3", Biomes: []Biome{DeepsBiome, RuinsBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Tarrasque = Monster{Name: "The Tarrasque", Level: 30, MonsterAlignment: NeutralAlignment, Move: "burrow", Attack: "", Page: "256", StatBlock: "AC 22, HP 140, ATK 4 thrash (near) +13 (3d10 + sever) and 1 bite (near) +13 (5d10 + sever + swallow), MV triple near (burrow, swim), S +7, D +2, C +5, I -3, W +1, Ch -3, AL N, LV 30", Tags: []MonsterTag{LegendaryTag}}
var TenEyedOracle = Monster{Name: "The Ten-Eyed Oracle", Level: 18, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "", Page: "255", StatBlock: "AC 17, HP 85, ATK 2d4 eyestalk ray, MV near (fly), S +4, D +5, C +4, I +5, W +4, Ch +4, AL C, LV 18", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{LegendaryTag}}
var WanderingMerchant = Monster{Name: "The Wandering Merchant", Level: 15, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "258", StatBlock: "AC 16 (mithral chainmail), HP 71, ATK 4 +3 vorpal bastard sword +9 (1d10 + lop), MV near, S +3, D +3, C +4, I +3, W +4, Ch +5, AL L, LV 15", Tags: []MonsterTag{LegendaryTag}}
var Thief = Monster{Name: "Thief", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "259", StatBlock: "AC 13 (leather), HP 13, ATK 1 dagger (close/near) +2 (1d4) or 1 shortsword +0 (1d6), MV near, S+0, D +2, C +0, I +0, W +0, Ch +1, AL N, LV 3", Tags: []MonsterTag{HumanoidTag}}
var Thug = Monster{Name: "Thug", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "259", StatBlock: "AC 13 (leather + shield), HP 4, ATK 1 shortsword +1 (1d6), MV near, S +1, D +0, C +0, I -1, W +1, Ch -1, AL C, LV 1", Tags: []MonsterTag{HumanoidTag}}
var Treant = Monster{Name: "Treant", Level: 8, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "259", StatBlock: "AC 14, HP 38, ATK 3 slam +8 (1d10) or 1 rock (far) +8 (2d12), MV near, S +4, D -1, C +2, I +2, W +3, Ch +1, AL N, LV 8", Biomes: []Biome{ArcticBiome, ForestBiome, JungleBiome, MountainBiome}, Tags: []MonsterTag{PlantTag}}
var Triceratops = Monster{Name: "Triceratops", Level: 7, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "208", StatBlock: "AC 17, HP 35, ATK 2 horns +6 (1d10) or 1 charge, MV near, S +4, D -1, C +4, I -3, W +1, Ch -3, AL N, LV 7", Biomes: []Biome{GrasslandBiome, JungleBiome}, Tags: []MonsterTag{DinosaurTag}}
var Troll = Monster{Name: "Troll", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "259", StatBlock: "AC 12, HP 24, ATK 2 claw +4 (1d6) and 1 bite +4 (1d10), MV near, S+3, D +2, C +2, I -1, W +0, Ch -1, AL C, LV 5", Biomes: []Biome{CaveBiome, DeepsBiome, ForestBiome, JungleBiome, MountainBiome, RiverCoastBiome}, Tags: []MonsterTag{GiantTag}}
var TrollFrost = Monster{Name: "Troll, Frost", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "260", StatBlock: "AC 13, HP 34, ATK 2 claw +5 (1d8) and 1 bite +5 (1d12), MV near, S +3, D +2, C +3, I -1, W +0, Ch -1, AL C, LV 7", Biomes: []Biome{ArcticBiome}, Tags: []MonsterTag{GiantTag}}
var Tyrannosaurus = Monster{Name: "Tyrannosaurus", Level: 9, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "208", StatBlock: "AC 13, HP 44, ATK 3 bite +8 (2d12), MV double near, S +5, D +1, C +4, I -3, W +1, Ch -3, AL N, LV 9", Biomes: []Biome{GrasslandBiome, JungleBiome, SwampBiome}, Tags: []MonsterTag{DinosaurTag}}
var Unicorn = Monster{Name: "Unicorn", Level: 4, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "260", StatBlock: "AC 12, HP 20, ATK 1 hooves +3 (1d6), MV double near, S +3, D +2, C +2, I +1, W +2, Ch +3, AL L, LV 4", Biomes: []Biome{ForestBiome, GrasslandBiome}, Tags: []MonsterTag{CelestialTag}}
var Vampire = Monster{Name: "Vampire", Level: 11, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "blood drain", Page: "260", StatBlock: "AC 15, HP 52, ATK 3 bite +7 (1d8 + blood drain) or 1 charm, MV near (climb), S +4, D +3, C +3, I +1, W+3, Ch +4, AL C, LV 11", Tags: []MonsterTag{UndeadTag}}
var VampireSpawn = Monster{Name: "Vampire Spawn", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "blood drain", Page: "261", StatBlock: "AC 13 (leather), HP 25, ATK 2 bite +4 (1d8 + blood drain), MV near (climb), S +3, D +2, C +3, I -1, W +1, Ch +2, AL C, LV 5", Tags: []MonsterTag{UndeadTag}}
var Velociraptor = Monster{Name: "Velociraptor", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "208", StatBlock: "AC 13, HP 10, ATK 1 claw +3 (1d6), MV double near, S -1, D +3, C +1, I-2, W +1, Ch -3, AL N, LV 2", Biomes: []Biome{GrasslandBiome, JungleBiome}, Tags: []MonsterTag{DinosaurTag}}
var VioletFungus = Monster{Name: "Violet Fungus", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "261", StatBlock: "AC 7, HP 9, ATK 2 tendril (near) +0 (1d4), MV close, S -3, D -2, C +0, I -4, W -3, Ch -4, AL N, LV 2", Biomes: []Biome{CaveBiome, ForestBiome, RuinsBiome, TombBiome}, Tags: []MonsterTag{PlantTag}}
var Viperian = Monster{Name: "Viperian", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "262", StatBlock: "AC 13, HP 13, ATK 2 scimitar +2 (1d6) or 1 javelin (close/far) +2 (1d4), MV near, S +1, D +1, C +0, I+0, W +1, Ch +0, AL C, LV 3", Biomes: []Biome{JungleBiome}, Tags: []MonsterTag{MonstrosityTag}}
var ViperianOphid = Monster{Name: "Viperian Ophid", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "262", StatBlock: "AC 14, HP 28, ATK 3 falchion +5 (1d10) or 2 longbow (far) +3 (1d8), MV near (climb), S +4, D +2, C +1, I+1, W +1, Ch +1, AL C, LV 6", Biomes: []Biome{JungleBiome}, Tags: []MonsterTag{MonstrosityTag}}
var ViperianWizard = Monster{Name: "Viperian Wizard", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "spell", Page: "262", StatBlock: "AC 13, HP 37, ATK 1 dagger (close/near) +2 (1d4) or 2 spell +5, MV near, S +0, D +1, C +0, I +3, W +1, Ch +1, AL C, LV 8", Biomes: []Biome{JungleBiome}, Tags: []MonsterTag{MonstrosityTag}}
var VoidSpawn = Monster{Name: "Void Spawn", Level: 7, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "toxin", Page: "241", StatBlock: "AC 13, HP 34, ATK 2 scythe +6 (1d10) and 1 tentacles (1d12 + toxin), MV near (fly), S +4, D +1, C+3, I +0, W +1, Ch -1, AL C, LV 7", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{OutsiderTag}}
var VoidSpider = Monster{Name: "Void Spider", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "poison", Page: "241", StatBlock: "AC 13, HP 23, ATK 2 bite +4 (1d8 + poison), MV near (climb), S +3, D +3, C +1, I -1, W +1, Ch -2, AL C, LV 5", Biomes: []Biome{DeepsBiome}, Tags: []MonsterTag{OutsiderTag}}
var Vulture = Monster{Name: "Vulture", Level: 1, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "263", StatBlock: "AC 10, HP 5, ATK 1 tear +1 (1d4), MV near (fly), S +1, D +0, C +1, I -3, W +1, Ch -3, AL N, LV 1", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{AnimalTag}}
var WaspGiant = Monster{Name: "Wasp, Giant", Level: 2, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "venom", Page: "263", StatBlock: "AC 13, HP 9, ATK 1 sting +3 (1d6 + venom), MV near (fly), S +1, D +3, C +0, I -3, W +0, Ch -3, AL N, LV 2", Biomes: []Biome{ForestBiome, RuinsBiome}, Tags: []MonsterTag{InsectTag, GiantTag}}
var Wererat = Monster{Name: "Wererat", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "263", StatBlock: "AC 13 (leather), HP 14, ATK 2 bite +2 (1d6), MV near (climb), S +1, D+2, C +1, I -1, W +1, Ch -1, AL C, LV 3", Biomes: []Biome{TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Werewolf = Monster{Name: "Werewolf", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "263", StatBlock: "AC 12, HP 20, ATK 2 rend +3 (1d6), MV double near, S +3, D +2, C +2, I +0, W +1, Ch +0, AL C, LV 4", Biomes: []Biome{ArcticBiome, CaveBiome, DesertBiome, ForestBiome, GrasslandBiome, JungleBiome, MountainBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Wight = Monster{Name: "Wight", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "life drain", Page: "263", StatBlock: "AC 14 (chainmail), HP 15, ATK 1 bastard sword +3 (1d10) and 1 life drain +3, MV near, S +3, D +1, C +2, I +1, W +0, Ch +3, AL C, LV 3", Tags: []MonsterTag{UndeadTag}}
var WillOWisp = Monster{Name: "Will-O-Wisp", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "life drain", Page: "264", StatBlock: "AC 13, HP 10, ATK 1 life drain +3, MV near (fly), S -3, D +3, C +1, I -1, W -1, Ch -2, AL C, LV 2", Tags: []MonsterTag{UndeadTag}}
var Wolf = Monster{Name: "Wolf", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "264", StatBlock: "AC 12, HP 10, ATK 1 bite +2 (1d6), MV double near, S +2, D +2, C +1, I-2, W +1, Ch +0, AL N, LV 2", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag}}
var WolfDire = Monster{Name: "Wolf, Dire", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "264", StatBlock: "AC 12, HP 19, ATK 2 bite +4 (1d8), MV double near, S +3, D +2, C +1, I-1, W +1, Ch +0, AL N, LV 4", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome}, Tags: []MonsterTag{AnimalTag, DireTag}}
var WolfWinter = Monster{Name: "Wolf, Winter", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "ice", Page: "264", StatBlock: "AC 12, HP 23, ATK 2 bite +4 (1d6) or 1 frost breath, MV double near, S +3, D +2, C +1, I +0, W +1, Ch +0, AL C, LV 5", Biomes: []Biome{ArcticBiome, MountainBiome}, Tags: []MonsterTag{AnimalTag}}
var Worg = Monster{Name: "Worg", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "265", StatBlock: "AC 11, HP 14, ATK 1 bite +3 (1d6), MV double near, S +2, D +1, C +1, I-2, W +1, Ch -2, AL C, LV 3", Biomes: []Biome{ForestBiome, GrasslandBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Wraith = Monster{Name: "Wraith", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "life drain", Page: "265", StatBlock: "AC 14, HP 36, ATK 3 death touch +6 (1d10 + life drain), MV near (fly), S -4, D +4, C +0, I +0, W +0, Ch +3, AL C, LV 8", Tags: []MonsterTag{UndeadTag}}
var Wyvern = Monster{Name: "Wyvern", Level: 8, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "poison", Page: "265", StatBlock: "AC 15, HP 37, ATK 2 rend +6 (1d8) and 1 stinger +6 (1d6 + poison), MV double near (fly), S +4, D +2, C +1, I -3, W +1, Ch -3, AL N, LV 8", Biomes: []Biome{ArcticBiome, CaveBiome, GrasslandBiome, MountainBiome}, Tags: []MonsterTag{DragonTag}}
var Zombie = Monster{Name: "Zombie", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "265", StatBlock: "AC 8, HP 11, ATK 1 slam +2 (1d6), MV near, S +2, D -2, C +2, I -2, W -2, Ch -3, AL C, LV 2", Tags: []MonsterTag{UndeadTag}}

// Cursed Scrolls 1-3
var Bittermold = Monster{Name: "Bittermold", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS1-46", StatBlock: "AC 12, HP 5, ATK 1 shortsword +1 (1d6) or 1 sling (far) +2 (1d4), MV near, S +1, D +2, C +1, I +0, W +0, Ch +0, AL C, LV 1", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Bogthorn = Monster{Name: "Bogthorn", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "poison", Page: "CS1-46", StatBlock: "AC 13, HP 11, ATK 1 stab +0 (1d4) or 1 thorn hail (near) +2 (1d4 + poison), MV near (climb), S +0, D +2, C +2, I -3, W +1, Ch -3, AL C, LV 2", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Dralech = Monster{Name: "Dralech", Level: 6, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS1-46", StatBlock: "AC 13, HP 29, ATK 3 bone axe +5 (2d6), MV near, S +4, D +1, C +2, I +0, W +0, Ch +1, AL C, LV 6", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var GordockBreeg = Monster{Name: "Gordock Breeg", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS1-46", StatBlock: "AC 15 (chainmail), HP 19, ATK 1 bastard sword +3 (1d10) or 1 sling (far) +2 (1d4), MV near, S +3, D +2, C +1, I +1, W +2, Ch +2, AL N, LV 4", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Hexling = Monster{Name: "Hexling", Level: 2, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS1-46", StatBlock: "AC 12, HP 11, ATK 1 chill touch +2 (1d6 + energy drain), MV near, S +0, D +2, C +2, I +0, W +1, Ch +0, AL C, LV 2", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{FiendTag}}
var Howler = Monster{Name: "Howler", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS1-46", StatBlock: "AC 15 (leather + shield), HP 5, ATK 1 club +2 (1d4) or 1 sling (far) +2 (1d4), MV near, S +2, D +2, C +1, I +0, W +0, Ch +0, AL N, LV 1", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var IchorOoze = Monster{Name: "Ichor Ooze", Level: 3, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "CS1-47", StatBlock: "AC 12, HP 15, ATK 2 tendril +3 (1d6), MV near (climb), S +3, D +2, C +2, I -3, W +1, Ch -3, AL N, LV 3", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{OozeTag}}
var MarrowFiend = Monster{Name: "Marrow Fiend", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "climb", Attack: "", Page: "CS1-47", StatBlock: "AC 15, HP 39, ATK 2 claws +7 (1d10) and 1 sap gout (near line) +5 (2d6 + sap), MV near (climb), S +4, D +4, C +3, I +2, W +3, Ch +3, AL C, LV 8", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Mugdulblub = Monster{Name: "Mugdulblub", Level: 12, MonsterAlignment: ChaoticAlignment, Move: "climb,swim", Attack: "", Page: "CS1-47", StatBlock: "AC 16, HP 58, ATK 3 tendril (near) +8 (2d8) and 1 dissolve, MV near (climb, swim), S +5, D +3, C +4, I +2, W +3, Ch +4, AL C, LV 12", Tags: []MonsterTag{LegendaryTag}}
var MutantCatfish = Monster{Name: "Mutant Catfish", Level: 2, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "poison", Page: "CS1-47", StatBlock: "AC 11, HP 11, ATK 1 claw +2 (1d6) or 1 barb (near) +2 (1d4 + poison), MV near (swim), S +2, D +1, C +2, I -2, W +0, Ch -2, AL N, LV 2", Biomes: []Biome{RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var PlogrinaBittermold = Monster{Name: "Plogrina Bittermold", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS1-48", StatBlock: "AC 14, HP 25, ATK 2 tendril (near) +4 (1d6), MV near, S +0, D +4, C +3, I +1, W +1, Ch +3, AL C, LV 5", Biomes: []Biome{ForestBiome, GrasslandBiome, RiverCoastBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{HumanoidTag}}
var Skrell = Monster{Name: "Skrell", Level: 1, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS1-48", StatBlock: "AC 13, HP 5, ATK 1 claw +2 (1d6), MV double near, S +2, D +3, C +1, I -3, W +1, Ch -3, AL C, LV 1", Biomes: []Biome{ForestBiome, GrasslandBiome, JungleBiome, SwampBiome}, Tags: []MonsterTag{DinosaurTag}}
var TarBat = Monster{Name: "Tar Bat", Level: 1, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "", Page: "CS1-48", StatBlock: "AC 13, HP 4, ATK 1 bite +3 (1d4), MV near (fly), S -3, D +3, C +0, I -3, W +1, Ch -3, AL N, LV 1", Biomes: []Biome{CaveBiome, RuinsBiome, SwampBiome, TombBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Willowman = Monster{Name: "The Willowman", Level: 13, MonsterAlignment: ChaoticAlignment, Move: "teleport", Attack: "", Page: "CS1-48", StatBlock: "AC 17, HP 61, ATK 3 finger needle +9 (2d10) and 1 terrify, MV near (teleport), S +5, D +7, C +3, I +4, W +4, Ch +5, AL C, LV 13", Tags: []MonsterTag{LegendaryTag}}
var CamelSilver = Monster{Name: "Camel, Silver", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-40", StatBlock: "AC 10, HP 13, ATK 1 hoof +3 (1d6) or 1 spit (near) +0 (1d4), MV double near, S +3, D +0, C +4, I -2, W +1, Ch +0, AL N, LV 2", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{AnimalTag}}
var CanyonApe = Monster{Name: "Canyon Ape", Level: 6, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "CS2-40", StatBlock: "AC 13, HP 29, ATK 3 rend +6 (2d6), MV near (climb), S +4, D +3, C +2, I -1, W +1, Ch -1, AL N, LV 6", Biomes: []Biome{DesertBiome, MountainBiome}, Tags: []MonsterTag{AnimalTag}}
var Donkey = Monster{Name: "Donkey", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-40", StatBlock: "AC 10, HP 6, ATK 1 hoof +3 (1d4), MV near, S +3, D +0, C +2, I -2, W +1, Ch -2, AL N, LV 1", Biomes: []Biome{DesertBiome, ForestBiome, GrasslandBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Dunefiend = Monster{Name: "Dunefiend", Level: 4, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS2-40", StatBlock: "AC 14, HP 20, ATK 2 tear +4 (1d8) or 1 howl, MV near, S +3, D +4, C +2, I +0, W +1, Ch +0, AL C, LV 4", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{FiendTag}}
var DustDevil = Monster{Name: "Dust Devil", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS2-40", StatBlock: "AC 16, HP 36, ATK 2 lacerate +5 (2d8), MV double near, S +4, D +4, C +0, I -2, W +0, Ch -2, AL C, LV 8", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{ElementalTag}}
var HeroGladiator = Monster{Name: "Hero, Gladiator", Level: 5, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-41", StatBlock: "AC 16 (chainmail + shield), HP 25, ATK 3 bastard sword +5 (1d8) or 1 spear (close/near) +5 (1d6), MV near, S +3, D +1, C +3, I +0, W +1, Ch +1, AL N, LV 5", Tags: []MonsterTag{HumanoidTag}}
var HorseWar = Monster{Name: "Horse, War", Level: 3, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-41", StatBlock: "AC 11, HP 15, ATK 1 hooves +4 (1d6), MV double near, S +3, D +1, C +2, I -3, W +1, Ch -1, AL N, LV 3", Tags: []MonsterTag{AnimalTag}}
var Mirage = Monster{Name: "Mirage", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS2-41", StatBlock: "AC 6, HP 32, ATK 1 leech, MV near, S -4, D -4, C -4, I +2, W +1, Ch +4, AL C, LV 8", Tags: []MonsterTag{FiendTag}}
var RasGodai = Monster{Name: "Ras-Godai", Level: 3, MonsterAlignment: ChaoticAlignment, Move: "teleport", Attack: "", Page: "CS2-41", StatBlock: "AC 13 (leather), HP 13, ATK 1 razor chain (near) +4 (1d6), MV near (teleport), S +1, D +2, C +0, I +0, W +1, Ch +0, AL C, LV 3", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{HumanoidTag}}
var RookiePitFighter = Monster{Name: "Rookie, Pit-Fighter", Level: 1, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-41", StatBlock: "AC 14 (leather + shield), HP 5, ATK 1 shortsword +1 (1d6) or 1 javelin (close/far) +1 (1d4), MV near, S +1, D +1, C +1, I +0, W +0, Ch +0, AL N, LV 1", Tags: []MonsterTag{HumanoidTag}}
var Scrag = Monster{Name: "Scrag", Level: 2, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "CS2-43", StatBlock: "AC 12, HP 11, ATK 1 claw +2 (1d6), MV near (climb), S +2, D +2, C +2, I -2, W +1, Ch -3, AL N, LV 2", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{MonstrosityTag}}
var ScragWar = Monster{Name: "Scrag, War", Level: 3, MonsterAlignment: NeutralAlignment, Move: "climb", Attack: "", Page: "CS2-43", StatBlock: "AC 12, HP 15, ATK 1 claw +3 (1d6), MV near (climb), S +3, D +2, C +2, I -2, W +1, Ch -3, AL N, LV 3", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Siruul = Monster{Name: "Siruul", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS2-43", StatBlock: "AC 14 (leather), HP 9, ATK 1 scimitar +3 (1d6) or 1 longbow (far) +3 (1d8), MV double near (mount), S +0, D +3, C +0, I +1, W +1, Ch +1, AL N, LV 2", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{HumanoidTag}}
var Scourge = Monster{Name: "The Scourge", Level: 18, MonsterAlignment: ChaoticAlignment, Move: "fly", Attack: "electric", Page: "CS2-43", StatBlock: "AC 17, HP 84, ATK 3 rend +10 (2d10) or lightning breath, MV double near (fly), S +6, D +3, C +3, I +4, W +4, Ch +5, AL C, LV 18", Biomes: []Biome{DesertBiome}, Tags: []MonsterTag{LegendaryTag}}
var DrakeGreater = Monster{Name: "Drake, Greater", Level: 8, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "fire", Page: "CS3-44", StatBlock: "AC 15, HP 38, ATK 2 claw +7 (1d12) or 1 fire gout, MV double near (fly), S +4, D +3, C +2, I -2, W +2, Ch +0, AL N, LV 8", Biomes: []Biome{ArcticBiome, CaveBiome, MountainBiome}, Tags: []MonsterTag{DragonTag}}
var DrakeLesser = Monster{Name: "Drake, Lesser", Level: 6, MonsterAlignment: NeutralAlignment, Move: "fly", Attack: "fire", Page: "CS3-44", StatBlock: "AC 13, HP 28, ATK 2 claw +5 (1d10) or 1 fire spit, MV double near (fly), S +3, D +2, C +1, I -2, W +1, Ch +0, AL N, LV 6", Biomes: []Biome{ArcticBiome, CaveBiome, MountainBiome}, Tags: []MonsterTag{DragonTag}}
var Draugr = Monster{Name: "Draugr", Level: 5, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS3-45", StatBlock: "AC 16 (chainmail + shield), HP 25, ATK 2 bastard sword +4 (1d8), MV near, S +3, D +1, C +3, I +0, W +0, Ch +3, AL C, LV 5", Tags: []MonsterTag{UndeadTag}}
var Dverg = Monster{Name: "Dverg", Level: 3, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "CS3-45", StatBlock: "AC 13 (chainmail), HP 14, ATK 1 greataxe +2 (1d10), MV near, S +2, D +0, C +1, I +1, W +1, Ch +0, AL L, LV 3", Biomes: []Biome{ArcticBiome, CaveBiome, DeepsBiome, MountainBiome}, Tags: []MonsterTag{HumanoidTag}}
var Nord = Monster{Name: "Nord", Level: 2, MonsterAlignment: NeutralAlignment, Move: "", Attack: "", Page: "CS3-45", StatBlock: "AC 15 (chainmail + shield), HP 10, ATK 1 greataxe +2 (1d8) or 1 shield wall, MV near, S +2, D +0, C +1, I +0, W +0, Ch +0, AL N, LV 2", Biomes: []Biome{ArcticBiome, OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{HumanoidTag}}
var TrollDeep = Monster{Name: "Troll, Deep", Level: 8, MonsterAlignment: ChaoticAlignment, Move: "", Attack: "", Page: "CS3-46", StatBlock: "AC 13, HP 39, ATK 2 claw +7 (1d10), MV near, S +4, D +2, C +3, I -1, W +1, Ch +1, AL C, LV 8", Biomes: []Biome{ArcticBiome, CaveBiome, DeepsBiome, MountainBiome}, Tags: []MonsterTag{MonstrosityTag}}
var SeaSerpent = Monster{Name: "Sea Serpent", Level: 12, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "CS3-46", StatBlock: "AC 14, HP 58, ATK 3 bite +8 (2d12), MV double near (swim), S +5, D +0, C +4, I -3, W +1, Ch -3, AL N, LV 12", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{MonstrosityTag}}
var SeaNymph = Monster{Name: "Sea Nymph", Level: 3, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "CS3-46", StatBlock: "AC 13, HP 13, ATK 1 slash +2 (1d6) or 1 sing, MV near (swim), S +0, D +3, C +0, I +1, W +1, Ch +2, AL N, LV 3", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{FeyTag}}
var Orca = Monster{Name: "Orca", Level: 8, MonsterAlignment: NeutralAlignment, Move: "swim", Attack: "", Page: "CS3-46", StatBlock: "AC 12, HP 39, ATK 2 bite +5 (1d10), MV double near (swim), S +4, D +0, C +3, I -2, W +1, Ch -2, AL N, LV 8", Biomes: []Biome{OceanBiome, RiverCoastBiome}, Tags: []MonsterTag{AnimalTag}}
var Oracle = Monster{Name: "Oracle", Level: 4, MonsterAlignment: NeutralAlignment, Move: "", Attack: "spell", Page: "CS3-46", StatBlock: "AC 10, HP 19, ATK 1 stave +2 (1d4) or 1 spell +3, MV near, S +1, D +0, C +1, I +0, W +2, Ch +1, AL N, LV 4", Tags: []MonsterTag{HumanoidTag}}
var Werebear = Monster{Name: "Werebear", Level: 7, MonsterAlignment: LawfulAlignment, Move: "", Attack: "", Page: "CS3-47", StatBlock: "AC 11, HP 34, ATK 2 claw +6 (1d8), MV near, S +4, D +1, C +3, I +0, W +2, Ch -1, AL L, LV 7", Biomes: []Biome{ArcticBiome, ForestBiome, MountainBiome, RiverCoastBiome}, Tags: []MonsterTag{MonstrosityTag}}
var Valkyrie = Monster{Name: "Valkyrie", Level: 14, MonsterAlignment: LawfulAlignment, Move: "fly", Attack: "", Page: "CS3-47", StatBlock: "AC 17 (plate mail + shield), HP66, ATK 3 blessed spear (near) +9 (3d6), MV double near (fly), S +4, D +3, C +3, I +3, W +4, Ch +5, AL L, LV 14", Tags: []MonsterTag{CelestialTag}}
