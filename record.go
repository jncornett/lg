package lg

import "time"

type Record struct {
	Format string
	Args   []interface{}
	Level
	time.Time
}

func NewRecord(lvl Level, format string, args ...interface{}) Record {
	return Record{
		Format: format,
		Args:   args,
		Level:  lvl,
		Time:   time.Now(),
	}
}
