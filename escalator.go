package lg

type Elevator struct {
	Offset int
	Loggerf
}

func (e *Elevator) Logf(lvl Level, format string, args ...interface{}) {
	e.Loggerf.Logf(e.Loggerf.GetLevel()+Level(e.Offset), format, args...)
}

func (e *Elevator) GetLevel() Level {
	return e.Loggerf.GetLevel() - Level(e.Offset)
}
