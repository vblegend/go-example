package log

// Logger is a generic logging interface
type Logger interface {
	// Init initialises options
	Init(options ...Option) error
	// Options The Logger options
	Options() Options
	// Log writes a log entry
	Log(level Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level Level, format string, v ...interface{})
	// String returns the name of logger
	String() string
}
