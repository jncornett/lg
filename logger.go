package lg

import (
	"io"
	"os"
)

type Logger struct {
	Loggerf
	Wrap
}

func New(w io.Writer, f Formatter, lvl Level) *Logger {
	if w == nil {
		w = os.Stderr
	}
	if f == nil {
		f = DefaultFormatter
	}
	lg := &Logger{
		Loggerf: &Writer{
			Writer:    w,
			Formatter: f,
			Level:     lvl,
		},
	}
	lg.Wrap.Loggerf = lg.Loggerf
	return lg
}
