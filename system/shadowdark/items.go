package shadowdark

import (
	"fmt"
	"slices"
	"strings"

	"github.com/genrpg/utils"
)

type MagicItem interface {
	String() string
	HasBenefit(...int)
	HasCurse()
	SetPersonality(MagicItemPersonality)
	Init()
}

type NamedMagicItem struct {
	Personality MagicItemPersonality
	Benefit     *string
	Curse       *string
	Name        string
}

func (n *NamedMagicItem) String() string {
	output := n.Name
	if n.Benefit != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Benefit"), *n.Benefit)
	}
	if n.Curse != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Curse"), *n.Curse)
	}
	personality := n.Personality.String()
	if personality != "" {
		output += fmt.Sprintf("\n%s", personality)
	}
	return output
}

func (n *NamedMagicItem) HasBenefit(num ...int) {
	n.Benefit = new(string)
}

func (n *NamedMagicItem) HasCurse() {
	n.Curse = new(string)
}

func (n *NamedMagicItem) SetPersonality(p MagicItemPersonality) {
	n.Personality = p
}

func (n *NamedMagicItem) Init() {
	*n = NewNamedMagicItem(*n)
}

func NewNamedMagicItem(seed NamedMagicItem) NamedMagicItem {
	if seed.Benefit != nil {
		*seed.Benefit = getUtilityBenefit()
	}
	if seed.Curse != nil {
		*seed.Curse = magicUtilityCurses[utils.TD(12)]
	}
	resolvePersonality(&seed.Personality)
	return seed
}

// Personality
type MagicItemPersonality struct {
	Virtue *string
	Flaw   *string
	Trait  *string
}

func resolvePersonality(seed *MagicItemPersonality) {
	if seed.Virtue != nil {
		*seed.Virtue = magicItemVirtue()
	}
	if seed.Flaw != nil {
		*seed.Flaw = magicItemFlaw()
	}
	if seed.Trait != nil {
		*seed.Trait = magicItemTraits[utils.TD(16)]
	}
}

func (p MagicItemPersonality) String() string {
	if p.Virtue == nil && p.Flaw == nil && p.Trait == nil {
		return ""
	}
	output := fmt.Sprintf("Personality: %s", *p.Trait)
	if p.Virtue != nil {
		output += fmt.Sprintf(" | Virtue: %s", *p.Virtue)
	}
	if p.Flaw != nil {
		output += fmt.Sprintf(" | Flaw: %s", *p.Flaw)
	}
	return output
}

var magicItemVirtues = [][2]any{{"Insists on protecting people and creatures it likes", nil}, {"Warns its wielder if it senses impending danger", nil}, {"Gladly translates Primordial for its wielder", nil}, {"Senses hiding creatures within near, but not exact place", nil}, {"Owed a favor by a", []string{"unicorn", "unicorn", "dragon", "noble"}}, {"Commands the respect of the followers of a god", nil}, {"Occasionally remembers useful ancient history", nil}, {"Imparts pleasant dreams and good sleep to its wielder", nil}, {"Coaches its wielder on the right things to say in a situation", nil}, {"Sometimes provides helpful strategic advice", nil}, {"Occasionally notices important details others have missed", nil}, {"Tries to mediate disagreements between conscious items", nil}, {"Calming presence to", []string{"dogs", "horses", "cats", "birds"}}, {"Has an extremely acute sense of smell", nil}, {"Knows the direction of the nearest running water", nil}, {"Lawful, intimidating to chaotic creatures", nil}, {"Neutral, intimidating to lawful and chaotic creatures", nil}, {"Chaotic, intimidating to lawful creatures", nil}, {"Has legitimate prophecies but isn't sure of their meaning", nil}, {"Can undo a great", []string{"evil", "lie", "spell", "alliance"}}}

func magicItemVirtue() string {
	virtue := magicItemVirtues[utils.TD(20)]
	output := virtue[0].(string)
	if virtue[1] != nil {
		output += " " + virtue[1].([]string)[utils.TD(4)]
	}
	return output
}

var magicItemFlaws = [][2]any{{"Afraid of", []string{"the dark", "vermin", "heights", "water"}}, {"Preferred a past owner and always draws comparisons", nil}, {"Chatters while wielder is trying to concentrate", nil}, {"Dislikes", []string{"elves", "dwarves", "humans", "goblins"}}, {"Tries to get wielder into fights so it \"has something to do\"", nil}, {"Does not want to be separated from wielder for any reason", nil}, {"Objects to", []string{"gambling", "carousing", "stealth", "theft"}}, {"Accuses everyone of lying; is correct once in a while", nil}, {"Won't harm creatures", []string{"lawful", "lawful", "neutral", "chaotic"}}, {"Believes its wielder is a pawn in its apocalyptic scheme", nil}, {"Constantly tries to escape its current wielder", nil}, {"Demands its wielder observe its god's strict rituals", nil}, {"Insists on being reunited with its creator, living or dead", nil}, {"Can't stand other conscious magic items", nil}, {"Refuses to be used for \"unimportant\" or \"boring\" tasks", nil}, {"Purposefully goes magically inert when mad at its wielder", nil}, {"Insists on being meticulously cleaned every day", nil}, {"Loves the color purple and despises all other colors", nil}, {"Objects to", []string{"negotiating", "fighting", "fighting", "planning"}}, {"Pretends to know information it doesn't know", nil}}

func magicItemFlaw() string {
	flaw := magicItemFlaws[utils.TD(20)]
	output := flaw[0].(string)
	if flaw[1] != nil {
		output += " " + flaw[1].([]string)[utils.TD(4)]
	}
	return output
}

var magicItemTraits = []string{"Imperious", "Polite", "Puritanical", "Charming", "Anxious", "Righteous", "Critical", "Theatrical", "Bossy", "Noble", "Greedy", "Protective", "Impulsive", "Brave", "Vicious", "Loyal"}

// Armor
type MagicItemArmor struct {
	Personality MagicItemPersonality
	Bonus       *int
	Benefit     *string
	Curse       *string
	Name        string
	Type        string
	Feature     string
	Mithral     bool
}

var magicArmorFeatures = []string{"Demonic horned face", "Oak leaf motif", "Studded with shark teeth", "Dragon scales", "Bone or metal spikes", "Faint arcane runes", "Turtle shell plating", "Made of scorpion chitin", "Gilded metal/gold thread", "Scorched, smells burned", "Pearl-white fish scales", "Oozes blood", "Festooned with fungi", "Distant sound of ocean", "Set with crystals", "Draped in holy symbols", "Exudes tree sap", "Blurry, indistinct edges", "Large golden cat eye", "Covered in frost"}
var magicArmorBenefits = []string{"Once per day, deflect a ranged attack that would hit you", "Checks to stabilize you are easy (DC 9)", "You cannot be knocked over while you are conscious", "Undetected creatures do not have advantage to attack you", "You know Diabolic and are immune to fire, lava, and magma", "You are immune to the curses of one item you choose", "Once per day, gain advantage on all attacks for 3 rounds", "You have a +4 bonus to your death timers", "Gain immunity to a poison after suffering its effects once", "You know Celestial and can fly for 3 rounds once per day", "Treat critical hits against you as normal hits", "Ignore any damage dealt to you of 3 points or below"}
var magicArmorCurses = []string{"You take 2d10 damage if you remove this armor", "Your party cannot add CHA bonuses to reaction checks", "Mounts fear you and will not allow you to ride them", "DC 15 WIS first round of combat or attack nearest creature", "You take double damage from blunt/bludgeoning weapons", "Armor uses 5 gear slots and is extremely loud and clunky", "Ranged attacks against you have advantage", "Treat a natural 1 attack roll against you as a critical hit", "Beneficial spells that target you are hard to cast (DC 15)", "You have disadvantage on Dexterity checks", "There's a secret 1-in-6 chance each NPC ally will betray you", "You take double damage from silvered weapons"}

func (a *MagicItemArmor) Init() {
	*a = NewMagicArmor(*a)
}

func NewMagicArmor(seed MagicItemArmor) MagicItemArmor {
	if seed.Bonus == nil {
		seed.Bonus = new(int)
		switch utils.D(6, 2) {
		case 2, 3, 4, 5:
			*seed.Bonus = 0
		case 6, 7, 8:
			*seed.Bonus = 1
		case 9, 10, 11:
			*seed.Bonus = 2
		default:
			*seed.Bonus = 3
		}
	}
	if seed.Type == "" {
		switch utils.D(6, 2) {
		case 2, 3, 4, 5:
			seed.Type = "Leather"
		case 6, 7:
			seed.Type = "Chain mail"
		case 8, 9:
			seed.Type = "Shield"
		case 10, 11:
			seed.Type = "Plate mail"
		case 12:
			seed.Mithral = true
		}
		if seed.Mithral {
			seed.Type = mithralArmor()
		}
	}
	if seed.Feature == "" {
		seed.Feature = magicArmorFeatures[utils.TD(20)]
	}
	if seed.Benefit != nil {
		*seed.Benefit = magicArmorBenefits[utils.TD(12)]
	}
	if seed.Curse != nil {
		*seed.Curse = magicArmorCurses[utils.TD(len(magicArmorCurses))]
	}
	resolvePersonality(&seed.Personality)
	return seed
}

func mithralArmor() string {
	switch utils.D(6, 2) {
	case 6, 7:
		return "Chain mail"
	case 8, 9:
		return "Shield"
	case 10, 11:
		return "Plate mail"
	default:
		return mithralArmor()
	}
}

func (a *MagicItemArmor) String() string {
	var output string
	if a.Name != "" {
		output += fmt.Sprintf("%s\n", a.Name)
	}
	if a.Bonus != nil && *a.Bonus > 0 {
		output += fmt.Sprintf("+%d ", *a.Bonus)
	}
	if a.Mithral {
		output += "mithral "
	}
	if a.Type != "" {
		output += strings.ToLower(a.Type)
	}
	if a.Feature != "" {
		output += fmt.Sprintf("\n%s", utils.I(a.Feature))
	}
	if a.Benefit != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Benefit"), *a.Benefit)
	}
	if a.Curse != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Curse"), *a.Curse)
	}
	personality := a.Personality.String()
	if personality != "" {
		output += fmt.Sprintf("\n%s", personality)
	}
	return output
}

func (a *MagicItemArmor) HasBenefit(num ...int) {
	a.Benefit = new(string)
}

func (a *MagicItemArmor) HasCurse() {
	a.Curse = new(string)
}

func (a *MagicItemArmor) SetPersonality(p MagicItemPersonality) {
	a.Personality = p
}

type MagicItemPotion struct {
	MagicItemPersonality
	Features []string
	Mixing   []string
	Name     string
	Benefit  string
	Curse    string
}

var magicPotionFeatures1 = []string{"Spicy", "Clear as water", "Deep blue", "Citrus smell", "Sulfurous", "Fizzy", "Chilly", "Blood red"}
var magicPotionFeatures2 = []string{"Pickled spider inside ", "Green fumes", "Tiny stars and moon", "Gold flakes in liquid", "Swirling vortex", "Quiet whistling", "Rattles and shakes"}
var magicPotionFeatures3 = []string{"Blood red", "Eyeball inside", "Bubbling", "Purple streaks", "Flames on surface", "Floral smell", "Skull on bottle", "Warm", "Large molar inside", "Pink starbursts"}

var magicPotionMixingEffects1 = []string{"Drinker floats 2 rounds", "Gains personality: 1 virtue", "Becomes totally inert", "Reroll 2 new benefits", "Curse effects halved", "Gains 1 additional curse", "Drinker forgets 1 language", "Gains 1 additional benefit", "Gains personality: 1 flaw", "All effects are inverted", "Lose all benefits", "Drinker +1 to random stat"}
var magicPotionMixingEffects2 = []string{"Fumes: DC 12 CON or 1d4 dmg", "Loses all curses", "Gains personality: 1 flaw", "Drinker invisible 2 rounds", "Double one benefit's effects", "Explodes: DC 12 DEX or 1d6 dmg", "Gains personality: 1 virtue", "Benefit effects halved", "Double's one curse's effects", "Drinker DC 12 CON or 1d8 dmg", "Drinker -1 to random stat", "1 random effect is permanent"}

var magicPotionBenefits = [][2]any{{"Immune to %s for 5 rounds", []string{"fire", "cold", "electricity", "poison"}}, {"Heals %s", []string{"1d4", "2d6", "3d8", "4d10"}}, {"Read the minds of all creatures within near for 1 hour", nil}, {"Fly a near distance for 5 rounds", nil}, {"For 5 rounds, move far on your turn and still take an action", nil}, {"Become invisible for 5 rounds", nil}, {"Breathe underwater and know Merran language for 1 hour", nil}, {"A stat becomes 18 (+4) for 5 rounds", nil}, {"Turn into purple, flying gas for 5 rounds", nil}, {"Cures any disease or affliction affecting drinker", nil}, {"Speak to and understand animals for 1 hour", nil}, {"You are immune to all damage for 5 rounds", nil}}
var magicPotionCurses = [][2]any{{"DC 15 WIS check or attack nearest creature for 3 rounds", nil}, {"Turn into a 1 HP newt for 3 rounds", nil}, {"A stat becomes 3 (-4) for 1 hour", nil}, {"DC 15 CON check or take 2d10 damage", nil}, {"Forget all languages you know for 1 hour", nil}, {"Shrink to half size and disadvantage on attacks for 5 rounds", nil}, {"Sing at the top of your lungs for 3 rounds", nil}, {"You become magnetic to all metal near to you for 1 hour", nil}, {"You are compelled to jump into any pits you see for 1 hour", nil}, {"DC 15 CON check or go blind for 5 rounds", nil}, {"You are the source of an antimagic shell spell for 1 hour", nil}, {"%s petrify for 5 rounds", []string{"Arms", "Legs"}}}

func getPotionBenefit() string {
	entry := magicPotionBenefits[utils.TD(12)]
	benefit := entry[0].(string)
	if entry[1] != nil {
		options := entry[1].([]string)
		benefit = fmt.Sprintf(entry[0].(string), options[utils.TD(len(options))])
	}
	return benefit
}

func getPotionCurse() string {
	entry := magicPotionCurses[utils.TD(len(magicPotionCurses))]
	benefit := entry[0].(string)
	if entry[1] != nil {
		options := entry[1].([]string)
		benefit = fmt.Sprintf(entry[0].(string), options[utils.TD(len(options))])
	}
	return benefit
}

type MagicItemScrollWand struct {
	Personality MagicItemPersonality
	Benefit     *string
	Curse       *string
	Name        string
	Feature     string
	Spell       string
	SpellTier   int
}

var magicScrollFeatures = []string{"Branded on leather", "Etched on copper leaf", "Faded papyrus", "Stained parchment roll", "Carved into bone", "Chiseled on stone slats", "Etched into glass", "Tattooed on dragon skin"}
var magicWandFeatures = []string{"Carved from bone", "Blinking eye in handle", "Sleek starmetal", "Polished wood", "Obsidian with ivory tips", "Electrical sparks", "Jagged crystal", "Made of tiny skulls"}
var tier1spells = []string{"Alarm", "Burning hands", "Charm person", "Detect magic", "Feather fall", "Floating disk", "Hold portal", "Light", "Mage armor", "Magic missile", "Protection from evil", "Sleep"}
var tier2spells = []string{"Acid arrow", "Alter self", "Detect thoughts", "Fixed object", "Hold person", "Invisibility", "Knock", "Levitate", "Mirror image", "Misty step", "Silence", "Web"}
var tier3spells = []string{"Animate dead", "Dispel magic", "Fabricate", "Fireball", "Fly", "Gaseous form", "Illusion", "Lightning bolt", "Magic circle", "Protection from energy", "Sending", "Speak with dead"}
var tier4spells = []string{"Arcane eye", "Cloudkill", "Confusion", "Control water", "Dimension door", "Divination", "Passwall", "Polymorph", "Resilient sphere", "Stoneskin", "Telekinesis", "Wall of force"}
var tier5spells = []string{"Antimagic shell", "Create undead", "Disintegrate", "Hold monster", "Plane shift", "Power word kill", "Prismatic orb", "Scrying", "Shapechange", "Summon extraplanar", "Teleport", "Wish"}

func (s *MagicItemScrollWand) Init() {
	*s = NewMagicScrollWand(*s)
}

func NewMagicScrollWand(seed MagicItemScrollWand) MagicItemScrollWand {
	if seed.SpellTier != 0 {
		switch seed.SpellTier {
		case 1:
			seed.Spell = tier1spells[utils.TD(len(tier1spells))]
		case 2:
			seed.Spell = tier2spells[utils.TD(len(tier2spells))]
		case 3:
			seed.Spell = tier3spells[utils.TD(len(tier3spells))]
		case 4:
			seed.Spell = tier4spells[utils.TD(len(tier4spells))]
		case 5:
			seed.Spell = tier5spells[utils.TD(len(tier5spells))]
		}
	}
	if seed.Feature == "" {
		switch {
		case strings.Contains(seed.Name, "scroll"):
			seed.Feature = magicScrollFeatures[utils.TD(len(magicScrollFeatures))]
		case strings.Contains(seed.Name, "wand"):
			seed.Feature = magicWandFeatures[utils.TD(len(magicWandFeatures))]
		}
	}
	if seed.Benefit != nil {
		switch utils.D(6, 2) {
		case 12:
			*seed.Benefit = getWeaponBenefit()
		case 9, 10, 11:
			*seed.Benefit = getUtilityBenefit()
		case 7, 8:
			*seed.Benefit = getPotionBenefit()
		default:
			*seed.Benefit = magicArmorBenefits[utils.TD(12)]
		}
	}
	if seed.Curse != nil {
		switch utils.D(6, 2) {
		case 12:
			*seed.Curse = getWeaponCurse()
		case 9, 10, 11:
			*seed.Curse = magicUtilityCurses[utils.TD(len(magicUtilityCurses))]
		case 7, 8:
			*seed.Curse = getPotionCurse()
		default:
			*seed.Curse = magicArmorCurses[utils.TD(len(magicArmorCurses))]
		}
	}
	resolvePersonality(&seed.Personality)
	return seed
}

func (s *MagicItemScrollWand) String() string {
	// TODO: Resolve random scrolls and wands
	// Determine if scroll or wand based on name "Scroll of ..." & "Magic wand, ..."
	// TODO:
	var output string
	switch {
	case strings.Contains(s.Name, "scroll"):
		s.Name = fmt.Sprintf("Scroll of %s", utils.I(utils.B(s.Spell)))
	case strings.Contains(s.Name, "wand"):
		s.Name = fmt.Sprintf("Magic wand, %s", utils.I(utils.B(s.Spell)))
	}
	output += s.Name
	if s.Feature != "" {
		output += fmt.Sprintf("\n%s", utils.I(s.Feature))
	}
	if s.Benefit != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Benefit"), *s.Benefit)
	}
	if s.Curse != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Curse"), *s.Curse)
	}
	personality := s.Personality.String()
	if personality != "" {
		output += fmt.Sprintf("\n%s", personality)
	}
	return output
}

func (s *MagicItemScrollWand) HasBenefit(num ...int) {
	s.Benefit = new(string)
	*s.Benefit = magicArmorBenefits[utils.TD(12)]
}

func (s *MagicItemScrollWand) HasCurse() {
	s.Curse = new(string)
	*s.Curse = magicArmorCurses[utils.TD(len(magicArmorCurses))]
}

func (s *MagicItemScrollWand) SetPersonality(p MagicItemPersonality) {
	s.Personality = p
}

type MagicItemUtility struct {
	MagicItemPersonality
	Name    string
	Type    string
	Feature string
	Benefit string
	Curse   string
}

var magicUtilityTypes = []string{"Brooch", "Ring", "Boots", "Cloak", "Amulet", "Flask", "Tome", "Circlet", "Eyepatch", "Gauntlets", "Holy symbol", "Hat", "Goblet", "Helm", "Statuette", "Goggles", "Bag", "Rock", "Surcoat"}
var magicUtilityFeatures = []string{"Shaped like a raven", "Iridescent", "Cruel spikes and spines", "Made from a big frog", "Gem-studded", "Gold thread/hardware", "Made of basilisk hide", "Possessed by a spirit", "Made of shaped smoke", "Covered in small thorns", "Made with rare feathers", "Has tiny wings", "Slowly changes colors", "Shaped like a bat", "Tarnished silver hardware", "Made of spidersilk", "Hums quiet, sweet tones", "Jolt of pain at first touch", "Throbs like a heart"}
var magicUtilityBenefits = [][2]any{{"You can't be magically scryed upon or detected", nil}, {"Connects to an interdimensional pocket with 5 gear slots", nil}, {"%s becomes 18 (+4) while using/wearing item", []string{"STR", "DEX", "CON", "INT", "WIS", "CHA"}}, {"Once per day, teleport a near distance", nil}, {"Harmful spells that target you are DC 15 to cast", nil}, {"You're immune to %s", []string{"fire", "cold", "electricity", "poison"}}, {"Sense secret doors when they're within close range", nil}, {"You can see invisible and incorporeal creatures", nil}, {"Your movement isn't hindered by any terrain", nil}, {"You can hold your breath for 1 hour", nil}, {"You do not need to eat or drink to survive", nil}, {"You can walk on non-solid surfaces for 2 rounds at a time", nil}}
var magicUtilityCurses = []string{"Slowly rots all other non-magical items that touch it", "Deals 1d4 damage and leaves blisters whenever used", "Item attracts bad weather to its location", "You cannot be healed by magic; only by resting", "Crashes like a gong whenever wielder slays a creature", "Item attracts all undead within a far distance", "Temporarily loses magic if doused in water", "You have disadvantage on CON checks", "You are compelled to light parchment objects on fire", "You must drink blood once a day or take 1d8 damage", "Item must eat 1d10 gp a day or it loses its magic until fed", "Item has horrid smell that makes all your CHA checks hard"}

func getUtilityBenefit() string {
	entry := magicUtilityBenefits[utils.TD(12)]
	benefit := entry[0].(string)
	if entry[1] != nil {
		options := entry[1].([]string)
		benefit = fmt.Sprintf(entry[0].(string), options[utils.TD(len(options))])
	}
	return benefit
}

type MagicItemWeapon struct {
	Personality MagicItemPersonality
	Benefits    *[]string
	Curse       *string
	Bonus       *int
	Name        string
	Type        string
	Feature     string
}

var magicWeaponFeatures = []string{"Trails sparkles", "Starmetal", "Rusted and chipped", "Gem in pommel/handle", "Drips green ichor", "Moon motif and silvered", "Galaxies swirl on surface", "Ironwood", "Rune-scribed", "Faint, ghostly aura", "Inlaid with gold", "Trails incense", "Studded with gemstones", "Sparks dance on surface", "Shaped like an animal", "Carved from granite", "Dragonbone hardware", "Whispers in a language", "Drips ocean water", "Turns blood to rose petals"}
var magicWeaponTypes = []string{"Arrows", "Bastard sword", "Bastard sword", "Club", "Crossbow", "Crossbow bolts", "Dagger", "Dagger", "Greataxe", "Greatsword", "Javelin", "Longbow", "Longsword", "Longsword", "Mace", "Shortbow", "Shortsword", "Shortsword", "Staff", "Warhammer"}
var magicWeaponBenefits = [][2]any{{"Cut or smash through any material", nil}, {"Once per day, ignites for 5 rounds, deals 1d4 extra damage", nil}, {"DC 15 CHA check to command a wild animal within far", nil}, {"Behead the enemy on a critical hit", nil}, {"When you hit a creature, learn its True Name (see pg. 319)", nil}, {"Shoot a bolt of energy near with DEX, 1d6 damage", nil}, {"Once per day, deflect a melee attack that would hit you", nil}, {"Regain 1d6 hit points when you slay a creature", nil}, {"You have advantage on initiative rolls", nil}, {"Has thrown property (pg. 37), near distance, returns to you", nil}, {"Double damage to", []string{"undead", "undead", "demons", "dragons"}}, {"Reroll natural 1s once each when attacking with this weapon", nil}}
var magicWeaponCurses = [][2]any{{"You can't see", []string{"undead", "demons", "snakes", "spiders"}}, {"You are compelled to swallow all gemstones at first sight", nil}, {"Burn a straw doll daily or weapon temporarily loses magic", nil}, {"Any light source you hold immediately extinguishes", nil}, {"You must loudly praise a god whenever you see its symbol", nil}, {"Venomous creatures always target you with attacks", nil}, {"You turn into a rat every day at midnight for one hour", nil}, {"Your checks to swim are always extreme (DC 18)", nil}, {"You are burned by the touch of gold", nil}, {"Bathe weapon in blood daily or it temporarily loses its magic", nil}, {"You cannot wear armor while wielding this weapon", nil}, {"Weapon can possess you by winning contested CHA (+2)", nil}}

func getWeaponBenefit(iter ...int) string {
	var roll int
	if len(iter) == 0 {
		roll = utils.TD(12)
	} else {
		roll = iter[0]
	}
	entry := magicWeaponBenefits[roll]
	benefit := entry[0].(string)
	if entry[1] != nil {
		options := entry[1].([]string)
		benefit += " " + options[utils.TD(len(options))]
	}
	return benefit
}

func getWeaponCurse() string {
	entry := magicWeaponBenefits[utils.TD(len(magicWeaponCurses))]
	benefit := entry[0].(string)
	if entry[1] != nil {
		options := entry[1].([]string)
		benefit += " " + options[utils.TD(len(options))]
	}
	return benefit
}

func (w *MagicItemWeapon) Init() {
	*w = NewMagicWeapon(*w)
}

func NewMagicWeapon(seed MagicItemWeapon) MagicItemWeapon {
	if seed.Bonus == nil {
		seed.Bonus = new(int)
		switch utils.D(6, 2) {
		case 2, 3:
			*seed.Bonus = 0
		case 4, 5, 6, 7, 8, 9:
			*seed.Bonus = 1
		case 10, 11:
			*seed.Bonus = 2
		default:
			*seed.Bonus = 3
		}
	}
	if seed.Type == "" {
		seed.Type = magicWeaponTypes[utils.TD(20)]
	}
	if seed.Type == "Arrows" || seed.Type == "Crossbow bolts" {
		seed.Type += fmt.Sprintf(" (#: %d)", utils.D(6)+utils.D(6))
	}
	if seed.Feature == "" {
		seed.Feature = magicWeaponFeatures[utils.TD(20)]
	}
	if seed.Benefits != nil {
		var rolls []int
		for i, _ := range *seed.Benefits {
			roll := recursiveConsumedRoll(rolls, 12, 0)
			rolls = append(rolls, roll)
			(*seed.Benefits)[i] = getWeaponBenefit(roll)
		}
	}
	if seed.Curse != nil {
		*seed.Curse = getWeaponCurse()
	}
	resolvePersonality(&seed.Personality)
	return seed
}

func (w *MagicItemWeapon) String() string {
	var output string
	if w.Name != "" {
		output += fmt.Sprintf("%s\n", w.Name)
	}
	if w.Bonus != nil && *w.Bonus > 0 {
		output += fmt.Sprintf("+%d ", *w.Bonus)
	}
	if w.Type != "" {
		output += strings.ToLower(w.Type)
	}
	if w.Feature != "" {
		output += fmt.Sprintf("\n%s", utils.I(w.Feature))
	}
	for _, benefit := range *w.Benefits {
		output += fmt.Sprintf("\n%s: %s", utils.B("Benefit"), benefit)
	}
	if w.Curse != nil {
		output += fmt.Sprintf("\n%s: %s", utils.B("Curse"), *w.Curse)
	}
	personality := w.Personality.String()
	if personality != "" {
		output += fmt.Sprintf("\n%s", personality)
	}
	return output
}

func (w *MagicItemWeapon) HasBenefit(num ...int) {
	var benefits []string
	if len(num) > 0 {
		benefits = make([]string, num[0], num[0])
	} else {
		benefits = make([]string, 1, 1)
	}
	w.Benefits = &benefits
}

func (w *MagicItemWeapon) HasCurse() {
	w.Curse = new(string)
}

func (w *MagicItemWeapon) SetPersonality(p MagicItemPersonality) {
	w.Personality = p
}

func recursiveConsumedRoll(rolls []int, size, iter int) int {
	roll := utils.TD(size)
	if slices.Contains(rolls, roll) && iter < size {
		return recursiveConsumedRoll(rolls, size, iter+1)
	}
	return roll
}
