package logger

import (
	"fmt"
	"github.com/mgutz/ansi"
	"time"
)

type Logger interface {
	Info(f string, args ...any)
	Warning(f string, args ...any)
	Error(f string, args ...any)
	Critical(f string, args ...any)
}

type ColorLogger struct {
	name string
}

func NewLogger(name string) ColorLogger {
	return ColorLogger{name: name}
}

func (l ColorLogger) getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (l ColorLogger) Info(f string, args ...any) {
	fmt.Printf(fmt.Sprintf("[%s] [%s/%s]: %s\n", l.getTime(), l.name, ansi.Color("INFO", "green"), f), args...)
}

func (l ColorLogger) Warning(f string, args ...any) {
	fmt.Printf(fmt.Sprintf("[%s] [%s/%s]: %s\n", l.getTime(), l.name, ansi.Color("WARNING", "yellow+b"), ansi.Color(f, "yellow")), args...)
}

func (l ColorLogger) Error(f string, args ...any) {
	fmt.Printf(fmt.Sprintf("[%s] [%s/%s]: %s\n", l.getTime(), l.name, ansi.Color("ERROR", "red"), ansi.Color(f, "red")), args...)
}

func (l ColorLogger) Critical(f string, args ...any) {
	fmt.Printf(fmt.Sprintf("[%s] [%s/%s]: %s\n", l.getTime(), l.name, ansi.Color("CRITICAL", "red+b"), ansi.Color(f, "red+b")), args...)
}
