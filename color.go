package lg

import (
	"fmt"
	"io"

	"github.com/ttacon/chalk"
)

var ColorFormatter FormatterFunc = func(w io.Writer, r Record) {
	fmt.Fprintf(
		w,
		"[ %v ] %v\n",
		fmt.Sprint(getColor(r.Level), r.Level, chalk.Reset),
		fmt.Sprintf(r.Format, r.Args...),
	)
}

var colors = [...]chalk.Color{
	LvlDebug: chalk.Blue,
	LvlInfo:  chalk.Green,
	LvlError: chalk.Red,
}

func getColor(lvl Level) (c chalk.Color) {
	if lvl >= 0 && lvl < Level(len(colors)) {
		c = colors[lvl]
	}
	return
}
