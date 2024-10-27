package pirateborg

import (
	"fmt"
	"math"

	"github.com/genrpg/utils"
	"github.com/ttacon/chalk"
)

type Pirate struct {
	Name        string
	Description string
	Morale      int
	HP          int
	Armor       string
}

func GeneratePirate() Pirate {
	var pirate Pirate
	die1 := utils.D(12)
	die2 := utils.D(12)
	die3 := utils.D(12)
	tableAResult := pirateTableA[die1-1]
	tableBResult := pirateTableB[die2-1]
	tableCResult := pirateTableC[die3-1]
	pirate.Description = fmt.Sprintf("A %s with %s wielding a %s", tableAResult, tableBResult, tableCResult)
	pirate.Name = PirateName()
	pirate.Morale = int(math.Max(float64(die1), math.Max(float64(die2), float64(die3))))
	pirate.HP = 2 + int(math.Min(math.Min(float64(die1), float64(die2)), float64(die3)))
	switch {
	case pirate.Morale == die3:
		pirate.Armor = "-d4"
	case pirate.Morale == die2:
		pirate.Armor = "-d2"
	default:
		pirate.Armor = "no armor"
	}
	return pirate
}

func PirateName() string {
	var firstName string
	if utils.D(2) == 1 {
		firstName = pirateMaleName[utils.TableDie(36)]
	} else {
		firstName = pirateFemaleName[utils.TableDie(36)]
	}
	return fmt.Sprintf("%s %s \"%s %s\"", firstName, pirateSurname[utils.TableDie(36)], pirateNickname[utils.TableDie(36)], pirateNickname[utils.TableDie(36)])
}

func MysticalMishap() string {
	die := utils.D(20)
	switch die {
	case 1:
		return "All before you vanishes, and you are alone on a gray island. A figure robed in black can be seen on a boat in the distance.\n" + I("Welcome to"+chalk.Underline.TextStyle("purgatory"))
	case 2:
		return "Everyone within d20 feet vomits briny water and sea urchins (and loses 4 HP)."
	case 3:
		return "A " + chalk.Underline.TextStyle("wraith") + " appears, the ghost of an old captain. It is " + reaction[utils.D(6)+utils.D(6)] + ".\n" + I("It leaves after "+fmt.Sprint(utils.D(6))+" rounds.")
	case 4:
		return "You are possessed by an ancient, arcane spirit for " + fmt.Sprint(utils.D(4)) + " rounds. You are " + reaction[utils.D(6)+utils.D(6)] + " to others."
	case 5:
		return "You go insane, as you become convinced that this world is a dream and that you must wake up. Test SPIRIT DR12 every dawn to see if you come to your senses."
	case 6:
		return "Nothing happens" + I("... until the Kraken appears in "+fmt.Sprint(utils.D(6))+" days")
	case 7:
		return "Gravity behaves as if everything within 30' is underwater" + I("... for "+fmt.Sprint(utils.D(8))+" minutes")
	case 8:
		return fmt.Sprint(utils.D(12)) + "BIRDS! They swarm you and all around.\nHP 2 Morale - No Armor Bite d2.\n" + I("They flee when half are defeated.")
	case 9:
		return "All metal within 30' become molten hot and glows red. The effect lasts for 1 minute.\n" + I("Any one touching metal after the first round is burned for d2 damage.")
	case 10:
		return "Your vision permanently fills with water and the world around you looks submerged and obfuscated. Presence Tests involving sight are +4 DR from now on.\n" + I("Underwater you see like fish.")
	case 11:
		return "The weather and time of day are immediately, completely different.\n" + I("You and everything within a 3 mile radius have traveled "+
			fmt.Sprint(utils.D(12)+utils.D(12)+utils.D(12)+utils.D(12))+" hours into the future.")
	case 12:
		return "Clouds darken, winds pick up, the temperature drops " + fmt.Sprint(10*utils.D(10)) + " degrees.\n" + I("A thunderstorm moves in.")
	case 13:
		return "A massive earthquake shakes the ground. At sea, tidal waves swell."
	case 14:
		return "You've simply forgotten how to cast the ritual.\n" + I("You no longer know it. Remove it from your sheet.")
	case 15:
		return "Eldritch voices whisper in your ears. Test SPIRIT DR12 or take d2 damage out of horror.\n" + I("If you pass, then next time you consider accessing the arcane you discover you know one random new ritual")
	case 16:
		return "The ritual succeeds, but in the worst way possible. GM decides: different target, goes off at the wrong time, you cast it on yourself, etc."
	case 17:
		years := fmt.Sprint(utils.D(10) + utils.D(10) + utils.D(10))
		return "Your mind is lost at sea for what seems like " + years + " years, yet no time has passed.\n" + I("Others watch as you age "+years+" years in a matter of seconds")
	case 18:
		return "Hundreds of tiny crabs swarm you. They are harmless, and provide an extra -2 armor.\n" + I("They leave after "+fmt.Sprint(utils.D(6))+" days, but return every full moon.")
	case 19:
		return "Everyone you can see recovers " + fmt.Sprint(utils.D(6)+utils.D(6)+utils.D(6)) + " HP.\n" + I("including your enemies...")
	case 20:
		return "Time stops. You experience a moment of absolute cosmic understanding and bliss.\n" + I("When you gain your composure, time resumes and you have learned a new ritual of your choice and your Spirit score is increased by +1.")
	default:
		return ""
	}
}
