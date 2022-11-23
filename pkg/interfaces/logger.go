package interfaces

type Fields map[string]any

// Logger is an interface consisting of the core logger methods.
type Logger interface {
	Info(msg string, fields Fields)
	Error(msg string, fields Fields)
	Debug(msg string, fields Fields)
	Warn(msg string, fields Fields)
	Fatal(msg string, fields Fields)
	Panic(msg string, fields Fields)
}
