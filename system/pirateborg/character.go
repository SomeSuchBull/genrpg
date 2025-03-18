package pirateborg

import (
	"fmt"
	"slices"

	"github.com/genrpg/utils"
	"github.com/ttacon/chalk"
)

type Stats struct {
	Strength  int `json:"strength"`
	Agility   int `json:"agility"`
	Presence  int `json:"presence"`
	Toughness int `json:"toughness"`
	Spirit    int `json:"spirit"`
}

func (s Stats) String() string {
	return fmt.Sprintf("  Strength: %d\n  Agility: %d\n  Presence: %d\n  Toughness: %d\n  Spirit: %d",
		s.Strength, s.Agility, s.Presence, s.Toughness, s.Spirit)
}

// type Stats struct {
// 	Strength  int `json:"STRENGTH"`
// 	Agility   int `json:"AGILITY"`
// 	Presence  int `json:"PRESENCE"`
// 	Toughness int `json:"TOUGHNESS"`
// 	Spirit    int `json:"SPIRIT"`
// }

type PlayerCharacter struct {
	Character         Character   `json:"character"`
	Clothing          Clothing    `json:"clothing"`
	Hat               Hat         `json:"hat"`
	Stats             Stats       `json:"stats"`
	Features          []Feature   `json:"features"`
	Weapons           Weapons     `json:"weapons"`
	Gear              Gear        `json:"gear"`
	Class             PlayerClass `json:"class"`
	Name              string      `json:"name"`
	Nickname          string      `json:"nickname"`
	Armor             string      `json:"armor"`
	Container         string      `json:"container"`
	Money             string      `json:"money"`
	ThingOfImportance string      `json:"thingOfImportance"`
	HP                int         `json:"hp"`
	CarryingCapacity  int         `json:"carryingCapacity"`
}

type Weapons []Weapon

type Gear []Item

func (g Gear) String() string {
	var val string
	for _, i := range g {
		val += fmt.Sprintln(i.String())
	}
	return val
}

func (ws Weapons) String() string {
	var val string
	for _, w := range ws {
		val += fmt.Sprintln(w.String())
	}
	return val
}

type Character struct {
	Background          string `json:"background,omitempty"`
	DistinctiveFlaw     string `json:"distinctiveFlaw,omitempty"`
	PhysicalTrademark   string `json:"physicalTrademark,omitempty"`
	Idiosyncrasy        string `json:"idiosyncrasy,omitempty"`
	UnfortunateIncident string `json:"unfortunateIncident,omitempty"`
}

func (c Character) String() string {
	val := fmt.Sprintf("%s %s\n", chalk.Yellow.Color("Background:"), c.Background)
	val += fmt.Sprintf("%s %s\n", chalk.Yellow.Color("Distinctive Flaw:"), c.DistinctiveFlaw)
	val += fmt.Sprintf("%s %s\n", chalk.Yellow.Color("Physical Trademark:"), c.PhysicalTrademark)
	val += fmt.Sprintf("%s %s\n", chalk.Yellow.Color("Idiosyncrasy:"), c.Idiosyncrasy)
	val += fmt.Sprintf("%s %s", chalk.Yellow.Color("Unfortunate Incident or Condition:"), c.UnfortunateIncident)
	return val
}

type PlayerClass interface {
	Level() int
	LevelUp()
	String() string
	StartingFeatureBlurb() string
	GetFeatures() []Feature
	Description() string
	GetItems() []Item
	GetWeapons() []Weapon
	GetStatMods() Stats
	GetHPDie() int
	GetClothingDie() int
	GetHatDie() int
	GetWeaponDie() int
	GetDevilsLuck() string
}

type Feature interface {
	String() string
}

func GenerateCharacter(extra bool) {
	pc := NewCharacter(extra)
	fmt.Printf("%s %s\n%s %s\n%s\n%s\n",
		Red("Name:"), pc.Name,
		Red("Class:"), pc.Class,
		pc.Class.Description(),
		Red("Features:"))
	for _, feature := range pc.Features {
		fmt.Println(feature)
	}
	fmt.Printf("\n%s %d\n", Red("HP:"), pc.HP)
	fmt.Printf("%s\n%s\n", Red("Stats:"), pc.Stats)
	fmt.Printf("%s %s\n", Red("Devil's Luck:"), pc.Class.GetDevilsLuck())

	fmt.Printf("\n%s\n%s", Red("Weapons:"), pc.Weapons)
	fmt.Printf("%s %s\n", Red("Clothing:"), pc.Clothing)
	fmt.Printf("%s %s\n", Red("Hat:"), pc.Hat)
	fmt.Printf("%s\n%s", Red("Equipment:"), pc.Gear)
	fmt.Printf("%s %s\n", Red("Money:"), pc.Money)

	fmt.Printf("\n%s\n", pc.Character)
	fmt.Printf("%s %s\n", chalk.Yellow.Color("Thing of Importance:"), pc.ThingOfImportance)
}

func NewCharacter(additionalClasses ...bool) *PlayerCharacter {
	pc := &PlayerCharacter{}
	pc.Class = GetClass(additionalClasses...)
	pc.GetStartingStats()
	pc.StartingHP()
	pc.StartingWeapon()
	pc.StartingClothing()
	pc.StartingHat()
	pc.StartGear()
	pc.Features = append(pc.Features, pc.Class.GetFeatures()...)
	pc.WhoIsThisPerson()
	pc.CreateThingOfImportance()
	pc.GetRandomName()
	return pc
}

func GetClass(additionalClasses ...bool) PlayerClass {
	res := utils.D(6)
	if len(additionalClasses) > 0 && additionalClasses[0] {
		res = utils.D(8)
	}
	switch res {
	case 1:
		return NewBrute()
	case 2:
		return NewRapscallion()
	case 3:
		return NewBuccaneer()
	case 4:
		return NewSwashbuckler()
	case 5:
		return NewZealot()
	case 6:
		return NewSorcerer()
	case 7:
		return NewHauntedSoul()
	case 8:
		return NewTallTale()
	}
	return nil
}

func (pc *PlayerCharacter) GetStartingStats() {
	pc.Stats = Stats{
		Strength:  GenerateStat(),
		Agility:   GenerateStat(),
		Presence:  GenerateStat(),
		Toughness: GenerateStat(),
		Spirit:    GenerateStat(),
	}
	pc.ApplyClassStats()
	pc.CarryingCapacity = 8 + pc.Stats.Strength
}

func (pc *PlayerCharacter) ApplyClassStats() {
	classStats := pc.Class.GetStatMods()
	pc.Stats.Strength = pc.Stats.Strength + classStats.Strength
	if pc.Stats.Strength < -3 {
		pc.Stats.Strength = -3
	}
	if pc.Stats.Strength > 6 {
		pc.Stats.Strength = 6
	}
	pc.Stats.Agility = pc.Stats.Agility + classStats.Agility
	if pc.Stats.Agility < -3 {
		pc.Stats.Agility = -3
	}
	if pc.Stats.Agility > 6 {
		pc.Stats.Agility = 6
	}
	pc.Stats.Presence = pc.Stats.Presence + classStats.Presence
	if pc.Stats.Presence < -3 {
		pc.Stats.Presence = -3
	}
	if pc.Stats.Presence > 6 {
		pc.Stats.Presence = 6
	}
	pc.Stats.Toughness = pc.Stats.Toughness + classStats.Toughness
	if pc.Stats.Toughness < -3 {
		pc.Stats.Toughness = -3
	}
	if pc.Stats.Toughness > 6 {
		pc.Stats.Toughness = 6
	}
	pc.Stats.Spirit = pc.Stats.Spirit + classStats.Spirit
	if pc.Stats.Spirit < -3 {
		pc.Stats.Spirit = -3
	}
	if pc.Stats.Spirit > 6 {
		pc.Stats.Spirit = 6
	}
}

func GenerateStat() int {
	r := utils.D(6) * utils.D(6) * utils.D(6)
	switch {
	case slices.Contains([]int{3, 4}, r):
		return -3
	case slices.Contains([]int{5, 6}, r):
		return -2
	case slices.Contains([]int{7, 8}, r):
		return -1
	case slices.Contains([]int{9, 10, 11, 12}, r):
		return 0
	case slices.Contains([]int{13, 14}, r):
		return 1
	case slices.Contains([]int{15, 16}, r):
		return 2
	case slices.Contains([]int{17, 18}, r):
		return 3
	}
	return 0
}

func (pc *PlayerCharacter) StartingHP() {
	pc.HP = utils.D(pc.Class.GetHPDie()) + pc.Stats.Toughness
	if pc.HP < 1 {
		pc.HP = 1
	}
}

func (pc *PlayerCharacter) StartingWeapon() {
	d := pc.Class.GetWeaponDie()
	if d != 0 {
		pc.Weapons = append(pc.Weapons, RollForWeapon(d))
	}
	pc.Weapons = append(pc.Weapons, pc.Class.GetWeapons()...)
}
func (pc *PlayerCharacter) StartingClothing() {
	d := pc.Class.GetClothingDie()
	if d != 0 {
		pc.Clothing = RollForClothing(d)
	}
}
func (pc *PlayerCharacter) StartingHat() {
	d := pc.Class.GetHatDie()
	if d != 0 {
		pc.Hat = RollForHat(d)
	}
}

func RollForWeapon(d int) Weapon {
	switch utils.D(d) {
	case 1:
		return Weapon{Item: Item{Name: "Marlinspike or Belaying Pin"}, Damage: "d4"}
	case 2:
		return Weapon{Item: Item{Name: "Knife or Bayonet"}, Damage: "d4"}
	case 3:
		return Weapon{Item: Item{Name: "Smallsword or Machete"}, Damage: "d4"}
	case 4:
		return Weapon{Item: Item{Name: "Cat O' Nine Tails", Extra: "10' reach"}, Damage: "d4"}
	case 5:
		return Weapon{Item: Item{Name: "Boarding Axe"}, Damage: "d6"}
	case 6:
		return Weapon{Item: Item{Name: "Cutlass"}, Damage: "d6"}
	case 7:
		return Weapon{Item: Item{Name: "Flintlock Pistol", Extra: "reload 2 actions, range 30'"}, Damage: "2d4"}
	case 8:
		return Weapon{Item: Item{Name: "Finely Crafted Rapier"}, Damage: "d8"}
	case 9:
		return Weapon{Item: Item{Name: "Boarding Pike", Extra: "10' reach"}, Damage: "d10"}
	case 10:
		return Weapon{Item: Item{Name: "Musket", Extra: "reload 2 actions, range 150'"}, Damage: "2d6"}
	}
	return Weapon{}
}
func RollForClothing(d int) Clothing {
	r := utils.D(d)
	switch {
	case r == 3 || r == 4:
		return Clothing{Item: Item{Name: "common clothes", Price: "2s"}}
	case r == 5:
		return Clothing{Item: Item{Name: "old uniform", Price: "8s"}}
	case r == 6:
		return Clothing{Item: Item{Name: "fancy clothes", Price: "250s", Extra: I("You look amazing!")}, Armor: ""}
	case r == 7:
		return Clothing{Item: Item{Name: "leather armor", Price: "20s"}, Armor: "tier1: -d2 damage"}
	case r == 8:
		return Clothing{Item: Item{Name: "hide armor", Price: "25s"}, Armor: "tier1: -d2 damage"}
	case r == 9:
		return Clothing{Armor: "tier2: -d4 damage", Item: Item{Name: "chain shirt", Price: "100s",
			Extra: "DR +2 on AGILITY tests including DEFENSE."}}
	case r == 10:
		return Clothing{Armor: "tier3: -d6 damage", Item: Item{Name: "conquistador plate", Price: "200s",
			Extra: "DR +4 on AGILITY tests, DEFENSE is DR +2. " + I("You'll most likely sink and drown in water.")}}
	}
	return Clothing{Item: Item{Name: "rags"}}
}
func RollForHat(d int) Hat {
	r := utils.D(d)
	switch {
	case r == 5:
		return Hat{Item: Item{Name: "wig", Price: "8s"}}
	case r == 6:
		return Hat{Item: Item{Name: "bandana", Price: "2s"}}
	case r == 7:
		return Hat{Item: Item{Name: "cavalier", Price: "15s"}}
	case r == 8:
		return Hat{Item: Item{Name: "bicorne", Price: "15s"}}
	case r == 9:
		return Hat{Item: Item{Name: "plain tricorne", Price: "10s"}}
	case r == 10:
		return Hat{Item: Item{Name: "fancy tricorne", Price: "90s"}}
	case r == 11:
		return Hat{Item: Item{Name: "metal lined hat", Price: "20s"}, Armor: "-1 damage"}
	case r == 12:
		return Hat{Armor: "-1 damage", Item: Item{Name: "morion", Price: "90s",
			Extra: I("(conquistador helmet)") + " You can choose to ignore all damage from one attack but the helmet breaks."}}
	}
	return Hat{}
}

func (pc *PlayerCharacter) StartGear() {
	pc.Gear = append(pc.Gear, pc.Class.GetItems()...)
	pc.Container = RollContainer()
	pc.Gear = append(pc.Gear, RollCheapGear(pc.Stats.Presence))
	pc.Gear = append(pc.Gear, RollFancyGear())
}

func RollCheapGear(presence int) Item {
	d4 := utils.D(4)
	d6 := utils.D(6)
	switch utils.D(12) {
	case 1:
		return Item{Name: "lantern", Price: fmt.Sprintf("%ds", d6*5+10), Extra: fmt.Sprintf("%d hours of oil", d6)}
	case 2:
		return Item{Name: fmt.Sprintf("%d candles", d4), Price: fmt.Sprintf("%ds", d4), Extra: "1 hour each"}
	case 3:
		return Item{Name: "30' of rope", Price: "4s", Extra: ""}
	case 4:
		return Item{Name: "shovel", Price: "5s", Extra: ""}
	case 5:
		return Item{Name: "medical kit", Price: "15s",
			Extra: Red(fmt.Sprintf("stops bleeding/poison/infection and heals %s HP. %d uses", B("d6"), presence+4))}
	case 6:
		return Item{Name: "weighted dice", Price: "", Extra: ""}
	case 7:
		return Item{Name: "flint & steel", Price: "3s", Extra: ""}
	case 8:
		return Item{Name: "hammer & nails", Price: "8s", Extra: ""}
	case 9:
		return Item{Name: "mess kit", Price: "8s", Extra: ""}
	case 10:
		return Item{Name: "pipe & tobacco pouch", Price: "10s", Extra: ""}
	case 11:
		return Item{Name: fmt.Sprintf("%d torches", d6), Price: fmt.Sprintf("%ds", d6*2), Extra: "1 hour each"}
	case 12:
		switch utils.D(10) {
		case 1:
			return Item{Name: "a pet snake"}
		case 2:
			return Item{Name: "a pet rat"}
		case 3:
			return Item{Name: "a pet lizard"}
		case 4:
			return Item{Name: "a pet monkey"}
		case 5:
			return Item{Name: "a pet parrot"}
		case 6:
			return Item{Name: "a pet cat"}
		case 7:
			return Item{Name: "a pet dog"}
		case 8:
			return Item{Name: "a pet hawk"}
		case 9:
			return Item{Name: "a pet hermit crab"}
		case 10:
			return Item{Name: "a pet fish in a jar"}
		}

	}
	return Item{}
}
func RollFancyGear() Item {
	switch utils.D(12) {
	case 1:
		return Item{Name: "compass", Price: "75s", Extra: ""}
	case 2:
		return Item{Name: "spyglass", Price: "150s", Extra: ""}
	case 3:
		return Item{Name: "fishing rod", Price: "25s", Extra: ""}
	case 4:
		return GetAncientRelic(utils.D(20))
	case 5:
		return Item{Name: "bottle of fine rum", Price: "10s", Extra: ""}
	case 6:
		return Item{Name: "old pocket watch", Price: "45s", Extra: ""}
	case 7:
		return Item{Name: "blanket & pillow", Price: "5s", Extra: ""}
	case 8:
		return Item{Name: "ink, quill, parchment", Price: "20s", Extra: ""}
	case 9:
		return Item{Name: "worn out book", Price: "", Extra: ""}
	case 10:
		return Item{Name: "tent", Price: "25s", Extra: ""}
	case 11:
		return Item{Name: "whetstone", Price: "5s", Extra: ""}
	case 12:
		var instrument string
		switch utils.D(10) {
		case 1:
			instrument = "concertina"
		case 2:
			instrument = "drum"
		case 3:
			instrument = "flute"
		case 4:
			instrument = "fiddle"
		case 5:
			instrument = "banjo"
		case 6:
			instrument = "horn"
		case 7:
			instrument = "hurdy-gurdy"
		case 8:
			instrument = "guitar"
		case 9:
			instrument = "mandolin"
		case 10:
			instrument = "voice of an angel"
		}
		return Item{Name: instrument, Price: "250s+", Extra: ""}
	}
	return Item{}
}

func RollContainer() string {
	switch utils.D(6) {
	case 1:
		return I("bucket") + " for " + B("4") + " normal-sized items"
	case 2:
		return I("bandolier") + " for " + B("6") + " small-sized items"
	case 3:
		return I("satchel") + " for " + B("8") + " normal-sized items"
	case 4:
		return I("backpack") + " for " + B("10") + " normal-sized items"
	case 5:
		return I("large sea chest") + " for " + B("20") + " normal-sized items"
	case 6:
		return I("dinghy")
	}
	return ""
}

func (pc *PlayerCharacter) WhoIsThisPerson() {
	background := RollBackground()
	bg := background.String()
	if background.RestOfBackground != "" {
		bg += I(", you have " + background.RestOfBackground)
	}
	pc.Character = Character{
		Background:          bg,
		DistinctiveFlaw:     RollDistinctiveFlaw(),
		PhysicalTrademark:   RollPhysicalTrademark(),
		Idiosyncrasy:        RollIdiosyncrasy(),
		UnfortunateIncident: RollUnfortunateIncident(),
	}
	pc.Money = background.Money
	if background.Item.Name != "" {
		pc.Gear = append(pc.Gear, background.Item)
	}
}

func (pc *PlayerCharacter) CreateThingOfImportance() {
	pc.ThingOfImportance = RollThingOfImportance()
	pc.Gear = append(pc.Gear, Item{Name: pc.ThingOfImportance})

}

type CharacterBackground struct {
	Item             Item
	Name             string
	Money            string
	RestOfBackground string
}

func (cb CharacterBackground) String() string {
	return cb.Name
}

func GetBackgroundList() []CharacterBackground {
	backgroundList := []CharacterBackground{}
	for k, v := range backgrounds {
		backgroundList = append(backgroundList, CharacterBackground{
			Name:             k,
			Money:            v[0],
			Item:             Item{Name: v[1]},
			RestOfBackground: v[2],
		})
	}
	return backgroundList
}

func RollBackground() CharacterBackground {
	r := utils.TableDie(100)
	if r < 19 {
		r = 19
	}
	r -= 19
	return GetBackgroundList()[r]
}
func RollDistinctiveFlaw() string {
	return distinctiveFlaws[utils.TableDie(20)]
}
func RollPhysicalTrademark() string {
	return physicalTrademarks[utils.TableDie(20)]
}
func RollIdiosyncrasy() string {
	return idiosyncrasies[utils.TableDie(20)]
}
func RollUnfortunateIncident() string {
	return unfortunateIncidents[utils.TableDie(20)]
}

func RollThingOfImportance() string {
	return thingsOfImportance[utils.TableDie(100)]
}

func (pc *PlayerCharacter) GetRandomName() {
	pc.Name = PirateName()
}
