package redlog

import (
	"fmt"
	"strings"
)

const (
	OMNI_FLAG_LEVEL_ENABLED LogFlags = 0b00000001
)

type OmniLog struct {
	flags LogFlags
	logs  []ILog
}

func NewOmniLog(logs ...ILog) *OmniLog {
	return &OmniLog{logs: logs}
}

func (log *OmniLog) Level(level LogLevel) {
	if (log.flags & OMNI_FLAG_LEVEL_ENABLED) > 0 {
		for _, ilog := range log.logs {
			ilog.Level(level)
		}
	} else {
		fmt.Println("WARNING: must use OMNI_FLAG_LEVEL_ENABLED to enabled OmniFlag.Level()!")
	}
}

func (log *OmniLog) Flags(flags LogFlags) {
	log.flags = flags
}

func (log *OmniLog) Debug(message string) {
	for _, ilog := range log.logs {
		ilog.Debug(message)
	}
}

func (log *OmniLog) Info(message string) {
	for _, ilog := range log.logs {
		ilog.Info(message)
	}
}

func (log *OmniLog) Warning(message string) {
	for _, ilog := range log.logs {
		ilog.Warning(message)
	}
}

func (log *OmniLog) Error(message string) {
	for _, ilog := range log.logs {
		ilog.Error(message)
	}
}

func (log *OmniLog) Critical(message string) {
	for _, ilog := range log.logs {
		ilog.Critical(message)
	}
}

func (log *OmniLog) Close() error {
	errs := make([]string, 0)
	for _, ilog := range log.logs {
		err := ilog.Close()
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf(
			"1 or more errors occured during logs close: \n%s\n",
			strings.Join(errs, "\n"),
		)
	}
	return nil
}
