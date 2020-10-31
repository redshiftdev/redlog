package redlog

const LOG_FORMAT string = "%s | [%s] - %s\n"

//LogFlags container for flags
type LogFlags uint

type logChan chan logMessage
type errChan chan error

type logMessage struct {
	level   LogType
	message string
}

//ILog logging interface
type ILog interface {
	Level(level LogLevel)
	Flags(flags LogFlags)
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
	Critical(message string)
	Close() error
}

//BaseLog base log
type BaseLog struct {
	level LogLevel
	flags LogFlags
	io    logChan
	err   errChan
}

func NewBaseLog(level LogLevel) BaseLog {
	return BaseLog{
		level: level,
		io:    make(logChan),
		err:   make(errChan),
	}
}

func (log *BaseLog) Close() error {
	if log.io != nil {
		close(log.io)
		if log.err != nil {
			return <-log.err
		}
	}
	return nil
}

//Level sets write level
func (log *BaseLog) Level(level LogLevel) {
	log.level = level
}

//Flags sets write flags
func (log *BaseLog) Flags(flags LogFlags) {
	log.flags = flags
}

//Debug generates a DEBUG level log
func (log *BaseLog) Debug(message string) {
	log.io <- logMessage{
		level:   TYPE_DEBUG,
		message: message,
	}
}

//Info generates an INFO level log
func (log *BaseLog) Info(message string) {
	log.io <- logMessage{
		level:   TYPE_INFO,
		message: message,
	}
}

//Warning generates a WARNING level log
func (log *BaseLog) Warning(message string) {
	log.io <- logMessage{
		level:   TYPE_WARNING,
		message: message,
	}
}

//Error generates an ERROR level log
func (log *BaseLog) Error(message string) {
	log.io <- logMessage{
		level:   TYPE_ERROR,
		message: message,
	}
}

//Critical generates a CRITICAL level log
func (log *BaseLog) Critical(message string) {
	log.io <- logMessage{
		level:   TYPE_CRITICAL,
		message: message,
	}
}
