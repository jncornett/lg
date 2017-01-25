package lg

type Loggerf interface {
	Logf(lvl Level, format string, args ...interface{})
	GetLevel() Level
}
