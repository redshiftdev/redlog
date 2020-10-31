package redlog

import "strconv"

//LogLevel log level
type LogLevel uint8

//LEVEL log levels
const (
	LEVEL_DISABLED LogLevel = 0b00000000
	LEVEL_ALL      LogLevel = 0b11111111

	LEVEL_NON_DEBUG  LogLevel = LEVEL_ALL ^ LogLevel(TYPE_DEBUG)
	LEVEL_PRODUCTION LogLevel = LogLevel(TYPE_ERROR | TYPE_CRITICAL)
)

//String returns the human readable LogType
func (i LogLevel) String() string {
	switch i {
	case LEVEL_DISABLED:
		return "DISABLED"
	case LEVEL_ALL:
		return "ALL"
	case LEVEL_NON_DEBUG:
		return "NON DEBUG"
	case LEVEL_PRODUCTION:
		return "PRODUCTION"
	default:
		return "LogLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
