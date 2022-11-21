package interfaces

type Logger interface {
	Info(msg string, fields Fields)
	Error(msg string, fields Fields)
	Debug(msg string, fields Fields)
	Warn(msg string, fields Fields)
	Fatal(msg string, fields Fields)
	Panic(msg string, fields Fields)
}

type Fields map[string]any
