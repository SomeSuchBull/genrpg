package utils

import (
	"math/rand"

	"github.com/ttacon/chalk"
)

type StockingContext struct {
	Level int
}

func D(dieSize int) int {
	return rand.Intn(dieSize) + 1
}

func TableDie(dieSize int) int {
	return rand.Intn(dieSize)
}

func Chance(size int, c ...int) (happened int) {
	chance := 1
	if len(c) > 0 {
		chance = c[0]
	}
	if rand.Intn(size)/chance == 0 {
		happened = 1
	}
	return
}

func RoomStyle(val string) string {
	return chalk.Cyan.Color(chalk.Bold.TextStyle(chalk.Underline.TextStyle(val)))
}

func DetailStyle(val string) string {
	return chalk.Green.Color(chalk.Bold.TextStyle(val))
}

func SpellStyle(val string) string {
	return chalk.Magenta.Color(chalk.Bold.TextStyle(val))
}

func MonsterStyle(val string) string {
	return chalk.Red.Color(chalk.Bold.TextStyle(val))
}

// NOTE: I don't like this style, play around with other configs of it
func TableStyle(val string) string {
	return chalk.Black.NewStyle().WithBackground(chalk.White).Style(val)
	// WithTextStyle(chalk.Bold).Style(val)
}
