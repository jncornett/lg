package lg

type Null struct {
	Level
}

func (n Null) GetLevel() Level                    { return n.Level }
func (n Null) Logf(Level, string, ...interface{}) {}

var _ Loggerf = &Null{}
