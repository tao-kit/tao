package console

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
)

// Console describes an abstract printer.
type Console interface {
	Info(msg ...any)
	InfoF(format string, msg ...any)
	Debug(msg ...any)
	DebugF(format string, msg ...any)
	Warning(msg ...any)
	WarningF(format string, msg ...any)
	Error(msg ...any)
	ErrorF(format string, msg ...any)
	Fatal(msg ...any)
	FatalF(format string, msg ...any)
	Panic(msg ...any)
	PanicF(format string, msg ...any)
}

// Assert *colorConsole implements Console.
var _ Console = (*colorConsole)(nil)

// NewColorConsole returns an instance of Console.
func NewColorConsole() Console {
	return &colorConsole{}
}

type colorConsole struct{}

func (c *colorConsole) Info(msg ...any) {
	fmt.Println(aurora.Green(fmt.Sprint(msg...)))
}

func (c *colorConsole) InfoF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(aurora.Green(msg))
}

func (c *colorConsole) Debug(msg ...any) {
	fmt.Println(aurora.Blue(fmt.Sprint(msg...)))
}

func (c *colorConsole) DebugF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(aurora.Blue(msg))
}

func (c *colorConsole) Warning(msg ...any) {
	fmt.Println(aurora.Yellow(fmt.Sprint(msg...)))
}

func (c *colorConsole) WarningF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(aurora.Yellow(msg))
}

func (c *colorConsole) Error(msg ...any) {
	fmt.Println(aurora.Red(fmt.Sprint(msg...)))
}

func (c *colorConsole) ErrorF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	fmt.Println(aurora.Red(msg))
}

func (c *colorConsole) Fatal(msg ...any) {
	log.Fatalln(fmt.Sprint(msg...))
}

func (c *colorConsole) FatalF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Fatalln(aurora.Red(msg))
}

func (c *colorConsole) Panic(msg ...any) {
	panic(fmt.Sprint(msg...))
}

func (c *colorConsole) PanicF(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	panic(msg)
}
