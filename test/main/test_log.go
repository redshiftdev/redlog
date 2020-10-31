package main

import (
	"os"

	"github.com/redshiftdev/redlog"
)

func main() {
	// working directory assumed to be project root!
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileLog, err := redlog.NewFileLog(cwd + "/test/data/file.log")
	if err != nil {
		panic(err)
	}
	dbDriver, err := NewDatabaseDriver(cwd + "/test/data/sqlite.db")
	if err != nil {
		panic(err)
	}
	dbLog, err := redlog.NewDatabaseLog(dbDriver)
	consoleLog := redlog.NewConsoleLog()
	log := redlog.NewOmniLog(consoleLog, fileLog, dbLog)

	log.Debug("debug message")
	log.Info("info message")
	log.Warning("warning message")
	log.Error("error message")
	log.Critical("critical message")

	log.Level(redlog.LEVEL_ALL)
	log.Flags(redlog.OMNI_FLAG_LEVEL_ENABLED)
	log.Level(redlog.LEVEL_ALL)

	log.Debug("debug message")
	log.Info("info message")
	log.Warning("warning message")
	log.Error("error message")
	log.Critical("critical message")

	err = log.Close()
	if err != nil {
		panic(err)
	}
}
