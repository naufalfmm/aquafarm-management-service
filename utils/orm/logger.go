package orm

type Logger interface {
	Printf(format string, args ...interface{})
}
