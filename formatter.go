package lg

import "fmt"

type Formatter interface {
	Format(Record) string
}

type FormatterFunc func(Record) string

func (f FormatterFunc) Format(r Record) string { return f(r) }

var DefaultFormatter FormatterFunc = func(r Record) string {
	return fmt.Sprintf(
		"%v - %v - %v",
		r.Time,
		r.Level,
		fmt.Sprintf(r.Format, r.Args...),
	)
}
