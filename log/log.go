package log

import (
	"log"
	"os"
)

// ебашим Лог чтобы меньше писать кот
type Logger interface {
	Log(a interface{})
}

type ConsoleLog struct {
	logger *log.Logger
}

func (x *ConsoleLog) Init() {
	x.logger = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func (x *ConsoleLog) Log(a interface{}) {
	x.logger.Print(a)
}
