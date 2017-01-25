package lg

import (
	"fmt"
	"io"
	"time"
)

type Formatter interface {
	Format(io.Writer, Record)
}

type FormatterFunc func(io.Writer, Record)

func (f FormatterFunc) Format(w io.Writer, r Record) { f(w, r) }

var DefaultFormatter FormatterFunc = func(w io.Writer, r Record) {
	fmt.Fprintf(
		w,
		"%v.%v %v %v\n",
		r.Time.Format("2006-01-02 15:04:05"),
		r.Time.Nanosecond()/int(time.Millisecond),
		r.Level,
		fmt.Sprintf(r.Format, r.Args...),
	)
}
