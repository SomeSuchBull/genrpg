package pirateborg

import "github.com/ttacon/chalk"

type Item struct {
	Price string
	Name  string
	Extra string
}

type Weapon struct {
	Item
	Damage string
}

type Clothing struct {
	Item
	Armor string
}

type Hat struct {
	Item
	Armor string
}

type AncientRelic struct {
	Item
}

var AncientRelicRules = "RELICS can be used to access ancient magical powers. Use your action to activate one in your possession.\n" + B("After") + " using a RELIC, test " + chalk.Underline.TextStyle("SPIRIT DR12") + ". If you fail, you are stunned for the next round and that RELIC cannot be used until the next dawn.\nOn a Fumble, the relic is permanently destroyed or depleted, and then you lose " + B("d2 HP") + ". The GM decides the effect on a Crit, if any."

func GetAncientRelic(r int) Item {
	if r < 0 || r >= len(ancientRelics) {
		return Item{}
	}
	return ancientRelics[r]
}
