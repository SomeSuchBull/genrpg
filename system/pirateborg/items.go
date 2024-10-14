package pirateborg

import (
	"fmt"

	"github.com/ttacon/chalk"
)

type Item struct {
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
	Extra string `json:"extra,omitempty"`
}

type Weapon struct {
	Item
	Damage string `json:"damage,omitempty"`
}

func (i Item) String() string {
	name := i.Name
	price := i.Price
	extra := i.Extra
	if name == "" {
		name = "-"
	}
	if price == "" {
		price = "-"
	}
	if extra == "" {
		extra = "-"
	}
	return fmt.Sprintf("%s | %s | %s", name, price, extra)
}

func (w Weapon) String() string {
	name := w.Name
	dmg := w.Damage
	price := w.Price
	extra := w.Extra
	if dmg == "" {
		dmg = "-"
	}
	if price == "" {
		price = "-"
	}
	if extra == "" {
		extra = "-"
	}
	return fmt.Sprintf("%s | %s | %s | %s", name, dmg, price, extra)
}

type Clothing struct {
	Item
	Armor string `json:"armor,omitempty"`
}

func (i Clothing) String() string {
	return i.Item.String() + " | " + i.Armor
}

type Hat struct {
	Item
	Armor string `json:"armor,omitempty"`
}

func (i Hat) String() string {
	return i.Item.String() + " | " + i.Armor
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
