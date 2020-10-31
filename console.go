package redlog

import (
	"fmt"
	"time"
)

type ConsoleLog struct {
	BaseLog
}

func NewConsoleLog() *ConsoleLog {
	log := ConsoleLog{BaseLog: NewBaseLog(LEVEL_ALL)}
	go log.listen()
	return &log
}

func (log *ConsoleLog) listen() {
	for {
		msg, ok := <-log.io
		if !ok {
			break
		}
		if (uint8(log.level) & uint8(msg.level)) > 0 {
			fmt.Printf(LOG_FORMAT,
				time.Now().Format("2006-01-02 15:04:05"), msg.level, msg.message,
			)
		}
	}
	log.err <- nil
}
