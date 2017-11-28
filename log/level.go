package log

type Level int

const (
	DEBUG = Level(1)
	INFO  = Level(2)
	WARN  = Level(3)
	ERROR = Level(4)
	FATAL = Level(5)
)

func (l Level) String() string {
	switch l {
	case 1:
		return "DEBUG"
	case 2:
		return "INFO"
	case 3:
		return "WARNING"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	}
	panic("invalid Level")
}
