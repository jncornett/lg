package lg

type Elevator struct {
	Offset int
	Logger
}

func (e *Elevator) Logf(lvl Level, format string, args ...interface{}) {
	e.Logger.Logf(e.Logger.GetLevel()+Level(e.Offset), format, args...)
}

func (e *Elevator) GetLevel() Level {
	return e.Logger.GetLevel() - Level(e.Offset)
}
