package console

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Logger struct {
	Component string
}

func NewLogger(ModuleName string) *Logger {
	return &Logger{Component: strings.ToUpper(ModuleName)}
}

func (l *Logger) Info(obj ...any) {
	s := ""

	for _, v := range obj {
		s = s + fmt.Sprint(v) + " "
	}

	fmt.Fprintf(os.Stdout, "\r \r")
	log.Println(fmt.Sprintf("[%v|INFO]", l.Component), s)
}

func (l *Logger) Panic(obj ...any) {
	s := ""

	for _, v := range obj {
		s = s + fmt.Sprint(v) + " "
	}

	fmt.Fprintf(os.Stdout, "\r \r")
	log.Println(fmt.Sprintf("[%v|WARN]", l.Component), s)
}

func (l *Logger) Fatal(obj ...any) {
	s := ""

	for _, v := range obj {
		s = s + fmt.Sprint(v) + " "
	}

	fmt.Fprintf(os.Stdout, "\r \r")
	log.Fatalln(fmt.Sprintf("[%v|FATAL]", l.Component), s)
}
