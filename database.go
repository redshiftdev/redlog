package redlog

import (
	"fmt"
)

type DB interface {
	Start() error
	Write(level LogType, message string)
	Close() error
}

type DatabaseLog struct {
	BaseLog
	db DB
}

func NewDatabaseLog(db DB) (*DatabaseLog, error) {
	log := DatabaseLog{
		BaseLog: NewBaseLog(LEVEL_NON_DEBUG),
		db:      db,
	}
	go log.listen()
	err := <-log.err
	if err != nil {
		return nil, fmt.Errorf("unable to start database log: %v", err)
	}
	return &log, nil
}

//listen will listen for log messages and write them
func (log *DatabaseLog) listen() {
	err := log.db.Start()
	if err != nil {
		log.err <- fmt.Errorf("could start database log: %v", err)
		return
	}
	log.err <- nil
	for {
		msg, ok := <-log.io
		if !ok {
			break
		}
		if (uint8(log.level) & uint8(msg.level)) > 0 {
			log.db.Write(msg.level, msg.message)
		}
	}
	err = log.db.Close()
	if err != nil {
		log.err <- fmt.Errorf("could not close database log: %v", err)
		return
	}
	log.err <- nil
}

//Close closes the file
func (log *DatabaseLog) Close() error {
	err := log.BaseLog.Close()
	if err != nil {
		return fmt.Errorf("error closing database log: %v", err)
	}
	return nil
}
