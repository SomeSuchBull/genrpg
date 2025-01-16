package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	"github.com/ttacon/chalk"
)

type StockingContext struct {
	Level int
}

var TableDie = rand.Intn

func D(dieSize int) int {
	return TableDie(dieSize) + 1
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

var B = chalk.Bold.TextStyle
var I = chalk.Italic.TextStyle

func RoomStyle(val string) string {
	return chalk.Cyan.Color(B(chalk.Underline.TextStyle(val)))
}

func DetailStyle(val string) string {
	return chalk.Green.Color(B(val))
}

func SpellStyle(val string) string {
	return chalk.Magenta.Color(B(val))
}

func MonsterStyle(val string) string {
	return chalk.Red.Color(B(val))
}

// NOTE: I don't like this style, play around with other configs of it
func TableStyle(val string) string {
	return chalk.Black.NewStyle().WithBackground(chalk.White).Style(val)
	// WithTextStyle(chalk.Bold).Style(val)
}

func PrintJSON(j any) error {
	var out []byte
	var err error
	out, err = json.MarshalIndent(j, "", "    ")
	if err == nil {
		fmt.Println(string(out))
	}
	return err
}

type Distribution struct {
	RangeMin, RangeMax, ValueMin, ValueMax int
	Extreme                                float64
	ResultContinuos                        map[int]float64
	ResultDiscrete                         map[int]int
}

func (d *Distribution) Generate() error {
	if d.RangeMax < d.RangeMin {
		return fmt.Errorf("RangeMax < RangeMin")
	}
	if d.ValueMax < d.ValueMin {
		return fmt.Errorf("ValueMax < ValueMin")
	}
	d.ResultContinuos = make(map[int]float64)
	d.ResultDiscrete = make(map[int]int)
	begin := d.RangeMin
	end := d.RangeMax
	if d.Extreme != 0 && end-begin > 2 {
		if d.Extreme < 0 && d.Extreme > 1 {
			return fmt.Errorf("Extreme must be between 0 and 1")
		}
		min := (float64(d.ValueMax) + float64(d.ValueMin)) / 2 * d.Extreme
		max := (float64(d.ValueMax) + float64(d.ValueMin)) / 2 * (1 + d.Extreme)
		d.ResultContinuos[begin] = min
		d.ResultDiscrete[begin] = int(math.Round(min))
		d.ResultContinuos[end] = max
		d.ResultDiscrete[end] = int(math.Round(max))
		begin++
		end--
	}
	steps := end - begin
	increment := float64(d.ValueMax-d.ValueMin) / float64(steps)
	for i := begin; i <= end; i++ {
		d.ResultContinuos[i] = float64(d.ValueMin) + float64(i-begin)*increment
		d.ResultDiscrete[i] = int(math.Round(d.ResultContinuos[i]))
	}
	return nil
}
