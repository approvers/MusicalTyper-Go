package Logger

import (
	"log"
)

type logger struct {
	SectionName string
}

func NewLogger(Entryname string) logger {
	return logger{SectionName: Entryname}
}

func (e *logger) CheckError(err error) {
	if err != nil {
		log.Printf("[%s/Fatal] %s\n", e.SectionName, err)
		panic("") //for call stack.
	}
}

func (e *logger) FatalError(msg string) {
	log.Printf("[%s/Fatal] %s\n", e.SectionName, msg)
	panic("") //for call stack.
}

func (e *logger) FatalErrorWithoutExit(msg string) {
	log.Printf("[%s/Fatal] %s\n", e.SectionName, msg)
}

func (e *logger) Warn(msg string) {
	log.Printf("[%s/Warn] %s\n", e.SectionName, msg)
}
