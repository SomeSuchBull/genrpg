package knave

import (
	"fmt"

	"github.com/genrpg/utils"
)

// TIP: Roll twice on these tables and combine the results.
var engine = map[int]func(utils.StockingContext, *string) string{
	0: getEmpty, 1: getEmpty, 2: getStockingMonster, 3: getStockingMonster, 4: getSpecial, 5: getTrap,
}

// TODO: something with this ctx, like a dungeon theme, common monsters, etc.
func Stocking(rooms, level int64, verbose bool) string {
	// verboseOutput := ""
	fmt.Println("Theme | Detail | Detail")
	for i := 0; i < int(rooms); i++ {
		roll := td(6)
		f := engine[roll]
		room(f, utils.StockingContext{Level: int(level)}, i, nil)
	}
	return ""
}

func room(f func(utils.StockingContext, *string) string, ctx utils.StockingContext,
	roomNumber int, verboseOutput *string) {
	contents := f(ctx, verboseOutput)
	fmt.Printf("\n%s\n%s\n%s\n\n",
		utils.RoomStyle(fmt.Sprintf("%03d %s", roomNumber+1, recursiveTableRoll(rooms[td(100)], nil))),
		fmt.Sprintf("%s | %s | %s",
			utils.DetailStyle(recursiveTableRoll(roomThemes[td(100)], nil)),
			utils.DetailStyle(recursiveTableRoll(roomDetails[td(100)], nil)),
			utils.DetailStyle(recursiveTableRoll(roomDetails[td(100)], nil))),
		contents)
}

var levelTreasure = map[int]map[string]int{
	0: {"SP": d(6) * 100, "GP": c(2) * d(6) * 100, "Gems": c(20) * d(6), "Pieces of jewellery": c(50) * d(6), "Magic item": c(50)},
	1: {"SP": d(12) * 100, "GP": c(2) * d(6) * 100, "Gems": c(10) * d(6), "Pieces of jewellery": c(20) * d(6), "Magic item": c(25, 2)},
	2: {"SP": d(6) * 1000, "GP": d(6) * 200, "Gems": c(5) * d(6), "Pieces of jewellery": c(10) * d(6), "Magic item": c(10)},
	3: {"SP": d(6) * 2000, "GP": d(6) * 500, "Gems": c(10, 3) * d(6), "Pieces of jewellery": c(20, 3) * d(6), "Magic item": c(20, 3)},
	4: {"SP": d(6) * 5000, "GP": d(6) * 1000, "Gems": c(5, 2) * d(12), "Pieces of jewellery": c(5) * d(12), "Magic item": c(5)},
}
