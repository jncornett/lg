package lg

import "fmt"

type Wrap struct {
	Loggerf
}

func (w Wrap) Debug(args ...interface{}) { w.Log(LvlDebug, args...) }
func (w Wrap) Debugf(format string, args ...interface{}) {
	w.Logf(LvlDebug, format, args...)
}

func (w Wrap) Info(args ...interface{}) { w.Log(LvlInfo, args...) }
func (w Wrap) Infof(format string, args ...interface{}) {
	w.Logf(LvlInfo, format, args...)
}

func (w Wrap) Error(args ...interface{}) { w.Log(LvlError, args...) }
func (w Wrap) Errorf(format string, args ...interface{}) {
	w.Logf(LvlError, format, args...)
}

func (w Wrap) Fatal(args ...interface{}) {
	s := fmt.Sprint(args...)
	w.Log(LvlError, s)
	panic(s)
}

func (w Wrap) Fatalf(format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	w.Log(LvlError, s)
	panic(s)
}

func (w Wrap) Log(lvl Level, args ...interface{}) {
	w.Logf(lvl, fmt.Sprint(args...))
}

func (w Wrap) IsEnabledFor(lvl Level) bool {
	return IsEnabledFor(w, lvl)
}
