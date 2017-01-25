package lg

import (
	"fmt"
	"io"
	"log"
)

type Logger interface {
	Logf(lvl Level, format string, args ...interface{})
	GetLevel() Level
}

type Basic struct {
	io.Writer
	Level
	Formatter
}

func New(w io.Writer) *Basic {
	return &Basic{
		Writer:    w,
		Level:     LvlInfo,
		Formatter: DefaultFormatter,
	}
}

func (b *Basic) Logf(lvl Level, format string, args ...interface{}) {
	log.Print(lvl, b.GetLevel())
	if IsEnabledFor(b, lvl) {
		fmt.Fprintf(b, b.Format(NewRecord(b.Level, format, args...)))
	}
}

func (b *Basic) GetLevel() Level { return b.Level }

func IsEnabledFor(lg Logger, lvl Level) bool {
	return lg.GetLevel() <= lvl
}
