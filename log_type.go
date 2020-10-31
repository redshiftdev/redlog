package redlog

import "strconv"

//LogType bit flag
type LogType uint8

//TYPE log types
const (
	TYPE_DEBUG    LogType = 0b00000001
	TYPE_INFO     LogType = 0b00000010
	TYPE_WARNING  LogType = 0b00000100
	TYPE_ERROR    LogType = 0b00001000
	TYPE_CRITICAL LogType = 0b00010000
)

//String returns the human readable LogType
func (i LogType) String() string {
	switch i {
	case TYPE_DEBUG:
		return "DEBUG"
	case TYPE_INFO:
		return "INFO"
	case TYPE_WARNING:
		return "WARNING"
	case TYPE_ERROR:
		return "ERROR"
	case TYPE_CRITICAL:
		return "CRITICAL"
	default:
		return "LogType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
