package lg

import (
	"fmt"
	"io"
)

var CliFormatter FormatterFunc = func(w io.Writer, r Record) {
	fmt.Fprintf(
		w,
		"[ %v ] %v\n",
		r.Level,
		fmt.Sprintf(r.Format, r.Args...),
	)
}
