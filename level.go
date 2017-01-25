package lg

import "fmt"

type Level int

const (
	LvlDebug Level = iota
	LvlInfo
	LvlError
)

var levels = [...]string{
	LvlDebug: "Debug",
	LvlInfo:  "Info",
	LvlError: "Error",
}

func (lvl Level) String() (s string) {
	if lvl >= 0 && lvl < Level(len(levels)) {
		s = levels[lvl]
	}
	if s == "" {
		s = fmt.Sprintf("%T(%v)", lvl, lvl)
	}
	return
}
