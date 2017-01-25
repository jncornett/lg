package lg

import (
	"io"
	"os"
)

type Logger struct {
	io.Writer
	Level
	Formatter
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
		Writer:    w,
		Level:     lvl,
		Formatter: f,
	}
	lg.Wrap.Loggerf = lg
	return lg
}

func (b *Logger) Logf(lvl Level, format string, args ...interface{}) {
	if IsEnabledFor(b, lvl) {
		b.Format(b, NewRecord(b.Level, format, args...))
	}
}

func (b *Logger) GetLevel() Level { return b.Level }

func IsEnabledFor(lg Loggerf, lvl Level) bool {
	return lg.GetLevel() <= lvl
}
