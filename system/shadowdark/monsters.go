package shadowdark

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

type Biome string

func (b Biome) String() string {
	return string(b)
}

const (
	ArcticBiome    Biome = "arctic"
	CaveBiome      Biome = "cave"
	DeepsBiome     Biome = "deeps"
	DesertBiome    Biome = "desert"
	ForestBiome    Biome = "forest"
	GrasslandBiome Biome = "grassland"
	JungleTag      Biome = "jungle"
	MountainBiome  Biome = "mountain"
	OceanBiome     Biome = "ocean"
	RiverCoastTag  Biome = "river/coast"
	RuinsBiome     Biome = "ruins"
	SwampBiome     Biome = "swamp"
	TombBiome      Biome = "tomb"
)

type Alignment string

func (a Alignment) String() string {
	return string(a)
}

const (
	LawfulAlignment  Alignment = "l"
	NeutralAlignment Alignment = "n"
	ChaoticAlignment Alignment = "c"
)

type Monster struct {
	Name      string       `json:"name"`
	Tags      []MonsterTag `json:"tags"`
	Level     int          `json:"level"`
	Biomes    []Biome      `json:"biomes"`
	Alignment Alignment    `json:"alignment"`
	Move      string       `json:"move"`
	Attack    string       `json:"attack"`
	Page      string       `json:"page"`
	StatBlock string       `json:"stat_block"`
}
