package lg

import "io"

type Writer struct {
	io.Writer
	Formatter
	Level
}

func (w *Writer) Logf(lvl Level, format string, args ...interface{}) {
	if IsEnabledFor(w, lvl) {
		w.Format(w, NewRecord(lvl, format, args...))
	}
}

func (w *Writer) GetLevel() Level { return w.Level }

var _ Loggerf = &Writer{}
