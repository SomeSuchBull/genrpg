package knave

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

func getSpell(i int, spellParts ...string) string {
	switch {
	case i <= 2:
		return fmt.Sprintf("%s %s", spellParts[0], spellParts[1])
	case i <= 5:
		return fmt.Sprintf("The %s %s %s", spellParts[0], spellParts[1], spellParts[2])
	case i <= 8:
		return fmt.Sprintf("%s’s %s %s", spellParts[0], spellParts[1], spellParts[2])
	default:
		return fmt.Sprintf("%s’s %s %s %s", spellParts[0], spellParts[1], spellParts[2], spellParts[3])
	}
}

var spellFormulae = [][][]string{
	{elements, forms},
	{effects, forms},
	{effects, elements},
	{qualities, elements, forms},
	{qualities, effects, forms},
	{qualities, effects, elements},
	{wizardNames, elements, forms},
	{wizardNames, effects, forms},
	{wizardNames, effects, elements},
	{wizardNames, qualities, elements, forms},
	{wizardNames, qualities, effects, forms},
	{wizardNames, qualities, effects, elements},
}

// GetSpell returns a random spell
func GetRandomSpell(verbose bool) string {
	verboseOutput := ""
	spellParts := []string{}
	initialRoll := rand.Intn(12)
	spellFormulaOutput := spellFormulaeOutput[initialRoll]
	verboseOutput += tableRoll("Spell Formulae", initialRoll, spellFormulaOutput)
	spellFormulaRe := regexp.MustCompile(`\[([^\]]*)\]`)
	spellFormulaParts := []string{}
	for _, match := range spellFormulaRe.FindAllStringSubmatch(spellFormulaOutput, -1) {
		if len(match) > 1 {
			spellFormulaParts = append(spellFormulaParts, match[1])
		}
	}
	for i, table := range spellFormulae[initialRoll] {
		roll := rand.Intn(100)
		spellPart := table[roll]
		verboseOutput += tableRoll(spellFormulaParts[i], roll, spellPart)
		spellPart = recursiveTableRoll(spellPart, &verboseOutput)
		spellParts = append(spellParts, spellPart)
	}
	spell := getSpell(initialRoll, spellParts...)
	if verbose {
		verboseOutput += fmt.Sprintf("Spell: %s", spell)
		return verboseOutput
	}
	return spell
}

func GetSpell(i int64, verbose bool) string {
	verboseOutput := ""
	intelligence := int(i)
	roll := rand.Intn(100)
	spell := setSpells[roll]
	if intelligence > 0 {
		spell = strings.ReplaceAll(spell, "INT", fmt.Sprintf("%d", intelligence))
	}
	if verbose {
		verboseOutput += fmt.Sprintf("----\nRoll: %-3.02d\n%s", roll, spell)
		return verboseOutput
	}
	return spell
}

var setSpells = []string{
	`ADHERE
INT objects become sticky enough to hold a PC to a ceiling. Lasts until washed.`,
	`ANIMAL FRIENDSHIP
INT animals obey your orders as well as a trained dog for one day.`,
	`ANIMATE OBJECT
INT objects obey your orders. They move 15’ per round.`,
	`ANTROPOMOPHIZE
INT animals gain human intelligence for one day.`,
	`ARCANE EYE
You create a magic eye that flies around under your control for INT turns.
You can see through it as well as your normal eyes.`,
	`ASTRAL PRISON
An object is frozen in time and space within an invulnerable crystal shell for INT turns.`,
	`ATTRACT
INT + 1 objects are strongly magnetically attracted to each other if they come within 10’.`,
	`AUDITORY ILLUSION
You can create illusory sounds that seem to come from INT directions of your choice.`,
	`BABBLE
INT creatures must loudly and clearly repeat everything you think. They are otherwise mute.`,
	`BEAST FROM
You and your possessions turn into an animal for up to INT days.`,
	`BEFUDDLE
A creature is unable to form short-term memories for INT turns.`,
	`BEND FATE
Roll INT + 1 d20s. After this point, when any creature you can see makes a check,
use and discard one of the rolled results until they are all gone.`,
	`BODY SWAP
You switch bodies with a creature you touch for INT turns. If one body dies, the other dies as well.`,
	`CATHERINE
A woman wearing a blue dress appears for INT hours. She will obey polite, safe requests.`,
	`CHARM
INT humanoids believe they are close friends with you until proven otherwise.`,
	`COMMAND
A creature obeys a single, INT-word command that doesn’t harm it.`,
	`COMPREHEND
You are fluent in all languages for INT hours.`,
	`CONTROL PLANTS
Plants within INT × 10’ obey you. They move 5’ per round.`,
	`CONTROL WEATHER
You control your hex’s weather for INT hours.`,
	`DETECT MAGIC
Anything magical within line of sight glows and reveals its properties on request.
Lasts 1 day or until you make INT requests.`,
	`DISASSEBLE
INT body parts may be detached at will. You can still control them. Lasts until they are reattached.`,
	`DISGUISE
You may alter the look of INT humanoids as long as they remain humanoid.
Lasts until the subjects speak.`,
	`DISPLACE
An object appears to be up to INT × 10’ from its actual position.`,
	`DUPLICATE
Create INT fragile, Porcelain copies of items you can see.`,
	`EARTHQUAKE
The ground shakes violently for INT rounds.`,
	`ELASTICITY
Your body can stretch up to INT × 10’.`,
	`ELEMENTAL WALL
Creates a wall of ice or fire INT × 40’ long, 5’ wide and 10’ tall.
The wall can curve however you want.`,
	`FILCH
INT visible items teleport to your hands.`,
	`FOG CLOUD
Fog spreads out in an INT 10’ radius from you. Fades in one turn.`,
	`GRAVITY SHIFT
INT creatures can alter their “down” direction at will.`,
	`GREED
INT creatures become obsessed with possessing a visible item.`,
	`HASTE
INT creatures have their movement speed tripled.`,
	`HATRED
INT creatures start attacking each other for one turn or until one dies.`,
	`HEAR WHISPERS
A creature can hear all sounds up to 120’ away for INT turns.`,
	`HOVER
Make INT objects hover 2’ above the ground, frictionless.
They can support the weight of up to INT people.`,
	`HYPNOTIZE
A creature enters a trance and will answer INT yes or no questions.`,
	`ICY TOUCH
An ice layer spreads across a surface, up to INT × 10’ in radius.`,
	`INCREASE GRAVITY
The gravity within INT × 10’ of you triples.`,
	`INVISIBLE TETHER
INT objects within 10’ of each other cannot be moved more than 10’ apart from each other.`,
	`KNOCK
INT locks unlock.`,
	`LEAP
You can jump up to INT × 10’.`,
	`LIQUID AIR
The air within INT × 10’ of you becomes swimmable.`,
	`LOCK
A door cannot be opened by mundane means for INT turns.`,
	`MAGIC SUPPRESOR
All magic is nullified while within INT × 10’ of you.`,
	`MANSE
A furnished house with INT rooms appears for 1 day.
It has no food or gear and does not count as a safe haven.`,
	`MARBLE MADNESS
Your pockets refill with marbles every round for INT rounds.`,
	`MASQUERADE
All creatures within INT × 10’ of you are compelled to dance.`,
	`MINIATURIZE
You and INT other touched creatures become mouse-sized.`,
	`MIRROR IMAGE
INT illusory copies of you, under your control, appear.`,
	`MIRRORWALK
A mirror becomes a gate to another mirror you touched today.`,
	`MULTIARM
You gain INT extra arms.`,
	`NIGHT SPHERE
An INT × 40’ wide sphere of total darkness appears.`,
	`OBJECTIFY
INT willing creatures become inanimate, immobile objects of your choice for as long as they wish.
They can still hear and see.`,
	`OOZE FORM
Your body and gear become living slime for INT turns.`,
	`PACIFY
INT creatures develop an intense hatred of violence unless attacked.`,
	`PHANTOM COACH
A coach scoops up INT creatures (who are outdoors) and deposits them in a random adjacent hex.`,
	`PHOBIA
INT creatures become terrified of an object.`,
	`PIT
A pit 10’ wide and INT × 10’ deep opens in the ground.`,
	`PRIMEVAL SURGE
An object grows to the size of an elephant for INT turns. If it is a creature, it is enraged.`,
	`PSYCHOMETRY
The GM answers INT yes or no questions about an object.`,
	`PULL
An object of any size is pulled directly towards you with the force of INT men for one round.`,
	`PUSH
An object of any size is pushed directly away from you with the force of INT men for one round.`,
	`RAISE DEAD
INT unarmed skeletons rise from the ground to serve you.`,
	`READ MIND
You can hear the surface thoughts of creatures for INT turns.`,
	`REPEL
INT + 1 objects are strongly magnetically repelled from each other if they come within 10’.`,
	`SCRY
You can share the vision of a creature you touched today for INT turns.`,
	`SCULPT ELEMENTS
Inanimate material acts like clay in your hands for INT turns.`,
	`SHROUD
INT creatures are invisible for as long as they can hold their breath (CON × 3 rounds).`,
	`SHUFFLE
INT creatures switch places randomly.`,
	`SILENCE
All sound is deadened within 10’ of you for INT turns.`,
	`SLEEP
INT creatures fall asleep.`,
	`SMOKE FORM
Your body and gear become living smoke for INT turns.`,
	`SNAIL KNIGHT
In 10 minutes, a knight atop a giant snail rides into view.
He may aid you for INT days if he finds you worthy. The snail cannot move faster than a walk.`,
	`SNIFF
A creature can smell all scents up to 120’ away for INT turns.`,
	`SORT
Inanimate items sort themselves according to INT categories you set.
The categories must be visually verifiable.`,
	`SPEAK WITH DEAD
The spirit of a touched dead body appears and will answer INT questions (if it can).`,
	`SPECTACLE
A clearly unreal illusion appears under your control for INT days.
It may be up to the size of a palace and has full motion and sound.`,
	`SPELLSEIZE
Cast this as a reaction to another spell of level INT or less being cast
to make a temporary copy of it that you can cast within 1 day.`,
	`SPIDER CLIMB
You can climb surfaces like a spider for INT turns.`,
	`SUMMON CUBE
You may summon or banish a 5’ cube of earth 5 times per round for INT rounds.
Cubes must be affixed to the earth or to other cubes.`,
	`SUMMON IDOL
A carved stone statue up to INT × 10’ tall rises from the ground.`,
	`SWARM
You become a swarm of crows, rats, or piranhas for INT turns. You only take damage from area effects.`,
	`TELEKINESIS
You may mentally manipulate items (one at a time) up to 10 feet away for INT turns.`,
	`TELEPATHY
You can project your thoughts into a mind within INT hexes.`,
	`TELEPORT
An object teleports to a clear patch of ground up to INT × 40’ away from its origin point.`,
	`THAUMATURGIC ANCHOR
An object becomes the target of every spell cast within 120’ of it for INT turns.`,
	`THICKET
A thicket of trees and dense brush up to INT × 40’ wide sprouts up over the course of one round.`,
	`TIME JUMP
An object disappears as it jumps INT turns into the future. When it returns,
it destroys any matter in its space.`,
	`TIME RUSH
Time within INT × 10’ of you goes 10 times faster than the rest of the world.
Lasts 10 rounds (for you).`,
	`TIME SLOW
Time within INT × 10’ of you goes 10 times slower than the rest of the world.
Lasts 10 rounds (for you).`,
	`TRUTH SENSE
You can detect lies for INT hours.`,
	`UNRAVEL
Cast this as a reaction to another spell of level INT or less going off to nullify it.`,
	`UPWELL
A spring of seawater erupts, producing a thousand cubic feet of water per turn for INT turns.`,
	`VISION
You create an illusory object with full motion and sound that only one creature can sense.
Lasts INT turns.`,
	`VISUAL ILLUSION
You create INT silent, immobile, illusory objects that last until they are touched.`,
	`WARD
A silver circle 40’ across appears on the ground around you. Until you leave the circle,
INT types of things that you name cannot cross it.`,
	`WEB
You can shoot INT × 40’ of strong, sticky web. Lasts until burned.`,
	`WHIRLWIND
You create a vortex of air INT × 10’ wide that can deflect missiles.`,
	`WIZARD MARK
Your finger produces ulfire-colored paint for INT hours.
This paint is only visible to you, and can be seen at any distance, even through objects.`,
	`X-RAY VISION: You can see through INT feet of material.`,
}
