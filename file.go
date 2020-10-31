package redlog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//FileLog manages logfile access
type FileLog struct {
	BaseLog
	path string
}

//NewFileLog creates a new FileLog from a given path
func NewFileLog(path string) (*FileLog, error) {
	log := FileLog{
		BaseLog: NewBaseLog(LEVEL_PRODUCTION),
		path:    path,
	}
	go log.listen()
	err := <-log.err
	if err != nil {
		return nil, fmt.Errorf("unable to start file log: %v", err)
	}
	return &log, nil
}

//listen will listen for log messages and write them
func (log *FileLog) listen() {
	logDir := filepath.Dir(log.path)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err = os.MkdirAll(logDir, os.ModePerm); err != nil {
			log.err <- fmt.Errorf("could not initialize log directory: %v", err)
			return
		}
	} else if err != nil {
		log.err <- fmt.Errorf("unknown filesystem error: %v", err)
		return
	}
	file, err := os.OpenFile(log.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0664))
	if err != nil {
		log.err <- fmt.Errorf("could not open log file: %v", err)
		return
	}
	log.err <- nil
	for {
		msg, ok := <-log.io
		if !ok {
			break
		}
		if (uint8(log.level) & uint8(msg.level)) > 0 {
			fmt.Fprintf(file, LOG_FORMAT,
				time.Now().Format("2006-01-02 15:04:05"), msg.level, msg.message,
			)
		}
	}
	err = file.Close()
	if err != nil {
		log.err <- fmt.Errorf("could not close log file: %v", err)
		return
	}
	log.err <- nil
}

//Close closes the file
func (log *FileLog) Close() error {
	err := log.BaseLog.Close()
	if err != nil {
		return fmt.Errorf("error closing file log: %v", err)
	}
	return nil
}
